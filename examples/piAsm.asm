"".Add t=1 size=36 args=0x18 locals=0x0 leaf
	0x0000 00000 (floatAdd.go:3)	TEXT	"".Add(SB), $-4-24
	0x0000 00000 (floatAdd.go:3)	FUNCDATA	$0, gclocals·709a14768fab2805a378215c02f0d27f(SB)
	0x0000 00000 (floatAdd.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (floatAdd.go:3)	MOVD	$f64.0000000000000000(SB), F0
	0x0008 00008 (floatAdd.go:4)	MOVD	"".x(FP), F0
	0x000c 00012 (floatAdd.go:4)	MOVD	"".y+8(FP), F1
	0x0010 00016 (floatAdd.go:4)	ADDD	F1, F0
	0x0014 00020 (floatAdd.go:4)	MOVD	F0, "".~r2+16(FP)
	0x0018 00024 (floatAdd.go:4)	JMP	(R14)
	0x001c 00028 (floatAdd.go:4)	JMP	0(PC)
	0x0020 00032 (floatAdd.go:4)	WORD	$f64.0000000000000000(SB)
	0x0000 18 b0 9f e5 00 0b 9b ed 01 0b 9d ed 03 1b 9d ed  ................
	0x0010 01 0b 30 ee 05 0b 8d ed 00 f0 8e e2 fe ff ff ea  ..0.............
	0x0020 00 00 00 00                                      ....
	rel 32+4 t=1 $f64.0000000000000000+0
"".init t=1 size=108 args=0x0 locals=0x0
	0x0000 00000 (floatAdd.go:7)	TEXT	"".init(SB), $0-0
	0x0000 00000 (floatAdd.go:7)	MOVW	8(g), R1
	0x0004 00004 (floatAdd.go:7)	CMP	R1, R13
	0x0008 00008 (floatAdd.go:7)	BLS	88
	0x000c 00012 (floatAdd.go:7)	MOVW.W	R14, -4(R13)
	0x0010 00016 (floatAdd.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0010 00016 (floatAdd.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0010 00016 (floatAdd.go:7)	MOVBU	"".initdone·(SB), R0
	0x0018 00024 (floatAdd.go:7)	CMP	$1, R0
	0x001c 00028 (floatAdd.go:7)	BLS	$0, 36
	0x0020 00032 (floatAdd.go:7)	MOVW.P	4(R13), R15
	0x0024 00036 (floatAdd.go:7)	MOVBU	"".initdone·(SB), R0
	0x002c 00044 (floatAdd.go:7)	CMP	$1, R0
	0x0030 00048 (floatAdd.go:7)	BNE	$0, 60
	0x0034 00052 (floatAdd.go:7)	PCDATA	$0, $0
	0x0034 00052 (floatAdd.go:7)	CALL	runtime.throwinit(SB)
	0x0038 00056 (floatAdd.go:7)	UNDEF
	0x003c 00060 (floatAdd.go:7)	MOVW	$1, R0
	0x0040 00064 (floatAdd.go:7)	MOVB	R0, "".initdone·(SB)
	0x0048 00072 (floatAdd.go:7)	MOVW	$2, R0
	0x004c 00076 (floatAdd.go:7)	MOVB	R0, "".initdone·(SB)
	0x0054 00084 (floatAdd.go:7)	MOVW.P	4(R13), R15
	0x0058 00088 (floatAdd.go:7)	NOP
	0x0058 00088 (floatAdd.go:7)	MOVW	R14, R3
	0x005c 00092 (floatAdd.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0060 00096 (floatAdd.go:7)	JMP	0
	0x0064 00100 (floatAdd.go:7)	JMP	0(PC)
	0x0068 00104 (floatAdd.go:7)	WORD	"".initdone·(SB)
	0x0000 08 10 9a e5 01 00 5d e1 12 00 00 9a 04 e0 2d e5  ......].......-.
	0x0010 50 b0 9f e5 00 00 db e5 01 00 50 e3 00 00 00 9a  P.........P.....
	0x0020 04 f0 9d e4 3c b0 9f e5 00 00 db e5 01 00 50 e3  ....<.........P.
	0x0030 01 00 00 1a 00 00 00 eb fd bc fa f7 01 00 a0 e3  ................
	0x0040 20 b0 9f e5 00 00 cb e5 02 00 a0 e3 14 b0 9f e5   ...............
	0x0050 00 00 cb e5 04 f0 9d e4 0e 30 a0 e1 00 00 00 eb  .........0......
	0x0060 e6 ff ff ea fe ff ff ea 00 00 00 00              ............
	rel 52+4 t=8 runtime.throwinit+ebfffffe
	rel 92+4 t=8 runtime.morestack_noctxt+ebfffffe
	rel 104+4 t=1 "".initdone·+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb t=9 dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·709a14768fab2805a378215c02f0d27f t=9 dupok size=12
	0x0000 01 00 00 00 06 00 00 00 00 00 00 00              ............
"".initdone· t=34 size=1
"".Add·f t=9 dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 "".Add+0
"".init·f t=9 dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 "".init+0