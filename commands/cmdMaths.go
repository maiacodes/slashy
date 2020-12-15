package commands

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/maiacodes/slashy/interaction"
)

func cmdMaths(e *interaction.Event) *interaction.EventCallback {
	expression, err := govaluate.NewEvaluableExpression(e.Data.Options[0].Value)
	if err != nil {
		return e.Error("Bad expression!")
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return e.Error("Cannot be evaluated!")
	}
	return e.Reply(fmt.Sprintf("**🧮 `%v`  =  `%v`**", expression.String(), result))
}
