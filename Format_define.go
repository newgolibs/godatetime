package godatetime
import (
        "time"
)

//对象必须实现的接口方法
type FormatInterface interface {
    /**    返回格式完的字符串    */
    Invoke()string

}

/**
格式化一个指定的时间，传入的是time格式;
*/
type Format struct
{
    /*格式化分割，参考php*/
    FormatString string
    /*Time对象*/
    Time time.Time
}








//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ FormatInterface =new(Format)
