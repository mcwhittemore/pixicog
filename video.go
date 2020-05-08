package pixicog

import (
	"github.com/3d0c/gmf"
	"image"
	"io"
)

func PixicogFromVideoFileName(srcFileName string) (Pixicog, error) {
	// this is a copy and paste from https://github.com/3d0c/gmf/blob/f4b5acb7db5cbbda9a6209be1d0de5f552823f62/examples/video-to-goImage.go
	cog := Pixicog{}
	var swsctx *gmf.SwsCtx

	inputCtx, err := gmf.NewInputCtx(srcFileName)
	if err != nil {
		return cog, err
	}
	defer inputCtx.Free()

	srcVideoStream, err := inputCtx.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	if err != nil {
		return cog, err
	}

	codec, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	if err != nil {
		return cog, err
	}

	cc := gmf.NewCodecCtx(codec)
	defer gmf.Release(cc)

	cc.SetTimeBase(gmf.AVR{Num: 1, Den: 1})

	cc.SetPixFmt(gmf.AV_PIX_FMT_RGBA).SetWidth(srcVideoStream.CodecCtx().Width()).SetHeight(srcVideoStream.CodecCtx().Height())
	if codec.IsExperimental() {
		cc.SetStrictCompliance(gmf.FF_COMPLIANCE_EXPERIMENTAL)
	}

	if err := cc.Open(nil); err != nil {
		return cog, err
	}
	defer cc.Free()

	ist, err := inputCtx.GetStream(srcVideoStream.Index())
	if err != nil {
		return cog, err
	}
	defer ist.Free()

	// convert source pix_fmt into AV_PIX_FMT_RGBA
	// which is set up by codec context above
	icc := srcVideoStream.CodecCtx()
	if swsctx, err = gmf.NewSwsCtx(icc.Width(), icc.Height(), icc.PixFmt(), cc.Width(), cc.Height(), cc.PixFmt(), gmf.SWS_BICUBIC); err != nil {
		panic(err)
	}
	defer swsctx.Free()

	var (
		pkt        *gmf.Packet
		frames     []*gmf.Frame
		drain      int = -1
		frameCount int = 0
	)

	for {
		if drain >= 0 {
			break
		}

		pkt, err = inputCtx.GetNextPacket()
		if err != nil && err != io.EOF {
			if pkt != nil {
				pkt.Free()
			}
			return cog, err
		} else if err != nil && pkt == nil {
			drain = 0
		}

		if pkt != nil && pkt.StreamIndex() != srcVideoStream.Index() {
			continue
		}

		frames, err = ist.CodecCtx().Decode(pkt)
		if err != nil {
			return cog, err
		}

		// Decode() method doesn't treat EAGAIN and EOF as errors
		// it returns empty frames slice instead. Countinue until
		// input EOF or frames received.
		if len(frames) == 0 && drain < 0 {
			continue
		}

		if frames, err = gmf.DefaultRescaler(swsctx, frames); err != nil {
			panic(err)
		}

		packets, err := cc.Encode(frames, drain)
		if err != nil {
			return cog, err
		}
		if len(packets) == 0 {
			continue
		}

		for _, p := range packets {
			width, height := cc.Width(), cc.Height()

			img := new(image.RGBA)
			img.Pix = p.Data()
			img.Stride = 4 * width
			img.Rect = image.Rect(0, 0, width, height)

			cog = append(cog, img)
		}

		for i, _ := range frames {
			frames[i].Free()
			frameCount++
		}

		if pkt != nil {
			pkt.Free()
			pkt = nil
		}
	}

	for i := 0; i < inputCtx.StreamsCnt(); i++ {
		st, _ := inputCtx.GetStream(i)
		st.CodecCtx().Free()
		st.Free()
	}

	return cog, nil
}
