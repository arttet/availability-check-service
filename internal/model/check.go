package model

type StatusType = string

const (
	Ok   StatusType = "ok"
	Fail StatusType = "fail"
)

type Check struct {
	ID          int64      `db:"id"`
	Host        string     `db:"host"`
	Port        uint16     `db:"port"`
	Status      StatusType `db:"status"`
	Timeout     uint       `db:"timeout"`
	FailMessage string     `db:"fail_message"`
}

type Checks []*Check

func GetStatus(ok bool) StatusType {
	if ok {
		return Ok
	}
	return Fail
}
