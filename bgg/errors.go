package bgg

type RateLimitError struct {
	code int
}

func (e RateLimitError) Error() string {
	return "rate limit error received from bgg"
}

type NoResultFoundError struct{}

func (e NoResultFoundError) Error() string {
	return "no result found"
}
