package main

type ServerError string

func (e ServerError) Error() string {
    if e == "" {
        return "Internal server error"
    }
    return string(e)
}


var NOT_IMPLEMENTED ServerError = "Not implemented"
