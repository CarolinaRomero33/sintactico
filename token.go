package andrea

// Token represents a lexical token.
type Token int

const (
	
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK      // *
	COMMA         // ,
	DosPunt       // :
	ParDer        // (
	ParIsq        // )
	PuntCom       // ;
	LlaveDer      // {
	LlaveIsq      // }
	SaltoDeL      //n
	Asignaicon    // =
	corchetederecho // [
	corcheteizquierdo // ]

	// Keywords
	SELECT
	FROM
	si
	sw
	zend
	zINT // Tipodedato
	funcn
	zmain
	no
	para
	zcase
	vars // variables
	breaks
	while
	
	du
	zUNCHECKED
	
	
	cach
	zFINALLY
	zsize
	znew
	zout
	feach
	clase
	zthis
)
