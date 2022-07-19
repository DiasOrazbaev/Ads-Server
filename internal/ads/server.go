package ads

import (
	"fmt"
	realip "github.com/ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
	"log"
	"net"
)

type Server struct {
	geoip2 *geoip2.Reader
}

func NewServer(geoip2 *geoip2.Reader) *Server {
	return &Server{geoip2: geoip2}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":4040", s.handleHttp)
}

func (s *Server) handleHttp(ctx *fasthttp.RequestCtx) {
	ua := string(ctx.Request.Header.UserAgent())
	parsedUA := user_agent.New(ua)
	browserName, browserVersion := parsedUA.Browser()

	ip := realip.FromRequest(ctx)
	country, err := s.geoip2.Country(net.ParseIP(ip))
	if err != nil {
		log.Panicln(err)
	}

	u := User{Browser: browserName, Country: country.Country.IsoCode}
	campaigns := GetCampaigns()
	winner := MakeAuction(campaigns, &u)
	//if winner == nil {
	//	ctx.Redirect("https://yandex.kz", http.StatusSeeOther)
	//}
	//
	//ctx.Redirect(winner.ClickUrl, http.StatusFound)

	ctx.WriteString(fmt.Sprintf("User-Agent: %s\n", ua))
	ctx.WriteString(fmt.Sprintf("Browser name: %s\n", browserName))
	ctx.WriteString(fmt.Sprintf("Broser version: %s\n", browserVersion))
	ctx.WriteString(fmt.Sprintf("IP: %s\n", ip))
	ctx.WriteString(fmt.Sprintf("Country: %s\n", country.Country.IsoCode))
	ctx.WriteString(fmt.Sprintf("Winned: %v\n", winner))

}
