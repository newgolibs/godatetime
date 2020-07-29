package Diffsecondall_261_0_test

type  Diffsecondall_261_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            "2019-10-01 01:01:00",
            60,
            "2019-10-01 01:00:00",
            " - id=124",
        },
        []interface{}{
            "2019-10-01 01:00:00",
            -61,
            "2019-10-01 01:01:01",
            " - id=125",
        },
	}

}