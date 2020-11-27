package sqlite

// Chat represents a db record for chat.
type Chat struct {
	ID            int
	Type          string
	Title         string
	Username      string
	FirstName     string
	LastName      string
	Description   string
	PinnedMessage string
	IsActive      bool
}

// CatJoke represents a db record for cat_jokes table.
type CatJoke struct {
	ID   int
	Text string
	Day  string
}

// Poetry represents a db record for poetry table.
type Poetry struct {
	ID   int
	Text string
	Type int
}
