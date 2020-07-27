package Format

import (
	"github.com/newgolibs/godatetime"
	"github.com/newgolibs/godatetime/test/Format/invoke_255_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Invoke_255_0 struct {
}

func TestFormat_invoke_255_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range invoke_255_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Invoke_255_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Invoke_255_0) run(input, arg interface{}) interface{} {
	time_format_str := (godatetime.Format{FormatString: input.(string)}).Invoke()
	return time_format_str
}
