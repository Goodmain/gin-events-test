package ticketmaster

import (
	"encoding/json"
	"events-hackathon-go/core/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type TicketmasterImage struct {
	Ratio    string `json:"ratio" bson:"ratio"`
	URL      string `json:"url" bson:"url"`
	Width    int    `json:"width" bson:"width"`
	Height   int    `json:"height"  bson:"height"`
	Fallback bool   `json:"fallback" bson:"fallback"`
}

type TicketmasterEvent struct {
	Name   string              `json:"name" bson:"name"`
	Type   string              `json:"type" bson:"type"`
	ID     string              `json:"id" bson:"id"`
	Test   bool                `json:"test" bson:"test"`
	URL    string              `json:"url" bson:"url"`
	Locale string              `json:"locale" bson:"locale"`
	Images []TicketmasterImage `json:"images" bson:"images"`
	Sales  struct {
		Public struct {
			StartDateTime time.Time `json:"startDateTime" bson:"startDateTime"`
			StartTBD      bool      `json:"startTBD" bson:"startTBD"`
			StartTBA      bool      `json:"startTBA" bson:"startTBA"`
			EndDateTime   time.Time `json:"endDateTime" bson:"endDateTime"`
		} `json:"public" bson:"public"`
	} `json:"sales" bson:"sales"`
	Dates struct {
		Start struct {
			LocalDate      string    `json:"localDate" bson:"localDate"`
			LocalTime      string    `json:"localTime" bson:"localTime"`
			DateTime       time.Time `json:"dateTime" bson:"dateTime"`
			DateTBD        bool      `json:"dateTBD" bson:"dateTBD"`
			DateTBA        bool      `json:"dateTBA"  bson:"dateTBA"`
			TimeTBA        bool      `json:"timeTBA"  bson:"timeTBA"`
			NoSpecificTime bool      `json:"noSpecificTime" bson:"noSpecificTime"`
		} `json:"start" bson:"start"`
		Status struct {
			Code string `json:"code" bson:"code"`
		} `json:"status" bson:"status"`
		SpanMultipleDays bool `json:"spanMultipleDays" bson:"spanMultipleDays"`
	} `json:"dates" bson:"dates"`
	Classifications []struct {
		Primary bool `json:"primary" bson:"primary"`
		Segment struct {
			ID   string `json:"id" bson:"id"`
			Name string `json:"name" bson:"name"`
		} `json:"segment" bson:"segment"`
		Genre struct {
			ID   string `json:"id" bson:"id"`
			Name string `json:"name" bson:"name"`
		} `json:"genre" bson:"genre"`
		SubGenre struct {
			ID   string `json:"id" bson:"id"`
			Name string `json:"name" bson:"name"`
		} `json:"subGenre" bson:"subGenre"`
		Family bool `json:"family" bson:"family"`
	} `json:"classifications" bson:"classifications"`
	Outlets []struct {
		URL  string `json:"url" bson:"url"`
		Type string `json:"type" bson:"type"`
	} `json:"outlets" bson:"outlets"`
	Seatmap struct {
		StaticURL string `json:"staticUrl" bson:"staticUrl"`
	} `json:"seatmap" bson:"seatmap"`
	Links struct {
		Self struct {
			Href string `json:"href" bson:"href"`
		} `json:"self" bson:"self"`
		Attractions []struct {
			Href string `json:"href" bson:"href"`
		} `json:"attractions" bson:"attractions"`
		Venues []struct {
			Href string `json:"href" bson:"href"`
		} `json:"venues" bson:"venues"`
	} `json:"_links" bson:"_links"`
	Embedded struct {
		Venues []struct {
			Name       string `json:"name" bson:"name"`
			Type       string `json:"type" bson:"type"`
			ID         string `json:"id" bson:"id"`
			Test       bool   `json:"test" bson:"test"`
			Locale     string `json:"locale" bson:"locale"`
			PostalCode string `json:"postalCode" bson:"postalCode"`
			Timezone   string `json:"timezone" bson:"timezone"`
			City       struct {
				Name string `json:"name" bson:"name"`
			} `json:"city" bson:"city"`
			State struct {
				Name      string `json:"name" bson:"name"`
				StateCode string `json:"stateCode" bson:"stateCode"`
			} `json:"state" bson:"state"`
			Country struct {
				Name        string `json:"name" bson:"name"`
				CountryCode string `json:"countryCode" bson:"countryCode"`
			} `json:"country" bson:"country"`
			Address struct {
				Line1 string `json:"line1" bson:"line1"`
				Line2 string `json:"line2" bson:"line2"`
			} `json:"address" bson:"address"`
			Location struct {
				Longitude string `json:"longitude" bson:"longitude"`
				Latitude  string `json:"latitude" bson:"latitude"`
			} `json:"location" bson:"location"`
			Dmas []struct {
				ID int `json:"id" bson:"id"`
			} `json:"dmas" bson:"dmas"`
			UpcomingEvents struct {
				Total        int `json:"_total" bson:"_total"`
				Tmr          int `json:"tmr" bson:"tmr"`
				Ticketmaster int `json:"ticketmaster" bson:"ticketmaster"`
				Filtered     int `json:"_filtered" bson:"_filtered"`
			} `json:"upcomingEvents" bson:"upcomingEvents"`
			Links struct {
				Self struct {
					Href string `json:"href" bson:"href"`
				} `json:"self" bson:"self"`
			} `json:"_links" bson:"_links"`
		} `json:"venues" bson:"venues"`
		Attractions []struct {
			Name          string `json:"name" bson:"name"`
			Type          string `json:"type" bson:"type"`
			ID            string `json:"id" bson:"id"`
			Test          bool   `json:"test" bson:"test"`
			URL           string `json:"url" bson:"url"`
			Locale        string `json:"locale" bson:"locale"`
			ExternalLinks struct {
				Twitter []struct {
					URL string `json:"url" bson:"url"`
				} `json:"twitter" bson:"twitter"`
				Wiki []struct {
					URL string `json:"url" bson:"url"`
				} `json:"wiki" bson:"wiki"`
				Facebook []struct {
					URL string `json:"url" bson:"url"`
				} `json:"facebook" bson:"facebook"`
				Instagram []struct {
					URL string `json:"url" bson:"url"`
				} `json:"instagram" bson:"instagram"`
				Homepage []struct {
					URL string `json:"url" bson:"url"`
				} `json:"homepage" bson:"homepage"`
			} `json:"externalLinks" bson:"externalLinks"`
			Aliases         []string            `json:"aliases" bson:"aliases"`
			Images          []TicketmasterImage `json:"images" bson:"images"`
			Classifications []struct {
				Primary bool `json:"primary" bson:"primary"`
				Segment struct {
					ID   string `json:"id" bson:"id"`
					Name string `json:"name" bson:"name"`
				} `json:"segment" bson:"segment"`
				Genre struct {
					ID   string `json:"id" bson:"id"`
					Name string `json:"name" bson:"name"`
				} `json:"genre" bson:"genre"`
				SubGenre struct {
					ID   string `json:"id" bson:"id"`
					Name string `json:"name" bson:"name"`
				} `json:"subGenre" bson:"subGenre"`
				Type struct {
					ID   string `json:"id" bson:"id"`
					Name string `json:"name" bson:"name"`
				} `json:"type" bson:"type"`
				SubType struct {
					ID   string `json:"id" bson:"id"`
					Name string `json:"name" bson:"name"`
				} `json:"subType" bson:"subType"`
				Family bool `json:"family" bson:"family"`
			} `json:"classifications" bson:"classifications"`
			UpcomingEvents struct {
				Total        int `json:"_total" bson:"_total"`
				Tmr          int `json:"tmr" bson:"tmr"`
				Ticketmaster int `json:"ticketmaster" bson:"ticketmaster"`
				Filtered     int `json:"_filtered" bson:"_filtered"`
			} `json:"upcomingEvents" bson:"upcomingEvents"`
			Links struct {
				Self struct {
					Href string `json:"href" bson:"href"`
				} `json:"self" bson:"self"`
			} `json:"_links" bson:"_links"`
		} `json:"attractions" bson:"attractions"`
	} `json:"_embedded" bson:"_embedded"`
}

type TicketmasterEventList struct {
	Embedded struct {
		Events []TicketmasterEvent `json:"events" bson:"events"`
	} `json:"_embedded" bson:"_embedded"`
	Links struct {
		First struct {
			Href string `json:"href" bson:"href"`
		} `json:"first" bson:"first"`
		Self struct {
			Href string `json:"href" bson:"href"`
		} `json:"self" bson:"self"`
		Next struct {
			Href string `json:"href" bson:"href"`
		} `json:"next" bson:"next"`
		Last struct {
			Href string `json:"href" bson:"href"`
		} `json:"last" bson:"last"`
	} `json:"_links" bson:"_links"`
	Page struct {
		Size          int `json:"size" bson:"size"`
		TotalElements int `json:"totalElements" bson:"totalElements"`
		TotalPages    int `json:"totalPages" bson:"totalPages"`
		Number        int `json:"number" bson:"number"`
	} `json:"page" bson:"page"`
}

func (e *TicketmasterEvent) getImage(ratio string, width int) TicketmasterImage {
	for _, image := range e.Images {
		if image.Width == width && image.Ratio == ratio {
			return image
		}
	}

	return TicketmasterImage{URL: ""}
}

func (t *TicketmasterEventList) convertToEvents() []models.Event {
	var events []models.Event

	for _, data := range t.Embedded.Events {
		image := data.getImage("3_2", 1024)
		thumbnail := data.getImage("3_2", 640)

		event := models.Event{
			InitialID: data.ID,
			Name:      data.Name,
			URL:       data.URL,
			Image:     image.URL,
			Thumbnail: thumbnail.URL,
			City:      data.Embedded.Venues[0].City.Name,
			Place:     data.Embedded.Venues[0].Name,
			Address:   data.Embedded.Venues[0].Address.Line1,
			State:     data.Embedded.Venues[0].State.StateCode,
			Country:   data.Embedded.Venues[0].Country.CountryCode,
			Zip:       data.Embedded.Venues[0].PostalCode,
			Lng:       data.Embedded.Venues[0].Location.Longitude,
			Lat:       data.Embedded.Venues[0].Location.Latitude,
			Date:      data.Dates.Start.DateTime,
		}

		events = append(events, event)
	}

	return events
}

func LoadEvents(city string) ([]models.Event, bool) {
	var events []models.Event

	for i := 20; i >= 1; i-- {
		eventsList := parseEvents(city, i)

		if eventsList.Page.Size == 0 {
			break
		}

		events = append(events, eventsList.convertToEvents()...)
	}

	if len(events) == 0 {
		return nil, false
	}

	return events, true
}

func parseEvents(city string, page int) TicketmasterEventList {
	url := viper.Get("EVENTS_URL").(string) + "&countryCode=US&city=" + url.QueryEscape(city) + "&page=" + strconv.Itoa(page)
	response, error := http.Get(url)

	if error != nil {
		log.Println(error)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var result TicketmasterEventList
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}
