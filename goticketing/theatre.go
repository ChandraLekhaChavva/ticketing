package goticketing

// Theatre : describe what this function does
type Theatre struct {
	ID               int
	Screens          []Screen
	TotalTicketsSold int
}

// Show : describe what this function does
type Show struct {
	ID        int
	NoOfSeats int
}

// Screen : describe what this function does
type Screen struct {
	ID    int
	Shows []Show
}

// Cafeteria : describe what this function does
type Cafeteria struct {
	NoOfSodas        int
	TotalSodasSold   int
	TotalPopcornSold int
	TotalWaterSold   int
}

// Receipt : describe what this function does
type Receipt struct {
	ScreenNumber int
	ShowNumber   int
	PersonNumber int
}

// GetCafeteriaMoq : Mock data for Cafeteria
func GetCafeteriaMoq() Cafeteria {
	cafeteria := Cafeteria{
		NoOfSodas: 10,
	}
	return cafeteria
}

// GetTheatreMoq : Mock data for theatre
func GetTheatreMoq() Theatre {
	theatre := Theatre{
		ID: 1,
		Screens: []Screen{
			{ID: 1, Shows: []Show{
				{ID: 1, NoOfSeats: 2},
				{ID: 2, NoOfSeats: 2},
				{ID: 3, NoOfSeats: 2},
				{ID: 4, NoOfSeats: 2},
			}},
			{ID: 2, Shows: []Show{
				{ID: 1, NoOfSeats: 2},
				{ID: 2, NoOfSeats: 2},
				{ID: 3, NoOfSeats: 2},
				{ID: 4, NoOfSeats: 2},
			}},
			{ID: 3, Shows: []Show{
				{ID: 1, NoOfSeats: 2},
				{ID: 2, NoOfSeats: 2},
				{ID: 3, NoOfSeats: 2},
				{ID: 4, NoOfSeats: 2},
			}},
			{ID: 4, Shows: []Show{
				{ID: 1, NoOfSeats: 2},
				{ID: 2, NoOfSeats: 2},
				{ID: 3, NoOfSeats: 2},
				{ID: 4, NoOfSeats: 2},
			}},
			{ID: 5, Shows: []Show{
				{ID: 1, NoOfSeats: 2},
				{ID: 2, NoOfSeats: 2},
				{ID: 3, NoOfSeats: 2},
				{ID: 4, NoOfSeats: 2},
			}},
		},
	}
	return theatre
}
