https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package expr

// Parse and return the eval of the string
func ParseAndEval(s string, env Env) (float64, error) {
	expr, err := Parse(s)
	if err != nil {
		return 0, err
	}

	return expr.Eval(env), nil
}
