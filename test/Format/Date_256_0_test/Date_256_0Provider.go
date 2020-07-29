package Date_256_0_test

import (
    "time"
)

type  Date_256_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            "",
            time.Now().Format("2006-01-02"),
            "",
            " - id=115",
        },
	}

}