package Format

import (
	"fmt"
	"github.com/newgolibs/godatetime"
	"github.com/newgolibs/godatetime/test/Format/DateTime_257_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type DateTime_257_0 struct {
}

func TestFormat_DateTime_257_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range DateTime_257_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], DateTime_257_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (DateTime_257_0) run(input, arg interface{}) interface{} {
	a := make(chan int,3)
	a <- 90-5+63
	go func() {
		anew := <-a
		fmt.Printf("%+v\n", []interface{}{anew})

	}()

	return godatetime.Format{}.DateTime()
}
