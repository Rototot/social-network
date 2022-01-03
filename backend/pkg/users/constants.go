package users

//go:generate stringer -type=Gender -output=constants_generated.go
type Gender int8

const (
	Male Gender = iota
	Female
)

//
//func (g Gender) String() string {
//	//TODO implement me
//	return []string{"Male", "Female"}[g]
//}
