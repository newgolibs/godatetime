package DateDiff

import (
    "github.com/newgolibs/godatetime"
    "github.com/newgolibs/godatetime/test/DateDiff/Diffsecondall_261_0_test"
    "log"
    "testing"
    "github.com/stretchr/testify/assert"
	"time"
)

type  Diffsecondall_261_0 struct {

}

func TestDateDiff_Diffsecondall_261_0(t *testing.T){
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    for _, v := range Diffsecondall_261_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Diffsecondall_261_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Diffsecondall_261_0) run(input, arg interface{}) interface{} {
	input2 := input.(string)
	arg2 := arg.(string)
	input3, _ := time.Parse("2006-01-02 15:04:05", input2)
	arg3, _ := time.Parse("2006-01-02 15:04:05", arg2)
	diff := godatetime.DateDiff{From: input3, To: arg3}
	diffyear := diff.Diffsecondall()
	return diffyear
}