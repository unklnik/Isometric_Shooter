package main

import (
	"time"

	z "github.com/gen2brain/raylib-go/raylib"
)

func UP() {
	INP()
	UPGAME()
	TIMERS()
}

var (
	FRAMES      int32
	FRAMESTIMER time.Time
	SECONDS     int
)

func TIMERS() {

	//FRAMES
	if FRAMESTIMER.IsZero() {
		FRAMESTIMER = time.Now()
	}
	if time.Since(FRAMESTIMER) >= time.Second/time.Duration(FPS) {
		FRAMES++
		if FRAMES == FPS {
			SECONDS++
			FRAMES = 0
			FRAMESTIMER = time.Time{}
		}
	}
}
func uPRJPL() bool {
	clear := false
	for i := range prjPL {
		if prjPL[i].off {
			clear = true
		} else {
			if cMovePRJPL(prjPL[i]) {
				prjPL[i].cnt.X += prjPL[i].x
				prjPL[i].cnt.Y += prjPL[i].y
				prjPL[i].r = RECCNT(prjPL[i].cnt, prjPL[i].r.Width, prjPL[i].r.Height)
			} else {
				a := prjENDANIM
				a.rD = RECCNT(prjPL[i].cnt, UNIT*2, UNIT*2)
				playonceANIMS = append(playonceANIMS, a)
				prjPL[i].off = true

			}
		}
	}
	return clear
}
func cMovePRJPL(p PROJ) bool {
	canmove := true
	for i := range lev.b {
		if lev.b[i].solid {
			if cIBPOINT(lev.b[i], p.cnt) {
				canmove = false
				break
			}
		}
	}
	return canmove
}
func uPL() {

	//SHOOT
	if MSL && pl.prjT == 0 {
		mPRJPL()
	}
	//TIMERS
	if pl.prjT > 0 {
		pl.prjT--
	}

	//INP
	move := z.NewVector2(0, 0)
	if kR {
		move.X += 1
	}
	if kL {
		move.X -= 1
	}
	if kU {
		move.Y -= 1
	}
	if kD {
		move.Y += 1
	}
	move = NORMALIZE(move)
	if move.X == 0 && move.Y == 0 {
		pl.st = 0
	} else {
		if cPLMOVE(move) {
			pl.st = 1
			pl.cnt.X += move.X * pl.spd
			pl.cnt.Y += move.Y * pl.spd
			pl.r = RECCNT(pl.cnt, pl.siz, pl.siz)
			pl.vb1 = z.NewVector2(pl.r.X, pl.r.Y+pl.r.Height)
			pl.vb2 = pl.vb1
			pl.vb2.X += pl.r.Width / 2
			pl.vb3 = pl.vb1
			pl.vb3.X += pl.r.Width
			num := cIGPOINT(lev, pl.vb2)
			if num != BLANKNUM {
				pl.zi = lev.b[num].zi
			}
			if move.X < 0 {
				pl.lr = true
			} else if move.X > 0 {
				pl.lr = false
			}
		} else {
			pl.st = 0
		}

	}
}
func cPLMOVE(move z.Vector2) bool {
	canmove := true
	cnt := pl.cnt
	cnt.X += move.X * pl.spd
	cnt.Y += move.Y * pl.spd
	r := RECCNT(cnt, pl.siz, pl.siz)
	vb1 := z.NewVector2(r.X, r.Y+r.Height)
	vb2 := vb1
	vb2.X += r.Width / 2
	vb3 := vb1
	vb3.X += r.Width

	for i := range lev.b {
		if lev.b[i].solid {
			if cIBPOINT(lev.b[i], vb1) || cIBPOINT(lev.b[i], vb2) || cIBPOINT(lev.b[i], vb3) {
				canmove = false
				break
			}
		}
	}
	return canmove
}

func UPGAME() {

	uPL()
	clear := false
	if len(prjPL) > 0 {
		clear = uPRJPL()
	}
	if clear {
		prjPL = REMPROJOFF(prjPL)
	}

	//MOUSE
	if MSL && clickT == 0 {
		clickT = clickP
	}

}
