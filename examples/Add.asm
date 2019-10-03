"".Add STEXT nosplit size=23 args=0x18 locals=0x0
TEXT	"".Add(SB), NOSPLIT|ABIInternal, $0-24
FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
PCDATA	$2, $0
PCDATA	$0, $0
MOVSD	"".x+8(SP), X0
MOVSD	"".y+16(SP), X1
ADDSD	X1, X0
MOVSD	X0, "".~r2+24(SP)
RET