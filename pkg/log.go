package pkg

type Level int

const (
	Debug Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {

}
