package main

type user struct {
	Id string `json:"id,omitempty"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Image string `json:"image"`
	Group string `json:"group,omitempty"`
}

type event struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Location string `json:"location"`
}

type like struct {
	LikeId string `json:"like_id,omitempty"`
	EventId string `json:"event_id"`
	FirstUserId string `json:"first_user_id"`
	SecondUserId string `json:"second_user_id"`
}