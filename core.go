package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	FPS, SCRW, SCRH int32 = 60, 1920, 1080
	CNT             z.Vector2
	CAM             z.Camera2D
	UNIT            float32
	UNITDIVNUM      float32 = 24
	TEXRND          z.RenderTexture2D
	SHADER          z.Shader

	//OPTIONS
	FULLSCREEN, MUTEFX, MUTEMUSIC bool
)

func RESTART() {
	lev = IG{}
	pl = PLAYER{}
	mLEV()
	mPL()
}

func INITIAL() {
	//z.HideCursor()
	CNT = z.NewVector2(float32(SCRW/2), float32(SCRH/2))
	CAM.Target = CNT
	CAM.Offset = z.NewVector2(float32(SCRW/2), float32(SCRH/2))
	CAM.Rotation = 0.0
	CAM.Zoom = 1.0
	UNIT = float32(SCRH) / UNITDIVNUM
	TEXRND = z.LoadRenderTexture(SCRW*2, SCRH*2)
	SHADER = z.LoadShader("", "data/bloom.fs")
	//IM CORE
	//TILES = mIMSheetXY("im/tiles.png", 16, 16, 16, 8, 12, 20)
	//PATTERNS = mIMSheetFiles("im/patterns")
	ETC = mIMSheetFiles("im/etc")
	UI = mIMSheetFiles("im/ui")
	//IM GAME SPECIFIC
	blokIM = mIMSheetFiles("im/bloks")
	prjENDANIM = mAnimXY("im/prj/01.png", 0, 0, 64, 0, 9, 1, 20)
	prjANIMS = append(prjANIMS, mAnimXY("im/prj/02.png", 0, 0, 24, 0, 8, 1, 20))

	mFONTS()

	//GAME SPECIFIC INITIAL
	INITGAME()
}

func UNLOAD() {
	//CORE >> DO NOT DELETE
	if len(PATTERNS) > 0 {
		for i := range PATTERNS {
			z.UnloadTexture(PATTERNS[i].tex)
		}
	}
	if len(PATTERNS) > 0 {
		for i := range PATTERNS {
			z.UnloadTexture(PATTERNS[i].tex)
		}
	}
	if len(TILES) > 0 {
		for i := range TILES {
			z.UnloadTexture(TILES[i].tex)
		}
	}
	if len(ETC) > 0 {
		for i := range ETC {
			z.UnloadTexture(ETC[i].tex)
		}
	}
	if len(UI) > 0 {
		for i := range UI {
			z.UnloadTexture(UI[i].tex)
		}
	}
	if len(TEX) > 0 {
		for i := range TEX {
			z.UnloadTexture(TEX[i])
		}
	}

	z.UnloadFont(FONT1)
	z.UnloadFont(FONT2)
}
