package main

import z "github.com/gen2brain/raylib-go/raylib"

var (
	blokIM  []IM
	blankIM = IM{}
	lev     IG
	pl      PLAYER

	prjPL []PROJ

	prjENDANIM              ANIM
	prjANIMS, playonceANIMS []ANIM
)

type PLAYER struct {
	cnt, vb1, vb2, vb3 z.Vector2
	r                  z.Rectangle
	a                  ANIM
	spd, siz           float32
	st, zi, prjT       int
	lr                 bool
}

type PROJ struct {
	r         z.Rectangle
	x, y, max float32
	cnt       z.Vector2
	off       bool
}

func INITGAME() { //MARK: INIT GAME
	mLEV()
	mPL()
}
func mLEV() {
	lev = mISOGRIDcnt(CNT, UNIT, 20, 20, z.Magenta, z.Orange)
	lev = aIsoGridPerimeterIM(lev, blokIM[0], blokIM[0], blokIM[0], blokIM[0], true)
	lev = aIsoGridFloorIM(lev, blokIM[1])
	lev = aIsoGridRandomBloksIM(lev, blokIM[0], 50, true)
}
func mPRJPL() {
	pl.prjT = int(FPS) / 4
	siz := UNIT / 5
	spd := UNIT / 4
	p := PROJ{}
	p.max = spd
	if pl.lr {
		p.cnt = z.NewVector2(gunREC.X, gunREC.Y)
		p.cnt.X += gunREC.Width / 4
	} else {
		p.cnt = z.NewVector2(gunREC.X+gunREC.Width, gunREC.Y)
		p.cnt.X -= gunREC.Width / 4
	}
	if MS.Y > pl.cnt.Y {
		p.cnt.Y += gunREC.Height / 2
	}
	if MS.Y > pl.vb1.Y {
		p.cnt.Y += gunREC.Height / 2
	}
	//p.cnt = pl.cnt
	p.r = RECCNT(p.cnt, siz, siz)
	p.x, p.y = VELXY(p.cnt, MS, p.max)
	prjPL = append(prjPL, p)
}
func mPL() {
	pl.cnt = CNT
	pl.siz = UNIT + UNIT/2
	pl.r = RECCNT(pl.cnt, pl.siz, pl.siz)
	pl.vb1 = z.NewVector2(pl.r.X, pl.r.Y+pl.r.Height)
	pl.vb2 = pl.vb1
	pl.vb2.X += pl.r.Width / 2
	pl.vb3 = pl.vb1
	pl.vb3.X += pl.r.Width
	pl.zi = lev.b[cIGPOINT(lev, pl.vb2)].zi
	pl.a = mAnimSheetIM(ETC[1], 0, 0, 16, 0, 5, 1, 14)
	pl.spd = UNIT / 8

}

// MARK: MOVE

// MARK: MAKE

// UTILS
