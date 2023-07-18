package server

import (
	"encoding/json"
	"io"
	"math/rand"
	"net"
	"strings"

	"github.com/rs/zerolog"
)

type BookQuote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type Book struct {
	Quotes []*BookQuote
}

func NewBook(jsonQuotes []byte) (*Book, error) {
	var quotes []*BookQuote
	if err := json.Unmarshal(jsonQuotes, &quotes); err != nil {
		return nil, err
	}
	return &Book{Quotes: quotes}, nil
}

func (b *Book) GetRandQuote() *BookQuote {
	i := rand.Intn(len(b.Quotes))
	return b.Quotes[i]
}

func (b *Book) ServeRequest(conn net.Conn, requestLog zerolog.Logger) {
	requestLog.Info().Msg("write response")

	quote := b.GetRandQuote()
	quoteReader := strings.NewReader(quote.Quote)
	_, err := io.Copy(conn, quoteReader)
	if err != nil {
		requestLog.Warn().Err(err).Msg("failed to write response")
	}
}
