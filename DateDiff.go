package godatetime

import (
	"math"
)

/**    相差年份    */
func (this DateDiff) Diffyear() int {
	return this.From.Year() - this.To.Year()
}

/**    相差的总天数    */
func (this DateDiff) Diffdayall() int {
	sub := this.From.Sub(this.To)
	return int(math.Floor(sub.Hours() / 24))
}

/**    相差的总小时    */
func (this DateDiff) Diffhourall() int {
	sub := this.From.Sub(this.To)
	return int(math.Floor(sub.Hours()))
}

/**    相差的总秒数    */
func (this DateDiff)Diffsecondall()int{
	sub := this.From.Sub(this.To)
	return int(math.Floor(sub.Seconds()))
}

