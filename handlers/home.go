package handlers

import (
	"church/render"
	"log"
	"net/http"
)

type HomeData struct {
	ChurchName  string
	PastorsName string
	BibleVerse  string
	Activity    []Activity
}
type Activity struct {
	Day      string
	Service  string
	Time     string
	Preacher string
	Venue    string
	IsLive   bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		ChurchName:  "RCCG VICTORY PARISH",
		PastorsName: "Pastor E.A. Adeboye",
		BibleVerse:  "Hebrew 13:8",
		Activity: []Activity{
			{
				Day:      "Tuesday",
				Service:  "Digging Deep",
				Time:     "5:00 pm - 6:30 pm",
				Preacher: "Dc. Vincent Ayoka",
				Venue:    "Sunrise Parish",
				IsLive:   true,
			},
			{
				Day:      "Wednesday",
				Service:  "Dominion Hour",
				Time:     "5:00 pm - 6:30 pm",
				Preacher: "Picp. Isaac Iyeresi",
				Venue:    "Sunrise Parish",
				IsLive:   true,
			},
			{
				Day:      "Thursday",
				Service:  "Faith Clinic",
				Time:     "5:00 pm - 6:00 pm",
				Preacher: "Dcn. Onyeche Omale",
				Venue:    "Sunrise Parish",
				IsLive:   false,
			},
			{
				Day:      "Sunday",
				Service:  "Worship & Thanksgivingservice ",
				Time:     "8:00 am - 11:30 am",
				Preacher: "Pst. Isaac Iyasere",
				Venue:    "Sunrise Parish",
				IsLive:   true,
			},
		},
	}
	log.Println("User Visited Home Page")
	if err := render.Render(w, "home.html", data); err != nil {
		log.Println("Render error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
