package Diffyear_258_0_test

type  Diffyear_258_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            "2020-10-02",
            1,
            "2019-10-02",
            " - id=117",
        },
        []interface{}{
            "2019-10-02",
            -1,
            "2020-10-02",
            " - id=118",
        },
	}

}