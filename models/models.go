package models

type Hub struct {
	ID          string `json:"hudid"`
	Name        string `json:"hubname"`
	GeoLocation string `json:"geolocation"`
}

type Team struct {
	ID       string `json:"teamid"`
	Name     string `json:"teamname"`
	TeamType string `json:"teamtype"`
}

type User struct {
	ID   string `json:"userid"`
	Name string `json:"username"`
	Role string `json:"userrole"`
}

// Inner type role
type Role int

const (
	Admin     Role = iota + 1
	Manager        = 2
	Developer      = 3
	TechLead       = 4
)

// Inner type team
type TeamType int64

const (
	IT    TeamType = iota + 1
	HR             = 1
	BOARD          = 2
)

// Team Response type
type TeamResponse struct {
	Type    string `json:"type"`
	Data    []Team `json:"data"`
	Message string `json:"message"`
}

// User Response type
type UserResponse struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

// Hub Response type
type HubResponse struct {
	Type    string `json:"type"`
	Data    []Hub  `json:"data"`
	Message string `json:"message"`
}
