package invoke_255_0_test

import (
    "time"
)

type  Invoke_255_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            "Y-m-d H:i:s",
            time.Now().Format("2006-01-02 15:04:05"),
            "",
            "测试普通的Y-m-d H:i:s格式化效果 - id=113",
        },
        []interface{}{
            "Y-m-d",
            time.Now().Format("2006-01-02"),
            "",
            "测试普通的Y-m-d H:i:s格式化效果 - id=114",
        },
	}

}