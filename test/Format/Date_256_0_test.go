package Format

import (
	"fmt"
	"github.com/newgolibs/godatetime"
	"github.com/newgolibs/godatetime/test/Format/Date_256_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"testing"
	"time"
)

type Date_256_0 struct {
}

func TestFormat_Date_256_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Date_256_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Date_256_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Date_256_0) run(input, arg interface{}) interface{} {
	t1 := time.Now()
	t2 ,_:= time.Parse("2006-01-02 15:04:05", "2020-07-30 01:01:01")
	fmt.Printf("%+v\n", []interface{}{t1,"xxxx", t2})
	sub := t1.Sub(t2)
	fmt.Printf("%+v\n", []interface{}{sub,math.Ceil(sub.Hours()),math.Floor(sub.Hours()),math.Floor(sub.Minutes())})
	return godatetime.Format{}.Date()
}
