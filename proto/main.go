package proto

// recommendation proto
type RecommendationRequest struct {
	Require string
	Lat     float64 
	Lon     float64 
}

type RecommendationResponse struct {
	HotelIDs []string
}


// profile proto

type ProfileRequest struct {
	Lat float64
	Lon float64
}

type Address struct {
	StreetNumber string
	StreetName   string
	City         string
	State        string
	Country      string
	PostalCode   string
	Lat          float64
	Lon          float64
}

type Image struct {
	Url string
	Def bool
}

type Hotel struct {
	ID          string
	Name        string
	PhoneNumber string
	Description string
	Address     Address
	Images      []Image
}

type ProfileResponse struct {
	Hotels []Hotel
}
