package ads

import "sort"

type (
	User struct {
		Browser string
		Country string
	}

	Campaign struct {
		ClickUrl  string
		Price     float64
		Targeting Targeting
	}

	Targeting struct {
		Browser string
		Country string
	}

	filter func(in []*Campaign, u *User) (out []*Campaign)
)

var (
	filters = []filter{
		filterByBrowser,
		filterByCountry,
	}
)

func MakeAuction(in []*Campaign, u *User) (winner *Campaign) {
	campaigns := make([]*Campaign, len(in))
	copy(campaigns, in)
	for _, fun := range filters {
		fun(campaigns, u)
	}

	if len(campaigns) == 0 {
		return nil
	}

	sort.Slice(campaigns, func(i, j int) bool {
		return campaigns[i].Price > campaigns[j].Price
	})

	return campaigns[0]
}

func filterByBrowser(in []*Campaign, u *User) (out []*Campaign) {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Browser) == 0 {
			continue
		}

		if in[i].Targeting.Browser == u.Browser {
			continue
		}

		in[i] = in[0]
		in = in[1:]
	}
	return in
}

func filterByCountry(in []*Campaign, u *User) (out []*Campaign) {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Country) == 0 {
			continue
		}

		if in[i].Targeting.Country == u.Country {
			continue
		}

		in[i] = in[0]
		in = in[1:]
	}
	return in
}

func GetCampaigns() []*Campaign {
	return []*Campaign{
		{
			Price:     1,
			ClickUrl:  "https://clashofclans.com",
			Targeting: Targeting{Country: "RU"},
		},
		{
			Price:     2,
			ClickUrl:  "https://google.com",
			Targeting: Targeting{Browser: "Chrome"},
		},
		{
			Price:     3,
			Targeting: Targeting{Browser: "Firefox"},
			ClickUrl:  "https://google.kz",
		},
	}
}
