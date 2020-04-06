package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"log"
)

func main() {
	video, err := gocv.OpenVideoCapture("/Users/stutzenbergere/Desktop/Sekiro4pt2.mp4")
	if err != nil {
		log.Fatalln(err)
	}

	window := gocv.NewWindow("Sekiro")
	frame := gocv.NewMat()


	for {
		res := video.Read(&frame)
		if !res {
			break
		}

		res = video.Read(&frame)
		if !res {
			break
		}

		lower := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(17.0, 15.0, 175.0, 0.0), frame.Rows(), frame.Cols(),gocv.MatTypeCV8UC3)
		upper := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(55.0, 95.0, 200.0, 0.0), frame.Rows(), frame.Cols(), gocv.MatTypeCV8UC3)

		mask := gocv.NewMat()
		gocv.InRange(frame, lower, upper, &mask)

		frameMask := gocv.NewMat()
		gocv.BitwiseAndWithMask(frame, frame, &frameMask, mask)

		gray := gocv.NewMat()
		gocv.CvtColor(frameMask, &gray, gocv.ColorBGRToGray)

		blurred := gocv.NewMat()
		gocv.GaussianBlur(gray, &blurred, image.Pt(7, 7), 0, 0, 0)

		contours := gocv.FindContours(blurred, gocv.RetrievalExternal, gocv.ChainApproxSimple)
		if len(contours) > 0 {
			gocv.DrawContours(&frame, contours, -1, color.RGBA{0, 255, 0, 0}, 2)
		}
		window.IMShow(frame)

		if window.WaitKey(1) == 27 {
			break
		}
		//time.Sleep(83 * time.Millisecond)
	}
}
