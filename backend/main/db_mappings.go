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