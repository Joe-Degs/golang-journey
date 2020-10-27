package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// first we define our tokens.
// Token represents a lexical token
type Token int

// tokens represent series of characters
// they help us make sense out of data the lexers
// recieves
const (
	// special tokens
	ILLEGAL Token = iota
	EOF
	WS // whitespace

	// literals
	IDENT // fields, table_name

	// misc characters
	ASTERISK // *
	COMMA    // ,

	// keywords
	SELECT
	FROM
)

// character classes helps us group characters into appropriate types.
// we need to know if a character is whitespace or a letter

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\t'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// it is useful to define an 'EOF' character so we treat
// as any other character
var eof = rune(0)

// time for serious business
// A scanner type wraps around an input reader and provides us with basic
// functionality for reading, unreading, peeking ahead from the

// Scanner represents a lexical scanner
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new lexical scanner
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read returns the next rune from the buffered reader
// returns eof if error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back in the reader
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value
func (s *Scanner) Scan() (tok Token, lit string) {
	// read next rune
	ch := s.read()

	// if its a whitespace? scan all contigous whitespace
	// if its a letter consume an ident or reserved word
	if isWhiteSpace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// otherwise read individual character
	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanning contigous characters

// scanWhiteSpace scans current rune and all contigous whitespace.
func (s *Scanner) scanWhiteSpace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// read every subsequent whitespace character into buffer.
	// non-whitespace characters and eof will cause the loop to exit
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhiteSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent scans current rune and all contiguous ident rune
func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

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

	// if string matches a keyword then return the that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}
	// otherwise return as a regular identifier.
	return IDENT, buf.String()
}

// Parsing Input
// SelectStatement is an AST for select statements.
type SelectStatement struct {
	Fields    []string
	TableName string
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size(max=1)
	}
}

// NewParser returns a ready to use parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// if token is unscanned read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// if we have a token on the buffer we return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// otherwise read the next token from *Scanner
	tok, lit = p.s.Scan()
	// save it to the buffer in case we unscan later
	p.buf.tok, p.buf.lit = tok, lit
	return
}

// unscan pushes the previously read token back in the buffer
func (p *Parser) unscan() { p.buf.n = 1 }

// Parser doesnt care about whitespace so we find the next
// non-whitespace character
func (p *Parser) scanIgnoreWhiteSpace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// Parse returns an AST of the next SELECT statement from reader.
func (p *Parser) Parse() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	if tok, lit := p.scan(); tok != SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}

	// continously read comma delimited fields.
	for {
		// read a field.
		tok, lit := p.scanIgnoreWhiteSpace()
		if tok != IDENT && tok != ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}
		stmt.Fields = append(stmt.Fields, lit)

		// if next token is not comma break the loop.
		if tok, _ := p.scanIgnoreWhiteSpace(); tok != COMMA {
			p.unscan()
			break
		}
	}

	// next we want to see a FROM keyword.
	if tok, lit := p.scanIgnoreWhiteSpace(); tok != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	// next we want to get table name, it should be an IDENT token.
	tok, lit := p.scanIgnoreWhiteSpace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	return stmt, nil
}
