package utility

// IfThenElse - this function returns a value depending on given condition which evaluates to
// boolean value. it accepts
// condition - a condition which must evaluate to boolean
// it returns
// a - value given as argument when condition is true
// b - value given as argument when condition is false
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
