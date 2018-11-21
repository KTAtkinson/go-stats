package collector

type StatsError string

func (e StatsError) Error() string {
	if e == "" {
		return "Error with statistics"
	}

	return string(e)
}

var NOT_IMPLEMENTED StatsError = "Not implemented"
