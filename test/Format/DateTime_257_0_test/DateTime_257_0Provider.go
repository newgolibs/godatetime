package DateTime_257_0_test

import (
    "time"
)

type  DateTime_257_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            "",
            time.Now().Format("2006-01-02 15:04:05"),
            "",
            " - id=116",
        },
	}

}