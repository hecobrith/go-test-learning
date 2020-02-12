package simplemath

type MathExpr = string

// the magin in somthing inside this, is that you can call diferent functions deppending on the sutations
// expecting different results just by calling in different order
const (
	AddExpr = MathExpr("add")
	DivideExpr = MathExpr("divide")
)
func joinedExp() {
	addExpr := mathExpression(AddExpr)
	println(addExpr(2, 3))
}

func mathExpression( expr MathExpr ) func (float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case DivideExpr:
		return simplemath.Divide
	default: 
		return func(f float64, f2 float64) float64 {
			return 0
		}
	}
}

// so there is an option using the same callback theory... you can pass a function as an argumnet
func double(f1, f2 float64, mathExpr func(float64, float64) float64 ) float64 {
	 return 2 * mathExpr(f1, f2)
}