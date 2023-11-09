package inter

import (
	"crypto/rsa"
	"cryptopus"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTicketSeller struct {
	signingMethod jwt.SigningMethod
	key           interface{}
	parser        jwt.Parser
}

var (
	ket, err = cryptopus.ReadKeyFromFile("D:\\GOPHER\\goframe\\goframe\build\\key.pem")

	jwtTicketSellerExample JwtTicketSeller = JwtTicketSeller{
		signingMethod: jwt.SigningMethodPS512,
		key:           ket,
		parser: *jwt.NewParser(
			func(p *jwt.Parser) {
				b := p.validMethods
			},
		),
	}
)

type JwtTicket struct {
	Token    *jwt.Token
	StrToken *string
}

func (t *JwtTicket) TicketToModel() interface{} {
	// TODO: придумать модель тикета в базе данных
	return nil
}
func (t *JwtTicket) IsEqual(ticket *Ticket) bool {
	g, ok := (*ticket).(*JwtTicket)
	if !ok {
		return false
	}

	return t.StrToken == g.StrToken
}
func (t *JwtTicket) TicketToResponse() interface{} {
	// TODO: придумать модель ответа
	return nil
}

func (jts *JwtTicketSeller) BuyTicket(u *User) *Ticket {
	t := jwt.NewWithClaims(jts.signingMethod, jwt.MapClaims{
		"nickname":  u.Nickname,
		"join_at":   u.JoinAt.String(),
		"signed_at": time.Now().String(),
	})
	s, err := t.SignedString(jts.key)
	if err != nil {
		return nil
	}
	var ticketToUser Ticket = &JwtTicket{
		Token:    t,
		StrToken: &s,
	}
	return &ticketToUser
}

func (jts *JwtTicketSeller) IsValid(t *Ticket) bool {

}

func luckyMe() *jwt.Token {

	var (
		key *rsa.PrivateKey
		tkn *jwt.Token
		s   string
	)

	tkn = jwt.New(jwt.SigningMethodRS512)
	s, _ = tkn.SignedString(key)

	jwt.Parse()

	return jwt.New(jwt.SigningMethodRS512)
}
