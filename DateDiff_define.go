package godatetime
import (
        "time"
)

//对象必须实现的接口方法
type DateDiffInterface interface {
    /**    相差的总天数    */
    Diffdayall()int
    /**    相差的总小时    */
    Diffhourall()int
    /**    相差的总秒数    */
    Diffsecondall()int
    /**    相差年份    */
    Diffyear()int

}

/**
比较日期的差距;
*/
type DateDiff struct
{
    /*开始的时间*/
    From time.Time
    /*结束的时间*/
    To time.Time
}








//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ DateDiffInterface =new(DateDiff)
