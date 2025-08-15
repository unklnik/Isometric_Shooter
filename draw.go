package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	gunUD  bool
	gunRO  float32
	gunREC z.Rectangle
)

func DRAWCAM() {

	dISOGRIDfloor(lev)
	//PL PROJ
	if len(prjPL) > 0 {
		dPRJPL()
	}
	dISOGRIDim(lev)

	//PLAY ONCE ANIMS
	clear := false
	if len(playonceANIMS) > 0 {
		for i := range playonceANIMS {
			if playonceANIMS[i].off {
				clear = true
			} else {
				playonceANIMS[i] = dAnimRecONCE(playonceANIMS[i], playonceANIMS[i].rD)
			}
		}
	}
	if clear {
		playonceANIMS = REMANIMOFF(playonceANIMS)
	}
	//DEBUG
	if debug {
		dISOGRIDlines(lev, z.Magenta)
	}

}
func DRAWNOCAMSHADER() { //MARK: NOCAM SHADER
	//MOUSE CURSOR
	/*
		siz := UNIT + UNIT/2
		cursorREC = z.NewRectangle(MS.X, MS.Y, siz, siz)
		if clickT > 0 {
			dIMcolorSHADOW(ETC[5], cursorREC, sepiaDRK, CA(z.Black, 180), UNIT/5)
		} else {
			dIMcolorSHADOW(ETC[4], cursorREC, sepiaDRK, CA(z.Black, 180), UNIT/5)
		}
	*/

}
func DRAWNOCAMNOSHADER() { //MARK: NOCAM NOSHADER
	SCAN(1, 2, z.Fade(z.Black, 0.3))

	if debug {
		DEBUG()

	}
	if scrollONOFF {
		SCROLL()
	}
}
func dPRJPL() {
	for i := range prjPL {
		if !prjPL[i].off {
			//dREC(prjPL[i].r, z.Magenta)
			prjANIMS[0] = dAnimRecLoopShadow(prjANIMS[0], qResizeRecDOUBLE(prjPL[i].r), CA(z.Black, 120), UNIT/2)
		}
	}
}
func dPL() {

	if MS.X < pl.cnt.X {
		pl.lr = true
	} else {
		pl.lr = false
	}

	dIM(ETC[2], R(pl.r.X, pl.r.Y+pl.r.Height/2, pl.r.Width, pl.r.Height)) //SHADOW
	if pl.lr {
		switch pl.st {
		case 0:
			pl.a = dAnimRecLoopFlipColorPART(pl.a, 2, 4, pl.r, z.Magenta)
		case 1:
			pl.a = dAnimRecLoopFlipColor(pl.a, pl.r, z.Magenta)
		}
	} else {
		switch pl.st {
		case 0:
			pl.a = dAnimRecLoopColorPART(pl.a, 2, 4, pl.r, z.Magenta)
		case 1:
			pl.a = dAnimRecLoopColor(pl.a, pl.r, z.Magenta)
		}

	}
	//GUN
	if FRAMES%(FPS/4) == 0 {
		gunUD = !gunUD
	}
	r2 := qResizeRecOffsetSMLR(pl.r, pl.r.Width/8)
	r2.Y += pl.r.Height / 5
	if gunUD {
		r2.Y -= 4
	}
	if pl.lr {
		r2.X -= (r2.Width / 20) * 14
		//ADJUST HEIGHT GUN REC FOR ANGLE
		if gunRO >= 370 && gunRO <= 470 {
			r2.Y -= r2.Height / 4
		} else if gunRO >= 240 && gunRO <= 350 {
			r2.Y += r2.Height / 4
		}
		gunRO = Angl2Points(qRecCNT(r2), MS) + 180
		dIMroFlipColorSHADOW(ETC[3], r2, gunRO, UNIT/10, z.Gray, CA(z.Black, 120))
	} else {
		r2.X += (r2.Width / 20) * 14
		//ADJUST HEIGHT GUN REC FOR ANGLE
		if gunRO >= 10 && gunRO <= 120 {
			r2.Y += r2.Height / 4
		} else if gunRO >= 230 && gunRO <= 350 {
			r2.Y -= r2.Height / 4
		}
		gunRO = Angl2Points(qRecCNT(r2), MS)
		dIMroColorSHADOW(ETC[3], r2, gunRO, UNIT/10, z.Gray, CA(z.Black, 120))
	}
	gunREC = r2

	if debug {
		dRECLINE(gunREC, 2, z.Blue)
	}

	//DEBUG
	if debug {
		dCircCNT(pl.vb1, 4, z.Blue)
		dCircCNT(pl.vb2, 4, z.Blue)
		dCircCNT(pl.vb3, 4, z.Blue)
	}

}
