package andrea

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)


type Scanner struct {
	r *bufio.Reader
}


func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}


func (s *Scanner) Scan() (tok Token, lit string) {

	ch := s.read()

	
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case ':':
		return DosPunt, string(ch)
	case '(':
		return ParDer, string(ch)
	case ')':
		return ParIsq, string(ch)
	case ';':
		return PuntCom, string(ch)
	case '{':
		return LlaveDer, string(ch)
	case '}':
		return LlaveIsq, string(ch)
	case '\n':
		return SaltoDeL, string(ch)
	case '=':
		return Asignaicon, string(ch)
	case '[':
		return corchetederecho, string(ch)
	case ']':
		return corcheteizquierdo, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	case "si":
		return si, buf.String()
	case "no":
		return no, buf.String()
	case "zend":
		return zend, buf.String()
	case "sw":
		return sw, buf.String()
	case "zmain":
		return zmain, buf.String()
	case "funcn":
		return funcn, buf.String()
	case "para":
		return para, buf.String()
	case "zcase":
		return zcase, buf.String()
	case "vars":
		return vars, buf.String()
	case "break":
		return breaks, buf.String()
	case "while":
		return while, buf.String()
	case "du":
		return du, buf.String()
	case "zINT":
		return zINT, buf.String()
	case "tri":
		return tri, buf.String()
	case "cach":
		return cach, buf.String()
	case "zFINALLY":
		return zFINALLY, buf.String()
	case "zsize":
		return zsize, buf.String()
	case "znew":
		return znew, buf.String()
	case "zout":
		return zout, buf.String()
	case "zUNCHECKED":
		return zUNCHECKED, buf.String()
	case "feach":
		return feach, buf.String()
	case "zthis":
		return zthis, buf.String()
	}

	return IDENT, buf.String()

}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
