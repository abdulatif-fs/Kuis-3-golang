package structs

type datetime struct {
	Class    string `json:"_class"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Day      int    `json:"day"`
	Timezone string `json:"timezone"`
	Hour     int    `json:"hour"`
	Minute   int    `json:"minute"`
	Second   int    `json:"second"`
}

type Categories struct {
	ID         int      `json:"id"`
	Nama       string   `json:"nama"`
	Created_at datetime `json:"created_at"`
	Updated_at datetime `json:"updated_at"`
}

type Book struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Image_url    string   `json:"image_url"`
	Release_year int      `json:"release_year"`
	Price        string   `json:"price"`
	Total_page   int      `json:"total_price"`
	Thickness    string   `json:"thickness"`
	Created_at   datetime `json:"created_at"`
	Updated_at   datetime `json:"updated_at"`
	Category_id  int      `json:"category_id"`
}
