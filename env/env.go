package main

import (
	"log"
	"runtime"
)

func main() {
	log.Printf("You are using %v on a %v machine\n", runtime.Compiler, runtime.GOARCH)
	log.Printf("Using Go version %v\n", runtime.Version())
	log.Printf("Number of CPUs: %v\n", runtime.NumCPU())
	log.Printf("Number of Goroutines: %v\n", runtime.NumGoroutine())

	/*
		stone@stonedeMacBook-Pro env % go run env.go
		2020/12/18 10:26:02 You are using gc on a amd64 machine
		2020/12/18 10:26:02 Using Go version go1.14.13
		2020/12/18 10:26:02 Number of CPUs: 8
		2020/12/18 10:26:02 Number of Goroutines: 1
	*/

	/*
		stone@stonedeMacBook-Pro env % GOOS=darwin GOARCH=amd64 go tool compile -S env.go
		"".main STEXT size=468 args=0x0 locals=0x60
		        0x0000 00000 (env.go:8) TEXT    "".main(SB), ABIInternal, $96-0
		        0x0000 00000 (env.go:8) MOVQ    (TLS), CX
		        0x0009 00009 (env.go:8) CMPQ    SP, 16(CX)
		        0x000d 00013 (env.go:8) PCDATA  $0, $-2
		        0x000d 00013 (env.go:8) JLS     458
		        0x0013 00019 (env.go:8) PCDATA  $0, $-1
		        0x0013 00019 (env.go:8) SUBQ    $96, SP
		        0x0017 00023 (env.go:8) MOVQ    BP, 88(SP)
		        0x001c 00028 (env.go:8) LEAQ    88(SP), BP
		        0x0021 00033 (env.go:8) PCDATA  $0, $-2
		        0x0021 00033 (env.go:8) PCDATA  $1, $-2
		        0x0021 00033 (env.go:8) FUNCDATA        $0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
		        0x0021 00033 (env.go:8) FUNCDATA        $1, gclocals·669b766e4959babb5d70e9d2b4843b8e(SB)
		        0x0021 00033 (env.go:8) FUNCDATA        $2, gclocals·6e8d7ea4abad763909b26991048ee1fe(SB)
		        0x0021 00033 (env.go:8) FUNCDATA        $3, "".main.stkobj(SB)
		        0x0021 00033 (env.go:9) PCDATA  $0, $0
		        0x0021 00033 (env.go:9) PCDATA  $1, $1
		        0x0021 00033 (env.go:9) XORPS   X0, X0
		        0x0024 00036 (env.go:9) MOVUPS  X0, ""..autotmp_3+56(SP)
		        0x0029 00041 (env.go:9) MOVUPS  X0, ""..autotmp_3+72(SP)
		        0x002e 00046 (env.go:9) PCDATA  $0, $1
		        0x002e 00046 (env.go:9) LEAQ    type.string(SB), AX
		        0x0035 00053 (env.go:9) MOVQ    AX, ""..autotmp_3+56(SP)
		        0x003a 00058 (env.go:9) PCDATA  $0, $2
		        0x003a 00058 (env.go:9) LEAQ    ""..stmp_0(SB), CX
		        0x0041 00065 (env.go:9) PCDATA  $0, $1
		        0x0041 00065 (env.go:9) MOVQ    CX, ""..autotmp_3+64(SP)
		        0x0046 00070 (env.go:9) PCDATA  $0, $0
		        0x0046 00070 (env.go:9) MOVQ    AX, ""..autotmp_3+72(SP)
		        0x004b 00075 (env.go:9) PCDATA  $0, $3
		        0x004b 00075 (env.go:9) LEAQ    ""..stmp_1(SB), CX
		        0x0052 00082 (env.go:9) PCDATA  $0, $0
		        0x0052 00082 (env.go:9) MOVQ    CX, ""..autotmp_3+80(SP)
		        0x0057 00087 (env.go:9) PCDATA  $0, $3
		        0x0057 00087 (env.go:9) LEAQ    go.string."You are using %v on a %v machine\n"(SB), CX
		        0x005e 00094 (env.go:9) PCDATA  $0, $0
		        0x005e 00094 (env.go:9) MOVQ    CX, (SP)
		        0x0062 00098 (env.go:9) MOVQ    $33, 8(SP)
		        0x006b 00107 (env.go:9) PCDATA  $0, $3
		        0x006b 00107 (env.go:9) PCDATA  $1, $0
		        0x006b 00107 (env.go:9) LEAQ    ""..autotmp_3+56(SP), CX
		        0x0070 00112 (env.go:9) PCDATA  $0, $0
		        0x0070 00112 (env.go:9) MOVQ    CX, 16(SP)
		        0x0075 00117 (env.go:9) MOVQ    $2, 24(SP)
		        0x007e 00126 (env.go:9) MOVQ    $2, 32(SP)
		        0x0087 00135 (env.go:9) CALL    log.Printf(SB)
		        0x008c 00140 (env.go:10)        PCDATA  $0, $1
		        0x008c 00140 (env.go:10)        LEAQ    go.string."go1.14.13"(SB), AX
		        0x0093 00147 (env.go:10)        PCDATA  $0, $0
		        0x0093 00147 (env.go:10)        MOVQ    AX, (SP)
		        0x0097 00151 (env.go:10)        MOVQ    $9, 8(SP)
		        0x00a0 00160 (env.go:10)        CALL    runtime.convTstring(SB)
		        0x00a5 00165 (env.go:10)        PCDATA  $0, $1
		        0x00a5 00165 (env.go:10)        MOVQ    16(SP), AX
		        0x00aa 00170 (env.go:10)        PCDATA  $1, $2
		        0x00aa 00170 (env.go:10)        XORPS   X0, X0
		        0x00ad 00173 (env.go:10)        MOVUPS  X0, ""..autotmp_4+40(SP)
		        0x00b2 00178 (env.go:10)        PCDATA  $0, $2
		        0x00b2 00178 (env.go:10)        LEAQ    type.string(SB), CX
		        0x00b9 00185 (env.go:10)        PCDATA  $0, $1
		        0x00b9 00185 (env.go:10)        MOVQ    CX, ""..autotmp_4+40(SP)
		        0x00be 00190 (env.go:10)        PCDATA  $0, $0
		        0x00be 00190 (env.go:10)        MOVQ    AX, ""..autotmp_4+48(SP)
		        0x00c3 00195 (env.go:10)        PCDATA  $0, $1
		        0x00c3 00195 (env.go:10)        LEAQ    go.string."Using Go version %v\n"(SB), AX
		        0x00ca 00202 (env.go:10)        PCDATA  $0, $0
		        0x00ca 00202 (env.go:10)        MOVQ    AX, (SP)
		        0x00ce 00206 (env.go:10)        MOVQ    $20, 8(SP)
		        0x00d7 00215 (env.go:10)        PCDATA  $0, $1
		        0x00d7 00215 (env.go:10)        PCDATA  $1, $0
		        0x00d7 00215 (env.go:10)        LEAQ    ""..autotmp_4+40(SP), AX
		        0x00dc 00220 (env.go:10)        PCDATA  $0, $0
		        0x00dc 00220 (env.go:10)        MOVQ    AX, 16(SP)
		        0x00e1 00225 (env.go:10)        MOVQ    $1, 24(SP)
		        0x00ea 00234 (env.go:10)        MOVQ    $1, 32(SP)
		        0x00f3 00243 (env.go:10)        CALL    log.Printf(SB)
		        0x00f8 00248 (<unknown line number>)    NOP
		        0x00f8 00248 ($GOROOT/src/runtime/debug.go:44)  MOVLQSX runtime.ncpu(SB), AX
		        0x00ff 00255 (env.go:11)        MOVQ    AX, (SP)
		        0x0103 00259 (env.go:11)        CALL    runtime.convT64(SB)
		        0x0108 00264 (env.go:11)        PCDATA  $0, $1
		        0x0108 00264 (env.go:11)        MOVQ    8(SP), AX
		        0x010d 00269 (env.go:11)        PCDATA  $1, $2
		        0x010d 00269 (env.go:11)        XORPS   X0, X0
		        0x0110 00272 (env.go:11)        MOVUPS  X0, ""..autotmp_4+40(SP)
		        0x0115 00277 (env.go:11)        PCDATA  $0, $2
		        0x0115 00277 (env.go:11)        LEAQ    type.int(SB), CX
		        0x011c 00284 (env.go:11)        PCDATA  $0, $1
		        0x011c 00284 (env.go:11)        MOVQ    CX, ""..autotmp_4+40(SP)
		        0x0121 00289 (env.go:11)        PCDATA  $0, $0
		        0x0121 00289 (env.go:11)        MOVQ    AX, ""..autotmp_4+48(SP)
		        0x0126 00294 (env.go:11)        PCDATA  $0, $1
		        0x0126 00294 (env.go:11)        LEAQ    go.string."Number of CPUs: %v\n"(SB), AX
		        0x012d 00301 (env.go:11)        PCDATA  $0, $0
		        0x012d 00301 (env.go:11)        MOVQ    AX, (SP)
		        0x0131 00305 (env.go:11)        MOVQ    $19, 8(SP)
		        0x013a 00314 (env.go:11)        PCDATA  $0, $1
		        0x013a 00314 (env.go:11)        PCDATA  $1, $0
		        0x013a 00314 (env.go:11)        LEAQ    ""..autotmp_4+40(SP), AX
		        0x013f 00319 (env.go:11)        PCDATA  $0, $0
		        0x013f 00319 (env.go:11)        MOVQ    AX, 16(SP)
		        0x0144 00324 (env.go:11)        MOVQ    $1, 24(SP)
		        0x014d 00333 (env.go:11)        MOVQ    $1, 32(SP)
		        0x0156 00342 (env.go:11)        CALL    log.Printf(SB)
		        0x015b 00347 (<unknown line number>)    NOP
		        0x015b 00347 ($GOROOT/src/runtime/debug.go:58)  CALL    runtime.gcount(SB)
		        0x0160 00352 ($GOROOT/src/runtime/debug.go:58)  MOVLQSX (SP), AX
		        0x0164 00356 (env.go:12)        MOVQ    AX, (SP)
		        0x0168 00360 (env.go:12)        CALL    runtime.convT64(SB)
		        0x016d 00365 (env.go:12)        PCDATA  $0, $1
		        0x016d 00365 (env.go:12)        MOVQ    8(SP), AX
		        0x0172 00370 (env.go:12)        PCDATA  $1, $2
		        0x0172 00370 (env.go:12)        XORPS   X0, X0
		        0x0175 00373 (env.go:12)        MOVUPS  X0, ""..autotmp_4+40(SP)
		        0x017a 00378 (env.go:12)        PCDATA  $0, $2
		        0x017a 00378 (env.go:12)        LEAQ    type.int(SB), CX
		        0x0181 00385 (env.go:12)        PCDATA  $0, $1
		        0x0181 00385 (env.go:12)        MOVQ    CX, ""..autotmp_4+40(SP)
		        0x0186 00390 (env.go:12)        PCDATA  $0, $0
		        0x0186 00390 (env.go:12)        MOVQ    AX, ""..autotmp_4+48(SP)
		        0x018b 00395 (env.go:12)        PCDATA  $0, $1
		        0x018b 00395 (env.go:12)        LEAQ    go.string."Number of Goroutines: %v\n"(SB), AX
		        0x0192 00402 (env.go:12)        PCDATA  $0, $0
		        0x0192 00402 (env.go:12)        MOVQ    AX, (SP)
		        0x0196 00406 (env.go:12)        MOVQ    $25, 8(SP)
		        0x019f 00415 (env.go:12)        PCDATA  $0, $1
		        0x019f 00415 (env.go:12)        PCDATA  $1, $0
		        0x019f 00415 (env.go:12)        LEAQ    ""..autotmp_4+40(SP), AX
		        0x01a4 00420 (env.go:12)        PCDATA  $0, $0
		        0x01a4 00420 (env.go:12)        MOVQ    AX, 16(SP)
		        0x01a9 00425 (env.go:12)        MOVQ    $1, 24(SP)
		        0x01b2 00434 (env.go:12)        MOVQ    $1, 32(SP)
		        0x01bb 00443 (env.go:12)        CALL    log.Printf(SB)
		        0x01c0 00448 (env.go:13)        MOVQ    88(SP), BP
		        0x01c5 00453 (env.go:13)        ADDQ    $96, SP
		        0x01c9 00457 (env.go:13)        RET
		        0x01ca 00458 (env.go:13)        NOP
		        0x01ca 00458 (env.go:8) PCDATA  $1, $-1
		        0x01ca 00458 (env.go:8) PCDATA  $0, $-2
		        0x01ca 00458 (env.go:8) CALL    runtime.morestack_noctxt(SB)
		        0x01cf 00463 (env.go:8) PCDATA  $0, $-1
		        0x01cf 00463 (env.go:8) JMP     0
		        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 b7  eH..%....H;a....
		        0x0010 01 00 00 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24  ...H..`H.l$XH.l$
		        0x0020 58 0f 57 c0 0f 11 44 24 38 0f 11 44 24 48 48 8d  X.W...D$8..D$HH.
		        0x0030 05 00 00 00 00 48 89 44 24 38 48 8d 0d 00 00 00  .....H.D$8H.....
		        0x0040 00 48 89 4c 24 40 48 89 44 24 48 48 8d 0d 00 00  .H.L$@H.D$HH....
		        0x0050 00 00 48 89 4c 24 50 48 8d 0d 00 00 00 00 48 89  ..H.L$PH......H.
		        0x0060 0c 24 48 c7 44 24 08 21 00 00 00 48 8d 4c 24 38  .$H.D$.!...H.L$8
		        0x0070 48 89 4c 24 10 48 c7 44 24 18 02 00 00 00 48 c7  H.L$.H.D$.....H.
		        0x0080 44 24 20 02 00 00 00 e8 00 00 00 00 48 8d 05 00  D$ .........H...
		        0x0090 00 00 00 48 89 04 24 48 c7 44 24 08 09 00 00 00  ...H..$H.D$.....
		        0x00a0 e8 00 00 00 00 48 8b 44 24 10 0f 57 c0 0f 11 44  .....H.D$..W...D
		        0x00b0 24 28 48 8d 0d 00 00 00 00 48 89 4c 24 28 48 89  $(H......H.L$(H.
		        0x00c0 44 24 30 48 8d 05 00 00 00 00 48 89 04 24 48 c7  D$0H......H..$H.
		        0x00d0 44 24 08 14 00 00 00 48 8d 44 24 28 48 89 44 24  D$.....H.D$(H.D$
		        0x00e0 10 48 c7 44 24 18 01 00 00 00 48 c7 44 24 20 01  .H.D$.....H.D$ .
		        0x00f0 00 00 00 e8 00 00 00 00 48 63 05 00 00 00 00 48  ........Hc.....H
		        0x0100 89 04 24 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0  ..$.....H.D$..W.
		        0x0110 0f 11 44 24 28 48 8d 0d 00 00 00 00 48 89 4c 24  ..D$(H......H.L$
		        0x0120 28 48 89 44 24 30 48 8d 05 00 00 00 00 48 89 04  (H.D$0H......H..
		        0x0130 24 48 c7 44 24 08 13 00 00 00 48 8d 44 24 28 48  $H.D$.....H.D$(H
		        0x0140 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44  .D$.H.D$.....H.D
		        0x0150 24 20 01 00 00 00 e8 00 00 00 00 e8 00 00 00 00  $ ..............
		        0x0160 48 63 04 24 48 89 04 24 e8 00 00 00 00 48 8b 44  Hc.$H..$.....H.D
		        0x0170 24 08 0f 57 c0 0f 11 44 24 28 48 8d 0d 00 00 00  $..W...D$(H.....
		        0x0180 00 48 89 4c 24 28 48 89 44 24 30 48 8d 05 00 00  .H.L$(H.D$0H....
		        0x0190 00 00 48 89 04 24 48 c7 44 24 08 19 00 00 00 48  ..H..$H.D$.....H
		        0x01a0 8d 44 24 28 48 89 44 24 10 48 c7 44 24 18 01 00  .D$(H.D$.H.D$...
		        0x01b0 00 00 48 c7 44 24 20 01 00 00 00 e8 00 00 00 00  ..H.D$ .........
		        0x01c0 48 8b 6c 24 58 48 83 c4 60 c3 e8 00 00 00 00 e9  H.l$XH..`.......
		        0x01d0 2c fe ff ff                                      ,...
		        rel 5+4 t=17 TLS+0
		        rel 49+4 t=16 type.string+0
		        rel 61+4 t=16 ""..stmp_0+0
		        rel 78+4 t=16 ""..stmp_1+0
		        rel 90+4 t=16 go.string."You are using %v on a %v machine\n"+0
		        rel 136+4 t=8 log.Printf+0
		        rel 143+4 t=16 go.string."go1.14.13"+0
		        rel 161+4 t=8 runtime.convTstring+0
		        rel 181+4 t=16 type.string+0
		        rel 198+4 t=16 go.string."Using Go version %v\n"+0
		        rel 244+4 t=8 log.Printf+0
		        rel 251+4 t=16 runtime.ncpu+0
		        rel 260+4 t=8 runtime.convT64+0
		        rel 280+4 t=16 type.int+0
		        rel 297+4 t=16 go.string."Number of CPUs: %v\n"+0
		        rel 343+4 t=8 log.Printf+0
		        rel 348+4 t=8 runtime.gcount+0
		        rel 361+4 t=8 runtime.convT64+0
		        rel 381+4 t=16 type.int+0
		        rel 398+4 t=16 go.string."Number of Goroutines: %v\n"+0
		        rel 444+4 t=8 log.Printf+0
		        rel 459+4 t=8 runtime.morestack_noctxt+0
		type..eq.[2]interface {} STEXT dupok size=179 args=0x18 locals=0x30
		        0x0000 00000 (<autogenerated>:1)        TEXT    type..eq.[2]interface {}(SB), DUPOK|ABIInternal, $48-24
		        0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
		        0x0009 00009 (<autogenerated>:1)        CMPQ    SP, 16(CX)
		        0x000d 00013 (<autogenerated>:1)        PCDATA  $0, $-2
		        0x000d 00013 (<autogenerated>:1)        JLS     169
		        0x0013 00019 (<autogenerated>:1)        PCDATA  $0, $-1
		        0x0013 00019 (<autogenerated>:1)        SUBQ    $48, SP
		        0x0017 00023 (<autogenerated>:1)        MOVQ    BP, 40(SP)
		        0x001c 00028 (<autogenerated>:1)        LEAQ    40(SP), BP
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $0, $-2
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $1, $-2
		        0x0021 00033 (<autogenerated>:1)        FUNCDATA        $0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
		        0x0021 00033 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
		        0x0021 00033 (<autogenerated>:1)        FUNCDATA        $2, gclocals·313a5bdbfadc4f007c002a3a3588596d(SB)
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $0, $1
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $1, $0
		        0x0021 00033 (<autogenerated>:1)        MOVQ    "".p+56(SP), AX
		        0x0026 00038 (<autogenerated>:1)        PCDATA  $0, $2
		        0x0026 00038 (<autogenerated>:1)        MOVQ    "".q+64(SP), CX
		        0x002b 00043 (<autogenerated>:1)        XORL    DX, DX
		        0x002d 00045 (<autogenerated>:1)        JMP     72
		        0x002f 00047 (<autogenerated>:1)        PCDATA  $0, $0
		        0x002f 00047 (<autogenerated>:1)        MOVQ    ""..autotmp_8+32(SP), BX
		        0x0034 00052 (<autogenerated>:1)        LEAQ    1(BX), DX
		        0x0038 00056 (<autogenerated>:1)        PCDATA  $0, $3
		        0x0038 00056 (<autogenerated>:1)        MOVQ    "".p+56(SP), BX
		        0x003d 00061 (<autogenerated>:1)        PCDATA  $0, $4
		        0x003d 00061 (<autogenerated>:1)        MOVQ    "".q+64(SP), SI
		        0x0042 00066 (<autogenerated>:1)        PCDATA  $0, $5
		        0x0042 00066 (<autogenerated>:1)        MOVQ    BX, AX
		        0x0045 00069 (<autogenerated>:1)        PCDATA  $0, $2
		        0x0045 00069 (<autogenerated>:1)        MOVQ    SI, CX
		        0x0048 00072 (<autogenerated>:1)        CMPQ    DX, $2
		        0x004c 00076 (<autogenerated>:1)        JGE     154
		        0x004e 00078 (<autogenerated>:1)        MOVQ    DX, BX
		        0x0051 00081 (<autogenerated>:1)        SHLQ    $4, DX
		        0x0055 00085 (<autogenerated>:1)        PCDATA  $0, $6
		        0x0055 00085 (<autogenerated>:1)        MOVQ    8(DX)(AX*1), SI
		        0x005a 00090 (<autogenerated>:1)        PCDATA  $0, $7
		        0x005a 00090 (<autogenerated>:1)        MOVQ    (DX)(AX*1), DI
		        0x005e 00094 (<autogenerated>:1)        MOVQ    (DX)(CX*1), R8
		        0x0062 00098 (<autogenerated>:1)        PCDATA  $0, $8
		        0x0062 00098 (<autogenerated>:1)        MOVQ    8(DX)(CX*1), DX
		        0x0067 00103 (<autogenerated>:1)        CMPQ    DI, R8
		        0x006a 00106 (<autogenerated>:1)        JNE     139
		        0x006c 00108 (<autogenerated>:1)        MOVQ    BX, ""..autotmp_8+32(SP)
		        0x0071 00113 (<autogenerated>:1)        MOVQ    DI, (SP)
		        0x0075 00117 (<autogenerated>:1)        PCDATA  $0, $9
		        0x0075 00117 (<autogenerated>:1)        MOVQ    SI, 8(SP)
		        0x007a 00122 (<autogenerated>:1)        PCDATA  $0, $0
		        0x007a 00122 (<autogenerated>:1)        MOVQ    DX, 16(SP)
		        0x007f 00127 (<autogenerated>:1)        CALL    runtime.efaceeq(SB)
		        0x0084 00132 (<autogenerated>:1)        CMPB    24(SP), $0
		        0x0089 00137 (<autogenerated>:1)        JNE     47
		        0x008b 00139 (<autogenerated>:1)        PCDATA  $1, $1
		        0x008b 00139 (<autogenerated>:1)        MOVB    $0, "".~r2+72(SP)
		        0x0090 00144 (<autogenerated>:1)        MOVQ    40(SP), BP
		        0x0095 00149 (<autogenerated>:1)        ADDQ    $48, SP
		        0x0099 00153 (<autogenerated>:1)        RET
		        0x009a 00154 (<autogenerated>:1)        MOVB    $1, "".~r2+72(SP)
		        0x009f 00159 (<autogenerated>:1)        MOVQ    40(SP), BP
		        0x00a4 00164 (<autogenerated>:1)        ADDQ    $48, SP
		        0x00a8 00168 (<autogenerated>:1)        RET
		        0x00a9 00169 (<autogenerated>:1)        NOP
		        0x00a9 00169 (<autogenerated>:1)        PCDATA  $1, $-1
		        0x00a9 00169 (<autogenerated>:1)        PCDATA  $0, $-2
		        0x00a9 00169 (<autogenerated>:1)        CALL    runtime.morestack_noctxt(SB)
		        0x00ae 00174 (<autogenerated>:1)        PCDATA  $0, $-1
		        0x00ae 00174 (<autogenerated>:1)        JMP     0
		        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 96  eH..%....H;a....
		        0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
		        0x0020 28 48 8b 44 24 38 48 8b 4c 24 40 31 d2 eb 19 48  (H.D$8H.L$@1...H
		        0x0030 8b 5c 24 20 48 8d 53 01 48 8b 5c 24 38 48 8b 74  .\$ H.S.H.\$8H.t
		        0x0040 24 40 48 89 d8 48 89 f1 48 83 fa 02 7d 4c 48 89  $@H..H..H...}LH.
		        0x0050 d3 48 c1 e2 04 48 8b 74 02 08 48 8b 3c 02 4c 8b  .H...H.t..H.<.L.
		        0x0060 04 0a 48 8b 54 0a 08 4c 39 c7 75 1f 48 89 5c 24  ..H.T..L9.u.H.\$
		        0x0070 20 48 89 3c 24 48 89 74 24 08 48 89 54 24 10 e8   H.<$H.t$.H.T$..
		        0x0080 00 00 00 00 80 7c 24 18 00 75 a4 c6 44 24 48 00  .....|$..u..D$H.
		        0x0090 48 8b 6c 24 28 48 83 c4 30 c3 c6 44 24 48 01 48  H.l$(H..0..D$H.H
		        0x00a0 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00 00 e9 4d  .l$(H..0.......M
		        0x00b0 ff ff ff                                         ...
		        rel 5+4 t=17 TLS+0
		        rel 128+4 t=8 runtime.efaceeq+0
		        rel 170+4 t=8 runtime.morestack_noctxt+0
		go.cuinfo.packagename. SDWARFINFO dupok size=0
		        0x0000 6d 61 69 6e                                      main
		go.info.runtime.Version$abstract SDWARFINFO dupok size=20
		        0x0000 04 72 75 6e 74 69 6d 65 2e 56 65 72 73 69 6f 6e  .runtime.Version
		        0x0010 00 01 01 00                                      ....
		go.info.runtime.NumCPU$abstract SDWARFINFO dupok size=19
		        0x0000 04 72 75 6e 74 69 6d 65 2e 4e 75 6d 43 50 55 00  .runtime.NumCPU.
		        0x0010 01 01 00                                         ...
		go.info.runtime.NumGoroutine$abstract SDWARFINFO dupok size=25
		        0x0000 04 72 75 6e 74 69 6d 65 2e 4e 75 6d 47 6f 72 6f  .runtime.NumGoro
		        0x0010 75 74 69 6e 65 00 01 01 00                       utine....
		go.string."gc" SRODATA dupok size=2
		        0x0000 67 63                                            gc
		go.string."amd64" SRODATA dupok size=5
		        0x0000 61 6d 64 36 34                                   amd64
		go.string."You are using %v on a %v machine\n" SRODATA dupok size=33
		        0x0000 59 6f 75 20 61 72 65 20 75 73 69 6e 67 20 25 76  You are using %v
		        0x0010 20 6f 6e 20 61 20 25 76 20 6d 61 63 68 69 6e 65   on a %v machine
		        0x0020 0a                                               .
		go.string."go1.14.13" SRODATA dupok size=9
		        0x0000 67 6f 31 2e 31 34 2e 31 33                       go1.14.13
		go.string."Using Go version %v\n" SRODATA dupok size=20
		        0x0000 55 73 69 6e 67 20 47 6f 20 76 65 72 73 69 6f 6e  Using Go version
		        0x0010 20 25 76 0a                                       %v.
		go.string."Number of CPUs: %v\n" SRODATA dupok size=19
		        0x0000 4e 75 6d 62 65 72 20 6f 66 20 43 50 55 73 3a 20  Number of CPUs:
		        0x0010 25 76 0a                                         %v.
		go.string."Number of Goroutines: %v\n" SRODATA dupok size=25
		        0x0000 4e 75 6d 62 65 72 20 6f 66 20 47 6f 72 6f 75 74  Number of Gorout
		        0x0010 69 6e 65 73 3a 20 25 76 0a                       ines: %v.
		go.loc."".main SDWARFLOC size=0
		go.info."".main SDWARFINFO size=93
		        0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
		        0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
		        0x0020 06 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 0b 00 00 00 00 06 00  ................
		        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0050 00 00 00 00 00 00 00 0c 00 00 00 00 00           .............
		        rel 0+0 t=24 type.[1]interface {}+0
		        rel 0+0 t=24 type.[2]interface {}+0
		        rel 9+8 t=1 "".main+0
		        rel 17+8 t=1 "".main+468
		        rel 27+4 t=30 gofile../Users/stone/Projects/go/stone/mastering-go/env/env.go+0
		        rel 33+4 t=29 go.info.runtime.NumCPU$abstract+0
		        rel 37+8 t=1 "".main+248
		        rel 45+8 t=1 "".main+255
		        rel 53+4 t=30 gofile../Users/stone/Projects/go/stone/mastering-go/env/env.go+0
		        rel 63+4 t=29 go.info.runtime.NumGoroutine$abstract+0
		        rel 67+8 t=1 "".main+347
		        rel 75+8 t=1 "".main+356
		        rel 83+4 t=30 gofile../Users/stone/Projects/go/stone/mastering-go/env/env.go+0
		go.range."".main SDWARFRANGE size=0
		go.debuglines."".main SDWARFMISC size=65
		        0x0000 04 02 03 02 14 0a cd 9c 06 2d 06 02 51 f6 06 55  .........-..Q..U
		        0x0010 06 91 06 41 04 06 06 02 3c 03 1d fa 04 02 03 63  ...A....<......c
		        0x0020 51 06 69 04 06 06 02 3c 03 2a fa 06 41 04 02 06  Q.i....<.*..A...
		        0x0030 03 56 33 06 69 06 02 3c f6 03 7f 6f 04 01 03 79  .V3.i..<...o...y
		        0x0040 01                                               .
		runtime.nilinterequal·f SRODATA dupok size=8
		        0x0000 00 00 00 00 00 00 00 00                          ........
		        rel 0+8 t=1 runtime.nilinterequal+0
		runtime.memequal64·f SRODATA dupok size=8
		        0x0000 00 00 00 00 00 00 00 00                          ........
		        rel 0+8 t=1 runtime.memequal64+0
		runtime.gcbits.01 SRODATA dupok size=1
		        0x0000 01                                               .
		type..namedata.*interface {}- SRODATA dupok size=16
		        0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
		type.*interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*interface {}-+0
		        rel 48+8 t=1 type.interface {}+0
		runtime.gcbits.02 SRODATA dupok size=1
		        0x0000 02                                               .
		type.interface {} SRODATA dupok size=80
		        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
		        0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        rel 24+8 t=1 runtime.nilinterequal·f+0
		        rel 32+8 t=1 runtime.gcbits.02+0
		        rel 40+4 t=5 type..namedata.*interface {}-+0
		        rel 44+4 t=6 type.*interface {}+0
		        rel 56+8 t=1 type.interface {}+80
		type..namedata.*[]interface {}- SRODATA dupok size=18
		        0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface
		        0x0010 7b 7d                                            {}
		type.*[]interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[]interface {}-+0
		        rel 48+8 t=1 type.[]interface {}+0
		type.[]interface {} SRODATA dupok size=56
		        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[]interface {}-+0
		        rel 44+4 t=6 type.*[]interface {}+0
		        rel 48+8 t=1 type.interface {}+0
		go.loc.type..eq.[2]interface {} SDWARFLOC dupok size=106
		        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0010 01 00 51 00 00 00 00 00 00 00 00 00 00 00 00 00  ..Q.............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 01 00 9c 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0050 00 00 00 00 00 00 02 00 91 08 00 00 00 00 00 00  ................
		        0x0060 00 00 00 00 00 00 00 00 00 00                    ..........
		        rel 0+8 t=53 type..eq.[2]interface {}+72
		        rel 8+8 t=53 type..eq.[2]interface {}+85
		        rel 35+8 t=53 type..eq.[2]interface {}+0
		        rel 43+8 t=53 type..eq.[2]interface {}+179
		        rel 70+8 t=53 type..eq.[2]interface {}+0
		        rel 78+8 t=53 type..eq.[2]interface {}+179
		go.info.type..eq.[2]interface {} SDWARFINFO dupok size=100
		        0x0000 03 74 79 70 65 2e 2e 65 71 2e 5b 32 5d 69 6e 74  .type..eq.[2]int
		        0x0010 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00 00 00  erface {}.......
		        0x0020 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
		        0x0030 01 0b 69 00 01 00 00 00 00 00 00 00 00 10 70 00  ..i...........p.
		        0x0040 00 01 00 00 00 00 00 00 00 00 10 71 00 00 01 00  ...........q....
		        0x0050 00 00 00 00 00 00 00 0f 7e 72 32 00 01 01 00 00  ........~r2.....
		        0x0060 00 00 00 00                                      ....
		        rel 0+0 t=24 type.*[2]interface {}+0
		        rel 0+0 t=24 type.bool+0
		        rel 0+0 t=24 type.int+0
		        rel 26+8 t=1 type..eq.[2]interface {}+0
		        rel 34+8 t=1 type..eq.[2]interface {}+179
		        rel 44+4 t=30 gofile..<autogenerated>+0
		        rel 53+4 t=29 go.info.int+0
		        rel 57+4 t=29 go.loc.type..eq.[2]interface {}+0
		        rel 66+4 t=29 go.info.*[2]interface {}+0
		        rel 70+4 t=29 go.loc.type..eq.[2]interface {}+35
		        rel 79+4 t=29 go.info.*[2]interface {}+0
		        rel 83+4 t=29 go.loc.type..eq.[2]interface {}+70
		        rel 94+4 t=29 go.info.bool+0
		go.range.type..eq.[2]interface {} SDWARFRANGE dupok size=0
		go.debuglines.type..eq.[2]interface {} SDWARFMISC dupok size=29
		        0x0000 04 01 0f 0a cd 06 cd 06 08 73 06 37 06 02 27 ff  .........s.7..'.
		        0x0010 06 41 06 73 06 41 06 73 04 01 03 00 01           .A.s.A.s.....
		type..eqfunc.[2]interface {} SRODATA dupok size=8
		        0x0000 00 00 00 00 00 00 00 00                          ........
		        rel 0+8 t=1 type..eq.[2]interface {}+0
		type..namedata.*[2]interface {}- SRODATA dupok size=19
		        0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
		        0x0010 20 7b 7d                                          {}
		type.*[2]interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 be 73 2d 71 08 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[2]interface {}-+0
		        rel 48+8 t=1 type.[2]interface {}+0
		runtime.gcbits.0a SRODATA dupok size=1
		        0x0000 0a                                               .
		type.[2]interface {} SRODATA dupok size=72
		        0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
		        0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 02 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 type..eqfunc.[2]interface {}+0
		        rel 32+8 t=1 runtime.gcbits.0a+0
		        rel 40+4 t=5 type..namedata.*[2]interface {}-+0
		        rel 44+4 t=6 type.*[2]interface {}+0
		        rel 48+8 t=1 type.interface {}+0
		        rel 56+8 t=1 type.[]interface {}+0
		type..namedata.*[1]interface {}- SRODATA dupok size=19
		        0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
		        0x0010 20 7b 7d                                          {}
		type.*[1]interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
		        rel 48+8 t=1 type.[1]interface {}+0
		type.[1]interface {} SRODATA dupok size=72
		        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
		        0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 01 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.nilinterequal·f+0
		        rel 32+8 t=1 runtime.gcbits.02+0
		        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
		        rel 44+4 t=6 type.*[1]interface {}+0
		        rel 48+8 t=1 type.interface {}+0
		        rel 56+8 t=1 type.[]interface {}+0
		""..inittask SNOPTRDATA size=40
		        0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
		        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0020 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 log..inittask+0
		        rel 32+8 t=1 runtime..inittask+0
		""..stmp_0 SRODATA size=16
		        0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
		        rel 0+8 t=1 go.string."gc"+0
		""..stmp_1 SRODATA size=16
		        0x0000 00 00 00 00 00 00 00 00 05 00 00 00 00 00 00 00  ................
		        rel 0+8 t=1 go.string."amd64"+0
		type..importpath.log. SRODATA dupok size=6
		        0x0000 00 00 03 6c 6f 67                                ...log
		type..importpath.runtime. SRODATA dupok size=10
		        0x0000 00 00 07 72 75 6e 74 69 6d 65                    ...runtime
		gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
		        0x0000 03 00 00 00 00 00 00 00                          ........
		gclocals·669b766e4959babb5d70e9d2b4843b8e SRODATA dupok size=11
		        0x0000 03 00 00 00 06 00 00 00 00 28 02                 .........(.
		gclocals·6e8d7ea4abad763909b26991048ee1fe SRODATA dupok size=12
		        0x0000 04 00 00 00 02 00 00 00 00 01 03 02              ............
		"".main.stkobj SRODATA size=40
		        0x0000 02 00 00 00 00 00 00 00 d0 ff ff ff ff ff ff ff  ................
		        0x0010 00 00 00 00 00 00 00 00 e0 ff ff ff ff ff ff ff  ................
		        0x0020 00 00 00 00 00 00 00 00                          ........
		        rel 16+8 t=1 type.[1]interface {}+0
		        rel 32+8 t=1 type.[2]interface {}+0
		gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
		        0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
		gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
		        0x0000 02 00 00 00 00 00 00 00                          ........
		gclocals·313a5bdbfadc4f007c002a3a3588596d SRODATA dupok size=18
		        0x0000 0a 00 00 00 06 00 00 00 00 01 03 08 28 21 23 22  ............(!#"
		        0x0010 24 04                                            $.

	*/

	/*
		stone@stonedeMacBook-Pro env % GOOS=darwin GOARCH=amd64 go build -gcflags -S env.go
		# command-line-arguments
		"".main STEXT size=468 args=0x0 locals=0x60
		        0x0000 00000 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) TEXT    "".main(SB), ABIInternal, $96-0
		        0x0000 00000 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) MOVQ    (TLS), CX
		        0x0009 00009 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) CMPQ    SP, 16(CX)
		        0x000d 00013 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $0, $-2
		        0x000d 00013 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) JLS     458
		        0x0013 00019 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $0, $-1
		        0x0013 00019 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) SUBQ    $96, SP
		        0x0017 00023 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) MOVQ    BP, 88(SP)
		        0x001c 00028 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) LEAQ    88(SP), BP
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $0, $-2
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $1, $-2
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) FUNCDATA        $0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) FUNCDATA        $1, gclocals·669b766e4959babb5d70e9d2b4843b8e(SB)
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) FUNCDATA        $2, gclocals·6e8d7ea4abad763909b26991048ee1fe(SB)
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) FUNCDATA        $3, "".main.stkobj(SB)
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $0
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $1, $1
		        0x0021 00033 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) XORPS   X0, X0
		        0x0024 00036 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVUPS  X0, ""..autotmp_3+56(SP)
		        0x0029 00041 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVUPS  X0, ""..autotmp_3+72(SP)
		        0x002e 00046 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $1
		        0x002e 00046 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) LEAQ    type.string(SB), AX
		        0x0035 00053 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    AX, ""..autotmp_3+56(SP)
		        0x003a 00058 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $2
		        0x003a 00058 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) LEAQ    ""..stmp_0(SB), CX
		        0x0041 00065 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $1
		        0x0041 00065 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    CX, ""..autotmp_3+64(SP)
		        0x0046 00070 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $0
		        0x0046 00070 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    AX, ""..autotmp_3+72(SP)
		        0x004b 00075 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $3
		        0x004b 00075 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) LEAQ    ""..stmp_1(SB), CX
		        0x0052 00082 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $0
		        0x0052 00082 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    CX, ""..autotmp_3+80(SP)
		        0x0057 00087 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $3
		        0x0057 00087 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) LEAQ    go.string."You are using %v on a %v machine\n"(SB), CX
		        0x005e 00094 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $0
		        0x005e 00094 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    CX, (SP)
		        0x0062 00098 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    $33, 8(SP)
		        0x006b 00107 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $3
		        0x006b 00107 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $1, $0
		        0x006b 00107 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) LEAQ    ""..autotmp_3+56(SP), CX
		        0x0070 00112 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) PCDATA  $0, $0
		        0x0070 00112 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    CX, 16(SP)
		        0x0075 00117 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    $2, 24(SP)
		        0x007e 00126 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) MOVQ    $2, 32(SP)
		        0x0087 00135 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:9) CALL    log.Printf(SB)
		        0x008c 00140 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $1
		        0x008c 00140 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        LEAQ    go.string."go1.14.13"(SB), AX
		        0x0093 00147 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $0
		        0x0093 00147 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    AX, (SP)
		        0x0097 00151 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    $9, 8(SP)
		        0x00a0 00160 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        CALL    runtime.convTstring(SB)
		        0x00a5 00165 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $1
		        0x00a5 00165 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    16(SP), AX
		        0x00aa 00170 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $1, $2
		        0x00aa 00170 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        XORPS   X0, X0
		        0x00ad 00173 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVUPS  X0, ""..autotmp_4+40(SP)
		        0x00b2 00178 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $2
		        0x00b2 00178 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        LEAQ    type.string(SB), CX
		        0x00b9 00185 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $1
		        0x00b9 00185 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    CX, ""..autotmp_4+40(SP)
		        0x00be 00190 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $0
		        0x00be 00190 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    AX, ""..autotmp_4+48(SP)
		        0x00c3 00195 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $1
		        0x00c3 00195 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        LEAQ    go.string."Using Go version %v\n"(SB), AX
		        0x00ca 00202 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $0
		        0x00ca 00202 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    AX, (SP)
		        0x00ce 00206 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    $20, 8(SP)
		        0x00d7 00215 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $1
		        0x00d7 00215 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $1, $0
		        0x00d7 00215 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        LEAQ    ""..autotmp_4+40(SP), AX
		        0x00dc 00220 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        PCDATA  $0, $0
		        0x00dc 00220 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    AX, 16(SP)
		        0x00e1 00225 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    $1, 24(SP)
		        0x00ea 00234 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        MOVQ    $1, 32(SP)
		        0x00f3 00243 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:10)        CALL    log.Printf(SB)
		        0x00f8 00248 (<unknown line number>)    NOP
		        0x00f8 00248 ($GOROOT/src/runtime/debug.go:44)  MOVLQSX runtime.ncpu(SB), AX
		        0x00ff 00255 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    AX, (SP)
		        0x0103 00259 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        CALL    runtime.convT64(SB)
		        0x0108 00264 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $1
		        0x0108 00264 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    8(SP), AX
		        0x010d 00269 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $1, $2
		        0x010d 00269 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        XORPS   X0, X0
		        0x0110 00272 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVUPS  X0, ""..autotmp_4+40(SP)
		        0x0115 00277 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $2
		        0x0115 00277 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        LEAQ    type.int(SB), CX
		        0x011c 00284 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $1
		        0x011c 00284 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    CX, ""..autotmp_4+40(SP)
		        0x0121 00289 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $0
		        0x0121 00289 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    AX, ""..autotmp_4+48(SP)
		        0x0126 00294 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $1
		        0x0126 00294 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        LEAQ    go.string."Number of CPUs: %v\n"(SB), AX
		        0x012d 00301 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $0
		        0x012d 00301 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    AX, (SP)
		        0x0131 00305 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    $19, 8(SP)
		        0x013a 00314 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $1
		        0x013a 00314 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $1, $0
		        0x013a 00314 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        LEAQ    ""..autotmp_4+40(SP), AX
		        0x013f 00319 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        PCDATA  $0, $0
		        0x013f 00319 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    AX, 16(SP)
		        0x0144 00324 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    $1, 24(SP)
		        0x014d 00333 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        MOVQ    $1, 32(SP)
		        0x0156 00342 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:11)        CALL    log.Printf(SB)
		        0x015b 00347 (<unknown line number>)    NOP
		        0x015b 00347 ($GOROOT/src/runtime/debug.go:58)  CALL    runtime.gcount(SB)
		        0x0160 00352 ($GOROOT/src/runtime/debug.go:58)  MOVLQSX (SP), AX
		        0x0164 00356 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    AX, (SP)
		        0x0168 00360 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        CALL    runtime.convT64(SB)
		        0x016d 00365 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $1
		        0x016d 00365 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    8(SP), AX
		        0x0172 00370 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $1, $2
		        0x0172 00370 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        XORPS   X0, X0
		        0x0175 00373 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVUPS  X0, ""..autotmp_4+40(SP)
		        0x017a 00378 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $2
		        0x017a 00378 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        LEAQ    type.int(SB), CX
		        0x0181 00385 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $1
		        0x0181 00385 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    CX, ""..autotmp_4+40(SP)
		        0x0186 00390 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $0
		        0x0186 00390 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    AX, ""..autotmp_4+48(SP)
		        0x018b 00395 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $1
		        0x018b 00395 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        LEAQ    go.string."Number of Goroutines: %v\n"(SB), AX
		        0x0192 00402 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $0
		        0x0192 00402 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    AX, (SP)
		        0x0196 00406 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    $25, 8(SP)
		        0x019f 00415 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $1
		        0x019f 00415 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $1, $0
		        0x019f 00415 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        LEAQ    ""..autotmp_4+40(SP), AX
		        0x01a4 00420 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        PCDATA  $0, $0
		        0x01a4 00420 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    AX, 16(SP)
		        0x01a9 00425 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    $1, 24(SP)
		        0x01b2 00434 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        MOVQ    $1, 32(SP)
		        0x01bb 00443 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:12)        CALL    log.Printf(SB)
		        0x01c0 00448 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:544)       MOVQ    88(SP), BP
		        0x01c5 00453 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:544)       ADDQ    $96, SP
		        0x01c9 00457 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:544)       RET
		        0x01ca 00458 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:544)       NOP
		        0x01ca 00458 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $1, $-1
		        0x01ca 00458 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $0, $-2
		        0x01ca 00458 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) CALL    runtime.morestack_noctxt(SB)
		        0x01cf 00463 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) PCDATA  $0, $-1
		        0x01cf 00463 (/Users/stone/Projects/go/stone/mastering-go/env/env.go:8) JMP     0
		        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 b7  eH..%....H;a....
		        0x0010 01 00 00 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24  ...H..`H.l$XH.l$
		        0x0020 58 0f 57 c0 0f 11 44 24 38 0f 11 44 24 48 48 8d  X.W...D$8..D$HH.
		        0x0030 05 00 00 00 00 48 89 44 24 38 48 8d 0d 00 00 00  .....H.D$8H.....
		        0x0040 00 48 89 4c 24 40 48 89 44 24 48 48 8d 0d 00 00  .H.L$@H.D$HH....
		        0x0050 00 00 48 89 4c 24 50 48 8d 0d 00 00 00 00 48 89  ..H.L$PH......H.
		        0x0060 0c 24 48 c7 44 24 08 21 00 00 00 48 8d 4c 24 38  .$H.D$.!...H.L$8
		        0x0070 48 89 4c 24 10 48 c7 44 24 18 02 00 00 00 48 c7  H.L$.H.D$.....H.
		        0x0080 44 24 20 02 00 00 00 e8 00 00 00 00 48 8d 05 00  D$ .........H...
		        0x0090 00 00 00 48 89 04 24 48 c7 44 24 08 09 00 00 00  ...H..$H.D$.....
		        0x00a0 e8 00 00 00 00 48 8b 44 24 10 0f 57 c0 0f 11 44  .....H.D$..W...D
		        0x00b0 24 28 48 8d 0d 00 00 00 00 48 89 4c 24 28 48 89  $(H......H.L$(H.
		        0x00c0 44 24 30 48 8d 05 00 00 00 00 48 89 04 24 48 c7  D$0H......H..$H.
		        0x00d0 44 24 08 14 00 00 00 48 8d 44 24 28 48 89 44 24  D$.....H.D$(H.D$
		        0x00e0 10 48 c7 44 24 18 01 00 00 00 48 c7 44 24 20 01  .H.D$.....H.D$ .
		        0x00f0 00 00 00 e8 00 00 00 00 48 63 05 00 00 00 00 48  ........Hc.....H
		        0x0100 89 04 24 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0  ..$.....H.D$..W.
		        0x0110 0f 11 44 24 28 48 8d 0d 00 00 00 00 48 89 4c 24  ..D$(H......H.L$
		        0x0120 28 48 89 44 24 30 48 8d 05 00 00 00 00 48 89 04  (H.D$0H......H..
		        0x0130 24 48 c7 44 24 08 13 00 00 00 48 8d 44 24 28 48  $H.D$.....H.D$(H
		        0x0140 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44  .D$.H.D$.....H.D
		        0x0150 24 20 01 00 00 00 e8 00 00 00 00 e8 00 00 00 00  $ ..............
		        0x0160 48 63 04 24 48 89 04 24 e8 00 00 00 00 48 8b 44  Hc.$H..$.....H.D
		        0x0170 24 08 0f 57 c0 0f 11 44 24 28 48 8d 0d 00 00 00  $..W...D$(H.....
		        0x0180 00 48 89 4c 24 28 48 89 44 24 30 48 8d 05 00 00  .H.L$(H.D$0H....
		        0x0190 00 00 48 89 04 24 48 c7 44 24 08 19 00 00 00 48  ..H..$H.D$.....H
		        0x01a0 8d 44 24 28 48 89 44 24 10 48 c7 44 24 18 01 00  .D$(H.D$.H.D$...
		        0x01b0 00 00 48 c7 44 24 20 01 00 00 00 e8 00 00 00 00  ..H.D$ .........
		        0x01c0 48 8b 6c 24 58 48 83 c4 60 c3 e8 00 00 00 00 e9  H.l$XH..`.......
		        0x01d0 2c fe ff ff                                      ,...
		        rel 5+4 t=17 TLS+0
		        rel 49+4 t=16 type.string+0
		        rel 61+4 t=16 ""..stmp_0+0
		        rel 78+4 t=16 ""..stmp_1+0
		        rel 90+4 t=16 go.string."You are using %v on a %v machine\n"+0
		        rel 136+4 t=8 log.Printf+0
		        rel 143+4 t=16 go.string."go1.14.13"+0
		        rel 161+4 t=8 runtime.convTstring+0
		        rel 181+4 t=16 type.string+0
		        rel 198+4 t=16 go.string."Using Go version %v\n"+0
		        rel 244+4 t=8 log.Printf+0
		        rel 251+4 t=16 runtime.ncpu+0
		        rel 260+4 t=8 runtime.convT64+0
		        rel 280+4 t=16 type.int+0
		        rel 297+4 t=16 go.string."Number of CPUs: %v\n"+0
		        rel 343+4 t=8 log.Printf+0
		        rel 348+4 t=8 runtime.gcount+0
		        rel 361+4 t=8 runtime.convT64+0
		        rel 381+4 t=16 type.int+0
		        rel 398+4 t=16 go.string."Number of Goroutines: %v\n"+0
		        rel 444+4 t=8 log.Printf+0
		        rel 459+4 t=8 runtime.morestack_noctxt+0
		type..eq.[2]interface {} STEXT dupok size=179 args=0x18 locals=0x30
		        0x0000 00000 (<autogenerated>:1)        TEXT    type..eq.[2]interface {}(SB), DUPOK|ABIInternal, $48-24
		        0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
		        0x0009 00009 (<autogenerated>:1)        CMPQ    SP, 16(CX)
		        0x000d 00013 (<autogenerated>:1)        PCDATA  $0, $-2
		        0x000d 00013 (<autogenerated>:1)        JLS     169
		        0x0013 00019 (<autogenerated>:1)        PCDATA  $0, $-1
		        0x0013 00019 (<autogenerated>:1)        SUBQ    $48, SP
		        0x0017 00023 (<autogenerated>:1)        MOVQ    BP, 40(SP)
		        0x001c 00028 (<autogenerated>:1)        LEAQ    40(SP), BP
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $0, $-2
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $1, $-2
		        0x0021 00033 (<autogenerated>:1)        FUNCDATA        $0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
		        0x0021 00033 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
		        0x0021 00033 (<autogenerated>:1)        FUNCDATA        $2, gclocals·313a5bdbfadc4f007c002a3a3588596d(SB)
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $0, $1
		        0x0021 00033 (<autogenerated>:1)        PCDATA  $1, $0
		        0x0021 00033 (<autogenerated>:1)        MOVQ    "".p+56(SP), AX
		        0x0026 00038 (<autogenerated>:1)        PCDATA  $0, $2
		        0x0026 00038 (<autogenerated>:1)        MOVQ    "".q+64(SP), CX
		        0x002b 00043 (<autogenerated>:1)        XORL    DX, DX
		        0x002d 00045 (<autogenerated>:1)        JMP     72
		        0x002f 00047 (<autogenerated>:1)        PCDATA  $0, $0
		        0x002f 00047 (<autogenerated>:1)        MOVQ    ""..autotmp_8+32(SP), BX
		        0x0034 00052 (<autogenerated>:1)        LEAQ    1(BX), DX
		        0x0038 00056 (<autogenerated>:1)        PCDATA  $0, $3
		        0x0038 00056 (<autogenerated>:1)        MOVQ    "".p+56(SP), BX
		        0x003d 00061 (<autogenerated>:1)        PCDATA  $0, $4
		        0x003d 00061 (<autogenerated>:1)        MOVQ    "".q+64(SP), SI
		        0x0042 00066 (<autogenerated>:1)        PCDATA  $0, $5
		        0x0042 00066 (<autogenerated>:1)        MOVQ    BX, AX
		        0x0045 00069 (<autogenerated>:1)        PCDATA  $0, $2
		        0x0045 00069 (<autogenerated>:1)        MOVQ    SI, CX
		        0x0048 00072 (<autogenerated>:1)        CMPQ    DX, $2
		        0x004c 00076 (<autogenerated>:1)        JGE     154
		        0x004e 00078 (<autogenerated>:1)        MOVQ    DX, BX
		        0x0051 00081 (<autogenerated>:1)        SHLQ    $4, DX
		        0x0055 00085 (<autogenerated>:1)        PCDATA  $0, $6
		        0x0055 00085 (<autogenerated>:1)        MOVQ    8(DX)(AX*1), SI
		        0x005a 00090 (<autogenerated>:1)        PCDATA  $0, $7
		        0x005a 00090 (<autogenerated>:1)        MOVQ    (DX)(AX*1), DI
		        0x005e 00094 (<autogenerated>:1)        MOVQ    (DX)(CX*1), R8
		        0x0062 00098 (<autogenerated>:1)        PCDATA  $0, $8
		        0x0062 00098 (<autogenerated>:1)        MOVQ    8(DX)(CX*1), DX
		        0x0067 00103 (<autogenerated>:1)        CMPQ    DI, R8
		        0x006a 00106 (<autogenerated>:1)        JNE     139
		        0x006c 00108 (<autogenerated>:1)        MOVQ    BX, ""..autotmp_8+32(SP)
		        0x0071 00113 (<autogenerated>:1)        MOVQ    DI, (SP)
		        0x0075 00117 (<autogenerated>:1)        PCDATA  $0, $9
		        0x0075 00117 (<autogenerated>:1)        MOVQ    SI, 8(SP)
		        0x007a 00122 (<autogenerated>:1)        PCDATA  $0, $0
		        0x007a 00122 (<autogenerated>:1)        MOVQ    DX, 16(SP)
		        0x007f 00127 (<autogenerated>:1)        CALL    runtime.efaceeq(SB)
		        0x0084 00132 (<autogenerated>:1)        CMPB    24(SP), $0
		        0x0089 00137 (<autogenerated>:1)        JNE     47
		        0x008b 00139 (<autogenerated>:1)        PCDATA  $1, $1
		        0x008b 00139 (<autogenerated>:1)        MOVB    $0, "".~r2+72(SP)
		        0x0090 00144 (<autogenerated>:1)        MOVQ    40(SP), BP
		        0x0095 00149 (<autogenerated>:1)        ADDQ    $48, SP
		        0x0099 00153 (<autogenerated>:1)        RET
		        0x009a 00154 (<autogenerated>:1)        MOVB    $1, "".~r2+72(SP)
		        0x009f 00159 (<autogenerated>:1)        MOVQ    40(SP), BP
		        0x00a4 00164 (<autogenerated>:1)        ADDQ    $48, SP
		        0x00a8 00168 (<autogenerated>:1)        RET
		        0x00a9 00169 (<autogenerated>:1)        NOP
		        0x00a9 00169 (<autogenerated>:1)        PCDATA  $1, $-1
		        0x00a9 00169 (<autogenerated>:1)        PCDATA  $0, $-2
		        0x00a9 00169 (<autogenerated>:1)        CALL    runtime.morestack_noctxt(SB)
		        0x00ae 00174 (<autogenerated>:1)        PCDATA  $0, $-1
		        0x00ae 00174 (<autogenerated>:1)        JMP     0
		        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 96  eH..%....H;a....
		        0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
		        0x0020 28 48 8b 44 24 38 48 8b 4c 24 40 31 d2 eb 19 48  (H.D$8H.L$@1...H
		        0x0030 8b 5c 24 20 48 8d 53 01 48 8b 5c 24 38 48 8b 74  .\$ H.S.H.\$8H.t
		        0x0040 24 40 48 89 d8 48 89 f1 48 83 fa 02 7d 4c 48 89  $@H..H..H...}LH.
		        0x0050 d3 48 c1 e2 04 48 8b 74 02 08 48 8b 3c 02 4c 8b  .H...H.t..H.<.L.
		        0x0060 04 0a 48 8b 54 0a 08 4c 39 c7 75 1f 48 89 5c 24  ..H.T..L9.u.H.\$
		        0x0070 20 48 89 3c 24 48 89 74 24 08 48 89 54 24 10 e8   H.<$H.t$.H.T$..
		        0x0080 00 00 00 00 80 7c 24 18 00 75 a4 c6 44 24 48 00  .....|$..u..D$H.
		        0x0090 48 8b 6c 24 28 48 83 c4 30 c3 c6 44 24 48 01 48  H.l$(H..0..D$H.H
		        0x00a0 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00 00 e9 4d  .l$(H..0.......M
		        0x00b0 ff ff ff                                         ...
		        rel 5+4 t=17 TLS+0
		        rel 128+4 t=8 runtime.efaceeq+0
		        rel 170+4 t=8 runtime.morestack_noctxt+0
		go.cuinfo.packagename.main SDWARFINFO dupok size=0
		        0x0000 6d 61 69 6e                                      main
		go.info.runtime.Version$abstract SDWARFINFO dupok size=20
		        0x0000 04 72 75 6e 74 69 6d 65 2e 56 65 72 73 69 6f 6e  .runtime.Version
		        0x0010 00 01 01 00                                      ....
		go.info.runtime.NumCPU$abstract SDWARFINFO dupok size=19
		        0x0000 04 72 75 6e 74 69 6d 65 2e 4e 75 6d 43 50 55 00  .runtime.NumCPU.
		        0x0010 01 01 00                                         ...
		go.info.runtime.NumGoroutine$abstract SDWARFINFO dupok size=25
		        0x0000 04 72 75 6e 74 69 6d 65 2e 4e 75 6d 47 6f 72 6f  .runtime.NumGoro
		        0x0010 75 74 69 6e 65 00 01 01 00                       utine....
		go.string."gc" SRODATA dupok size=2
		        0x0000 67 63                                            gc
		go.string."amd64" SRODATA dupok size=5
		        0x0000 61 6d 64 36 34                                   amd64
		go.string."You are using %v on a %v machine\n" SRODATA dupok size=33
		        0x0000 59 6f 75 20 61 72 65 20 75 73 69 6e 67 20 25 76  You are using %v
		        0x0010 20 6f 6e 20 61 20 25 76 20 6d 61 63 68 69 6e 65   on a %v machine
		        0x0020 0a                                               .
		go.string."go1.14.13" SRODATA dupok size=9
		        0x0000 67 6f 31 2e 31 34 2e 31 33                       go1.14.13
		go.string."Using Go version %v\n" SRODATA dupok size=20
		        0x0000 55 73 69 6e 67 20 47 6f 20 76 65 72 73 69 6f 6e  Using Go version
		        0x0010 20 25 76 0a                                       %v.
		go.string."Number of CPUs: %v\n" SRODATA dupok size=19
		        0x0000 4e 75 6d 62 65 72 20 6f 66 20 43 50 55 73 3a 20  Number of CPUs:
		        0x0010 25 76 0a                                         %v.
		go.string."Number of Goroutines: %v\n" SRODATA dupok size=25
		        0x0000 4e 75 6d 62 65 72 20 6f 66 20 47 6f 72 6f 75 74  Number of Gorout
		        0x0010 69 6e 65 73 3a 20 25 76 0a                       ines: %v.
		go.loc."".main SDWARFLOC size=0
		go.info."".main SDWARFINFO size=95
		        0x0000 03 6d 61 69 6e 2e 6d 61 69 6e 00 00 00 00 00 00  .main.main......
		        0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
		        0x0020 00 01 06 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 0b 00 00 00 00  ................
		        0x0040 06 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0050 00 00 00 00 00 00 00 00 00 0c 00 00 00 00 00     ...............
		        rel 0+0 t=24 type.[1]interface {}+0
		        rel 0+0 t=24 type.[2]interface {}+0
		        rel 11+8 t=1 "".main+0
		        rel 19+8 t=1 "".main+468
		        rel 29+4 t=30 gofile../Users/stone/Projects/go/stone/mastering-go/env/env.go+0
		        rel 35+4 t=29 go.info.runtime.NumCPU$abstract+0
		        rel 39+8 t=1 "".main+248
		        rel 47+8 t=1 "".main+255
		        rel 55+4 t=30 gofile../Users/stone/Projects/go/stone/mastering-go/env/env.go+0
		        rel 65+4 t=29 go.info.runtime.NumGoroutine$abstract+0
		        rel 69+8 t=1 "".main+347
		        rel 77+8 t=1 "".main+356
		        rel 85+4 t=30 gofile../Users/stone/Projects/go/stone/mastering-go/env/env.go+0
		go.range."".main SDWARFRANGE size=0
		go.debuglines."".main SDWARFMISC size=69
		        0x0000 04 02 03 02 14 0a cd 9c 06 2d 06 02 51 f6 06 55  .........-..Q..U
		        0x0010 06 91 06 41 04 07 06 02 3c 03 1d fa 04 02 03 63  ...A....<......c
		        0x0020 51 06 69 04 07 06 02 3c 03 2a fa 06 41 04 02 06  Q.i....<.*..A...
		        0x0030 03 56 33 06 69 06 02 3c 03 8f 04 fa 03 ec 7b 6f  .V3.i..<......{o
		        0x0040 04 01 03 79 01                                   ...y.
		runtime.nilinterequal·f SRODATA dupok size=8
		        0x0000 00 00 00 00 00 00 00 00                          ........
		        rel 0+8 t=1 runtime.nilinterequal+0
		runtime.memequal64·f SRODATA dupok size=8
		        0x0000 00 00 00 00 00 00 00 00                          ........
		        rel 0+8 t=1 runtime.memequal64+0
		runtime.gcbits.01 SRODATA dupok size=1
		        0x0000 01                                               .
		type..namedata.*interface {}- SRODATA dupok size=16
		        0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
		type.*interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*interface {}-+0
		        rel 48+8 t=1 type.interface {}+0
		runtime.gcbits.02 SRODATA dupok size=1
		        0x0000 02                                               .
		type.interface {} SRODATA dupok size=80
		        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
		        0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        rel 24+8 t=1 runtime.nilinterequal·f+0
		        rel 32+8 t=1 runtime.gcbits.02+0
		        rel 40+4 t=5 type..namedata.*interface {}-+0
		        rel 44+4 t=6 type.*interface {}+0
		        rel 56+8 t=1 type.interface {}+80
		type..namedata.*[]interface {}- SRODATA dupok size=18
		        0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface
		        0x0010 7b 7d                                            {}
		type.*[]interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[]interface {}-+0
		        rel 48+8 t=1 type.[]interface {}+0
		type.[]interface {} SRODATA dupok size=56
		        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[]interface {}-+0
		        rel 44+4 t=6 type.*[]interface {}+0
		        rel 48+8 t=1 type.interface {}+0
		go.loc.type..eq.[2]interface {} SDWARFLOC dupok size=106
		        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0010 01 00 51 00 00 00 00 00 00 00 00 00 00 00 00 00  ..Q.............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 01 00 9c 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0050 00 00 00 00 00 00 02 00 91 08 00 00 00 00 00 00  ................
		        0x0060 00 00 00 00 00 00 00 00 00 00                    ..........
		        rel 0+8 t=53 type..eq.[2]interface {}+72
		        rel 8+8 t=53 type..eq.[2]interface {}+85
		        rel 35+8 t=53 type..eq.[2]interface {}+0
		        rel 43+8 t=53 type..eq.[2]interface {}+179
		        rel 70+8 t=53 type..eq.[2]interface {}+0
		        rel 78+8 t=53 type..eq.[2]interface {}+179
		go.info.type..eq.[2]interface {} SDWARFINFO dupok size=100
		        0x0000 03 74 79 70 65 2e 2e 65 71 2e 5b 32 5d 69 6e 74  .type..eq.[2]int
		        0x0010 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00 00 00  erface {}.......
		        0x0020 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
		        0x0030 01 0b 69 00 01 00 00 00 00 00 00 00 00 10 70 00  ..i...........p.
		        0x0040 00 01 00 00 00 00 00 00 00 00 10 71 00 00 01 00  ...........q....
		        0x0050 00 00 00 00 00 00 00 0f 7e 72 32 00 01 01 00 00  ........~r2.....
		        0x0060 00 00 00 00                                      ....
		        rel 0+0 t=24 type.*[2]interface {}+0
		        rel 0+0 t=24 type.bool+0
		        rel 0+0 t=24 type.int+0
		        rel 26+8 t=1 type..eq.[2]interface {}+0
		        rel 34+8 t=1 type..eq.[2]interface {}+179
		        rel 44+4 t=30 gofile..<autogenerated>+0
		        rel 53+4 t=29 go.info.int+0
		        rel 57+4 t=29 go.loc.type..eq.[2]interface {}+0
		        rel 66+4 t=29 go.info.*[2]interface {}+0
		        rel 70+4 t=29 go.loc.type..eq.[2]interface {}+35
		        rel 79+4 t=29 go.info.*[2]interface {}+0
		        rel 83+4 t=29 go.loc.type..eq.[2]interface {}+70
		        rel 94+4 t=29 go.info.bool+0
		go.range.type..eq.[2]interface {} SDWARFRANGE dupok size=0
		go.debuglines.type..eq.[2]interface {} SDWARFMISC dupok size=29
		        0x0000 04 01 0f 0a cd 06 cd 06 08 73 06 37 06 02 27 ff  .........s.7..'.
		        0x0010 06 41 06 73 06 41 06 73 04 01 03 00 01           .A.s.A.s.....
		type..eqfunc.[2]interface {} SRODATA dupok size=8
		        0x0000 00 00 00 00 00 00 00 00                          ........
		        rel 0+8 t=1 type..eq.[2]interface {}+0
		type..namedata.*[2]interface {}- SRODATA dupok size=19
		        0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
		        0x0010 20 7b 7d                                          {}
		type.*[2]interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 be 73 2d 71 08 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[2]interface {}-+0
		        rel 48+8 t=1 type.[2]interface {}+0
		runtime.gcbits.0a SRODATA dupok size=1
		        0x0000 0a                                               .
		type.[2]interface {} SRODATA dupok size=72
		        0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
		        0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 02 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 type..eqfunc.[2]interface {}+0
		        rel 32+8 t=1 runtime.gcbits.0a+0
		        rel 40+4 t=5 type..namedata.*[2]interface {}-+0
		        rel 44+4 t=6 type.*[2]interface {}+0
		        rel 48+8 t=1 type.interface {}+0
		        rel 56+8 t=1 type.[]interface {}+0
		type..namedata.*[1]interface {}- SRODATA dupok size=19
		        0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
		        0x0010 20 7b 7d                                          {}
		type.*[1]interface {} SRODATA dupok size=56
		        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
		        0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.memequal64·f+0
		        rel 32+8 t=1 runtime.gcbits.01+0
		        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
		        rel 48+8 t=1 type.[1]interface {}+0
		type.[1]interface {} SRODATA dupok size=72
		        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
		        0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
		        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0040 01 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 runtime.nilinterequal·f+0
		        rel 32+8 t=1 runtime.gcbits.02+0
		        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
		        rel 44+4 t=6 type.*[1]interface {}+0
		        rel 48+8 t=1 type.interface {}+0
		        rel 56+8 t=1 type.[]interface {}+0
		go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tmastering-go\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2" SRODATA dupok size=86
		        0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0w...t..A.......
		        0x0010 70 61 74 68 09 63 6f 6d 6d 61 6e 64 2d 6c 69 6e  path.command-lin
		        0x0020 65 2d 61 72 67 75 6d 65 6e 74 73 0a 6d 6f 64 09  e-arguments.mod.
		        0x0030 6d 61 73 74 65 72 69 6e 67 2d 67 6f 09 28 64 65  mastering-go.(de
		        0x0040 76 65 6c 29 09 0a f9 32 43 31 86 18 20 72 00 82  vel)...2C1.. r..
		        0x0050 42 10 41 16 d8 f2                                B.A...
		""..inittask SNOPTRDATA size=40
		        0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
		        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
		        0x0020 00 00 00 00 00 00 00 00                          ........
		        rel 24+8 t=1 log..inittask+0
		        rel 32+8 t=1 runtime..inittask+0
		runtime.modinfo SDATA size=16
		        0x0000 00 00 00 00 00 00 00 00 56 00 00 00 00 00 00 00  ........V.......
		        rel 0+8 t=1 go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tmastering-go\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"+0
		""..stmp_0 SRODATA size=16
		        0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
		        rel 0+8 t=1 go.string."gc"+0
		""..stmp_1 SRODATA size=16
		        0x0000 00 00 00 00 00 00 00 00 05 00 00 00 00 00 00 00  ................
		        rel 0+8 t=1 go.string."amd64"+0
		type..importpath.log. SRODATA dupok size=6
		        0x0000 00 00 03 6c 6f 67                                ...log
		type..importpath.runtime. SRODATA dupok size=10
		        0x0000 00 00 07 72 75 6e 74 69 6d 65                    ...runtime
		type..importpath.unsafe. SRODATA dupok size=9
		        0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
		gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
		        0x0000 03 00 00 00 00 00 00 00                          ........
		gclocals·669b766e4959babb5d70e9d2b4843b8e SRODATA dupok size=11
		        0x0000 03 00 00 00 06 00 00 00 00 28 02                 .........(.
		gclocals·6e8d7ea4abad763909b26991048ee1fe SRODATA dupok size=12
		        0x0000 04 00 00 00 02 00 00 00 00 01 03 02              ............
		"".main.stkobj SRODATA size=40
		        0x0000 02 00 00 00 00 00 00 00 d0 ff ff ff ff ff ff ff  ................
		        0x0010 00 00 00 00 00 00 00 00 e0 ff ff ff ff ff ff ff  ................
		        0x0020 00 00 00 00 00 00 00 00                          ........
		        rel 16+8 t=1 type.[1]interface {}+0
		        rel 32+8 t=1 type.[2]interface {}+0
		gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
		        0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
		gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
		        0x0000 02 00 00 00 00 00 00 00                          ........
		gclocals·313a5bdbfadc4f007c002a3a3588596d SRODATA dupok size=18
		        0x0000 0a 00 00 00 06 00 00 00 00 01 03 08 28 21 23 22  ............(!#"
		        0x0010 24 04                                            $.

	*/

	/*
		stone@stonedeMacBook-Pro env % go tool compile -W env.go

		before walk main
		.   CALLFUNC l(9) tc(1)
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(9) esc(no) PTR-@0
		.   CALLFUNC-list
		.   .   LITERAL-"You are using %v on a %v machine\n" l(9) tc(1) string

		.   .   CONVIFACE l(9) esc(h) tc(1) implicit(true) INTER-@0
		.   .   .   NAME-main..stmp_0 l(9) x(0) class(PEXTERN) tc(1) used string

		.   .   CONVIFACE l(9) esc(h) tc(1) implicit(true) INTER-@0
		.   .   .   NAME-main..stmp_1 l(9) x(0) class(PEXTERN) tc(1) used string

		.   VARKILL l(9) tc(1)
		.   .   NAME-main..autotmp_3 l(9) x(0) class(PAUTO) esc(N) used ARRAY-@0

		.   DCL l(10)
		.   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string

		.   AS l(10) tc(1)
		.   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string

		.   INLMARK l(+10) x(0)

		.   AS2 l(10) tc(1)
		.   AS2-list
		.   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string
		.   AS2-rlist
		.   .   LITERAL-"go1.14.13" l(5) tc(1) string

		.   GOTO l(10) tc(1) main..i0

		.   LABEL l(10) tc(1) main..i0

		.   CALLFUNC l(10) tc(1)
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(10) esc(no) PTR-@0
		.   CALLFUNC-list
		.   .   LITERAL-"Using Go version %v\n" l(10) tc(1) string

		.   .   CONVIFACE l(10) esc(h) tc(1) implicit(true) INTER-@0
		.   .   .   CONVNOP l(10) tc(1) hascall string
		.   .   .   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string

		.   VARKILL l(10) tc(1)
		.   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) used ARRAY-@0

		.   DCL l(11)
		.   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   AS l(11) tc(1)
		.   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   INLMARK l(+11) x(1)

		.   AS2 l(11) tc(1)
		.   AS2-list
		.   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int
		.   AS2-rlist
		.   .   CONV l(44) tc(1) int
		.   .   .   NAME-runtime.ncpu l(1019) x(0) class(PEXTERN) tc(1) used int32

		.   GOTO l(11) tc(1) main..i1

		.   LABEL l(11) tc(1) main..i1

		.   CALLFUNC l(11) tc(1)
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(11) esc(no) PTR-@0
		.   CALLFUNC-list
		.   .   LITERAL-"Number of CPUs: %v\n" l(11) tc(1) string

		.   .   CONVIFACE l(11) esc(h) tc(1) implicit(true) INTER-@0
		.   .   .   CONVNOP l(11) tc(1) hascall int
		.   .   .   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   VARKILL l(11) tc(1)
		.   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) used ARRAY-@0

		.   DCL l(12)
		.   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   AS l(12) tc(1)
		.   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   INLMARK l(+12) x(2)

		.   AS l(58) tc(1)
		.   .   NAME-main..autotmp_5 l(58) x(0) class(PAUTO) esc(N) tc(1) assigned used int32
		.   .   CALLFUNC l(58) tc(1) int32
		.   .   .   NAME-runtime.gcount l(3797) x(0) class(PFUNC) tc(1) used FUNC-@0

		.   AS2 l(12) tc(1)
		.   AS2-list
		.   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int
		.   AS2-rlist
		.   .   CONV l(58) tc(1) int
		.   .   .   NAME-main..autotmp_5 l(58) x(0) class(PAUTO) esc(N) tc(1) assigned used int32

		.   VARKILL l(12) tc(1)
		.   .   NAME-main..autotmp_5 l(58) x(0) class(PAUTO) esc(N) tc(1) assigned used int32

		.   GOTO l(12) tc(1) main..i2

		.   LABEL l(12) tc(1) main..i2

		.   CALLFUNC l(12) tc(1)
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(12) esc(no) PTR-@0
		.   CALLFUNC-list
		.   .   LITERAL-"Number of Goroutines: %v\n" l(12) tc(1) string

		.   .   CONVIFACE l(12) esc(h) tc(1) implicit(true) INTER-@0
		.   .   .   CONVNOP l(12) tc(1) hascall int
		.   .   .   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   VARKILL l(12) tc(1)
		.   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) used ARRAY-@0
		after walk main
		.   CALLFUNC-init
		.   .   AS l(9) tc(1)
		.   .   .   NAME-main..autotmp_3 l(9) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   AS l(9) tc(1)
		.   .   .   NAME-main..autotmp_7 l(9) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   ADDR l(9) tc(1) PTR-@0
		.   .   .   .   NAME-main..autotmp_3 l(9) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   BLOCK l(9)
		.   .   BLOCK-list
		.   .   .   AS l(9) tc(1) hascall
		.   .   .   .   INDEX l(9) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(9) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main..autotmp_7 l(9) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   .   .   LITERAL-0 l(9) tc(1) int
		.   .   .   .   EFACE l(9) tc(1) INTER-@0
		.   .   .   .   .   ADDR l(9) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-type.string x(0) class(PEXTERN) tc(1) uint8
		.   .   .   .   .   ADDR l(9) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-main..stmp_0 l(9) x(0) class(PEXTERN) tc(1) addrtaken used string

		.   .   BLOCK l(9)
		.   .   BLOCK-list
		.   .   .   AS l(9) tc(1) hascall
		.   .   .   .   INDEX l(9) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(9) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main..autotmp_7 l(9) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   .   .   LITERAL-1 l(9) tc(1) int
		.   .   .   .   EFACE l(9) tc(1) INTER-@0
		.   .   .   .   .   ADDR l(9) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-type.string x(0) class(PEXTERN) tc(1) uint8
		.   .   .   .   .   ADDR l(9) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-main..stmp_1 l(9) x(0) class(PEXTERN) tc(1) addrtaken used string

		.   .   BLOCK l(9)
		.   .   BLOCK-list
		.   .   .   AS l(9) tc(1) hascall
		.   .   .   .   NAME-main..autotmp_6 l(9) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0
		.   .   .   .   SLICEARR l(9) tc(1) hascall SLICE-@0
		.   .   .   .   .   NAME-main..autotmp_7 l(9) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   CALLFUNC l(9) tc(1) hascall
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(9) esc(no) PTR-@0
		.   CALLFUNC-rlist
		.   .   LITERAL-"You are using %v on a %v machine\n" l(9) tc(1) string

		.   .   NAME-main..autotmp_6 l(9) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0

		.   VARKILL l(9) tc(1)
		.   .   NAME-main..autotmp_3 l(9) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   DCL l(10)
		.   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string

		.   AS l(10) tc(1)
		.   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string

		.   INLMARK l(+10) x(0)

		.   BLOCK l(10)
		.   BLOCK-list
		.   .   AS l(10) tc(1)
		.   .   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string
		.   .   .   LITERAL-"go1.14.13" l(5) tc(1) string

		.   GOTO l(10) tc(1) main..i0

		.   LABEL l(10) tc(1) main..i0

		.   CALLFUNC-init
		.   .   AS l(10) tc(1) hascall
		.   .   .   NAME-main..autotmp_8 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used UNSAFEPTR-@0
		.   .   .   CALLFUNC l(10) tc(1) hascall UNSAFEPTR-@0
		.   .   .   .   NAME-runtime.convTstring x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   .   CALLFUNC-rlist
		.   .   .   .   CONVNOP l(10) tc(1) string
		.   .   .   .   .   NAME-main.~R0 l(10) x(0) class(PAUTO) esc(no) tc(1) assigned used string

		.   .   AS l(10) tc(1)
		.   .   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   AS l(10) tc(1)
		.   .   .   NAME-main..autotmp_10 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   ADDR l(10) tc(1) PTR-@0
		.   .   .   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   BLOCK l(10)
		.   .   BLOCK-list
		.   .   .   AS l(10) tc(1) hascall
		.   .   .   .   INDEX l(10) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(10) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main..autotmp_10 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   .   .   LITERAL-0 l(10) tc(1) int
		.   .   .   .   EFACE l(10) tc(1) INTER-@0
		.   .   .   .   .   ADDR l(10) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-type.string x(0) class(PEXTERN) tc(1) uint8
		.   .   .   .   .   NAME-main..autotmp_8 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used UNSAFEPTR-@0

		.   .   BLOCK l(10)
		.   .   BLOCK-list
		.   .   .   AS l(10) tc(1) hascall
		.   .   .   .   NAME-main..autotmp_9 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0
		.   .   .   .   SLICEARR l(10) tc(1) hascall SLICE-@0
		.   .   .   .   .   NAME-main..autotmp_10 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   CALLFUNC l(10) tc(1) hascall
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(10) esc(no) PTR-@0
		.   CALLFUNC-rlist
		.   .   LITERAL-"Using Go version %v\n" l(10) tc(1) string

		.   .   NAME-main..autotmp_9 l(10) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0

		.   VARKILL l(10) tc(1)
		.   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   DCL l(11)
		.   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   AS l(11) tc(1)
		.   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   INLMARK l(+11) x(1)

		.   BLOCK-init
		.   .   AS l(11) tc(1)
		.   .   .   NAME-main..autotmp_11 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   .   CONV l(44) tc(1) int
		.   .   .   .   NAME-runtime.ncpu l(1019) x(0) class(PEXTERN) tc(1) used int32
		.   BLOCK l(11) hascall
		.   BLOCK-list
		.   .   AS l(11) tc(1)
		.   .   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int
		.   .   .   NAME-main..autotmp_11 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used int

		.   GOTO l(11) tc(1) main..i1

		.   LABEL l(11) tc(1) main..i1

		.   CALLFUNC-init
		.   .   AS l(11) tc(1) hascall
		.   .   .   NAME-main..autotmp_12 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used UNSAFEPTR-@0
		.   .   .   CALLFUNC l(11) tc(1) hascall UNSAFEPTR-@0
		.   .   .   .   NAME-runtime.convT64 x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   .   CALLFUNC-rlist
		.   .   .   .   CONVNOP l(11) tc(1) int
		.   .   .   .   .   NAME-main.~R0 l(11) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   .   AS l(11) tc(1)
		.   .   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   AS l(11) tc(1)
		.   .   .   NAME-main..autotmp_14 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   ADDR l(11) tc(1) PTR-@0
		.   .   .   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   BLOCK l(11)
		.   .   BLOCK-list
		.   .   .   AS l(11) tc(1) hascall
		.   .   .   .   INDEX l(11) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(11) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main..autotmp_14 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   .   .   LITERAL-0 l(11) tc(1) int
		.   .   .   .   EFACE l(11) tc(1) INTER-@0
		.   .   .   .   .   ADDR l(11) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-type.int x(0) class(PEXTERN) tc(1) uint8
		.   .   .   .   .   NAME-main..autotmp_12 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used UNSAFEPTR-@0

		.   .   BLOCK l(11)
		.   .   BLOCK-list
		.   .   .   AS l(11) tc(1) hascall
		.   .   .   .   NAME-main..autotmp_13 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0
		.   .   .   .   SLICEARR l(11) tc(1) hascall SLICE-@0
		.   .   .   .   .   NAME-main..autotmp_14 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   CALLFUNC l(11) tc(1) hascall
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(11) esc(no) PTR-@0
		.   CALLFUNC-rlist
		.   .   LITERAL-"Number of CPUs: %v\n" l(11) tc(1) string

		.   .   NAME-main..autotmp_13 l(11) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0

		.   VARKILL l(11) tc(1)
		.   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   DCL l(12)
		.   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   AS l(12) tc(1)
		.   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   INLMARK l(+12) x(2)

		.   AS l(58) tc(1) hascall
		.   .   NAME-main..autotmp_5 l(58) x(0) class(PAUTO) esc(N) tc(1) assigned used int32
		.   .   CALLFUNC l(58) tc(1) hascall int32
		.   .   .   NAME-runtime.gcount l(3797) x(0) class(PFUNC) tc(1) used FUNC-@0

		.   BLOCK-init
		.   .   AS l(12) tc(1)
		.   .   .   NAME-main..autotmp_15 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   .   CONV l(58) tc(1) int
		.   .   .   .   NAME-main..autotmp_5 l(58) x(0) class(PAUTO) esc(N) tc(1) assigned used int32
		.   BLOCK l(12) hascall
		.   BLOCK-list
		.   .   AS l(12) tc(1)
		.   .   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int
		.   .   .   NAME-main..autotmp_15 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used int

		.   VARKILL l(12) tc(1)
		.   .   NAME-main..autotmp_5 l(58) x(0) class(PAUTO) esc(N) tc(1) assigned used int32

		.   GOTO l(12) tc(1) main..i2

		.   LABEL l(12) tc(1) main..i2

		.   CALLFUNC-init
		.   .   AS l(12) tc(1) hascall
		.   .   .   NAME-main..autotmp_16 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used UNSAFEPTR-@0
		.   .   .   CALLFUNC l(12) tc(1) hascall UNSAFEPTR-@0
		.   .   .   .   NAME-runtime.convT64 x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   .   CALLFUNC-rlist
		.   .   .   .   CONVNOP l(12) tc(1) int
		.   .   .   .   .   NAME-main.~R0 l(12) x(0) class(PAUTO) esc(no) tc(1) assigned used int

		.   .   AS l(12) tc(1)
		.   .   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   AS l(12) tc(1)
		.   .   .   NAME-main..autotmp_18 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   ADDR l(12) tc(1) PTR-@0
		.   .   .   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		.   .   BLOCK l(12)
		.   .   BLOCK-list
		.   .   .   AS l(12) tc(1) hascall
		.   .   .   .   INDEX l(12) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(12) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main..autotmp_18 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   .   .   .   .   LITERAL-0 l(12) tc(1) int
		.   .   .   .   EFACE l(12) tc(1) INTER-@0
		.   .   .   .   .   ADDR l(12) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-type.int x(0) class(PEXTERN) tc(1) uint8
		.   .   .   .   .   NAME-main..autotmp_16 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used UNSAFEPTR-@0

		.   .   BLOCK l(12)
		.   .   BLOCK-list
		.   .   .   AS l(12) tc(1) hascall
		.   .   .   .   NAME-main..autotmp_17 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0
		.   .   .   .   SLICEARR l(12) tc(1) hascall SLICE-@0
		.   .   .   .   .   NAME-main..autotmp_18 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-@0
		.   CALLFUNC l(12) tc(1) hascall
		.   .   NAME-log.Printf l(319) x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   DDDARG l(12) esc(no) PTR-@0
		.   CALLFUNC-rlist
		.   .   LITERAL-"Number of Goroutines: %v\n" l(12) tc(1) string

		.   .   NAME-main..autotmp_17 l(12) x(0) class(PAUTO) esc(N) tc(1) assigned used SLICE-@0

		.   VARKILL l(12) tc(1)
		.   .   NAME-main..autotmp_4 l(10) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-@0

		before walk type..eq.[2]interface {}
		.   DCL l(1)
		.   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int

		.   RANGE l(1) colas(true) tc(1) ARRAY-@0
		.   .   DEREF l(1) tc(1) ARRAY-@0
		.   .   .   NAME-main.p g(2) l(1) x(0) class(PPARAM) tc(1) used PTR-@0
		.   RANGE-list
		.   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int
		.   RANGE-body
		.   .   IF l(1) tc(1)
		.   .   .   NE l(1) tc(1) bool
		.   .   .   .   INDEX l(1) tc(1) bounded INTER-@0
		.   .   .   .   .   DEREF l(1) tc(1) implicit(true) ARRAY-@0
		.   .   .   .   .   .   NAME-main.p g(2) l(1) x(0) class(PPARAM) tc(1) used PTR-@0
		.   .   .   .   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int
		.   .   .   .   INDEX l(1) tc(1) bounded INTER-@0
		.   .   .   .   .   DEREF l(1) tc(1) implicit(true) ARRAY-@0
		.   .   .   .   .   .   NAME-main.q g(3) l(1) x(8) class(PPARAM) tc(1) used PTR-@0
		.   .   .   .   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int
		.   .   IF-body
		.   .   .   RETURN l(1) tc(1)
		.   .   .   RETURN-list
		.   .   .   .   LITERAL-false l(1) tc(1) bool

		.   RETURN l(1) tc(1)
		.   RETURN-list
		.   .   LITERAL-true l(1) tc(1) bool
		after walk type..eq.[2]interface {}
		.   DCL l(1)
		.   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int

		.   FOR-init
		.   .   AS l(1) tc(1)
		.   .   .   NAME-main..autotmp_4 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int

		.   .   AS l(1) tc(1)
		.   .   .   NAME-main..autotmp_5 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   .   LITERAL-2 l(1) tc(1) int
		.   FOR l(1) colas(true) tc(1) ARRAY-@0
		.   .   LT l(1) tc(1) bool
		.   .   .   NAME-main..autotmp_4 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   .   NAME-main..autotmp_5 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   AS l(1) tc(1)
		.   .   .   NAME-main..autotmp_4 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   .   ADD l(1) tc(1) int
		.   .   .   .   NAME-main..autotmp_4 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int
		.   .   .   .   LITERAL-1 l(1) tc(1) int
		.   FOR-body
		.   .   AS l(1) tc(1)
		.   .   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int
		.   .   .   NAME-main..autotmp_4 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used int

		.   .   IF-init
		.   .   .   AS l(1) tc(1) hascall
		.   .   .   .   NAME-main..autotmp_6 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0
		.   .   .   .   INDEX l(1) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(1) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main.q g(3) l(1) x(8) class(PPARAM) tc(1) used PTR-@0
		.   .   .   .   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int

		.   .   .   AS l(1) tc(1) hascall
		.   .   .   .   NAME-main..autotmp_7 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0
		.   .   .   .   INDEX l(1) tc(1) bounded hascall INTER-@0
		.   .   .   .   .   DEREF l(1) tc(1) implicit(true) hascall ARRAY-@0
		.   .   .   .   .   .   NAME-main.p g(2) l(1) x(0) class(PPARAM) tc(1) used PTR-@0
		.   .   .   .   .   NAME-main.i g(4) l(1) x(0) class(PAUTO) tc(1) assigned used int
		.   .   IF l(1) tc(1)
		.   .   .   OROR l(1) tc(1) hascall bool
		.   .   .   .   NE l(1) tc(1) bool
		.   .   .   .   .   ITAB l(1) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-main..autotmp_7 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0
		.   .   .   .   .   ITAB l(1) tc(1) PTR-@0
		.   .   .   .   .   .   NAME-main..autotmp_6 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0
		.   .   .   .   NOT l(1) tc(1) hascall bool
		.   .   .   .   .   CALLFUNC l(1) tc(1) hascall bool
		.   .   .   .   .   .   NAME-runtime.efaceeq x(0) class(PFUNC) tc(1) used FUNC-@0
		.   .   .   .   .   CALLFUNC-rlist
		.   .   .   .   .   .   ITAB l(1) tc(1) PTR-@0
		.   .   .   .   .   .   .   NAME-main..autotmp_7 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0

		.   .   .   .   .   .   IDATA l(1) tc(1) UNSAFEPTR-@0
		.   .   .   .   .   .   .   NAME-main..autotmp_7 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0

		.   .   .   .   .   .   IDATA l(1) tc(1) UNSAFEPTR-@0
		.   .   .   .   .   .   .   NAME-main..autotmp_6 l(1) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-@0
		.   .   IF-body
		.   .   .   RETURN l(1) tc(1)
		.   .   .   RETURN-list
		.   .   .   .   AS l(1) tc(1)
		.   .   .   .   .   NAME-main.~r2 g(1) l(1) x(16) class(PPARAMOUT) bool
		.   .   .   .   .   LITERAL-false l(1) tc(1) bool

		.   RETURN l(1) tc(1)
		.   RETURN-list
		.   .   AS l(1) tc(1)
		.   .   .   NAME-main.~r2 g(1) l(1) x(16) class(PPARAMOUT) bool
		.   .   .   LITERAL-true l(1) tc(1) bool
		enter type..eq.[2]interface {}
		.   AS l(1)
		.   .   NAME-main.~r2 g(1) l(1) x(16) class(PPARAMOUT) bool

	*/
}
