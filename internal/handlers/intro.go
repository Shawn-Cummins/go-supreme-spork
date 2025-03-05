package handlers

import (
	"html"
	"net/http"
)

func (s *ServerImpl) GetIntro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	introText := s.DB.GetIntro().Background.Text
	decodedText := html.UnescapeString(introText)
	w.Write([]byte(decodedText))
}
