package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"reflect"
)

// Expression hold the object graph as defined in the rule semantic.
// an expression could hold a predicate or pair of logical operated expression.
type Expression struct {
	LeftExpression   *Expression
	RightExpression  *Expression
	LogicalOperator  LogicalOperator
	Predicate        *Predicate
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// Initialize this object graph with necessary context prior engine execution.
func (expr *Expression) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	expr.knowledgeContext = knowledgeContext
	expr.ruleCtx = ruleCtx
	expr.dataCtx = dataCtx

	if expr.LeftExpression != nil {
		expr.LeftExpression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if expr.RightExpression != nil {
		expr.RightExpression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if expr.Predicate != nil {
		expr.Predicate.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// AcceptExpression will store expression as they are defined in the rule script, into this object graph.
func (expr *Expression) AcceptExpression(expression *Expression) error {
	if expr.LeftExpression == nil {
		expr.LeftExpression = expression
	} else if expr.RightExpression == nil {
		expr.RightExpression = expression
	} else {
		return errors.Errorf("expression alredy set twice")
	}
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (expr *Expression) Evaluate() (reflect.Value, error) {
	if expr.Predicate != nil {
		return expr.Predicate.Evaluate()
	}
	lv, err := expr.LeftExpression.Evaluate()
	if err != nil {
		return lv, errors.Trace(err)
	}
	rv, err := expr.RightExpression.Evaluate()
	if err != nil {
		return rv, errors.Trace(err)
	}
	if rv.Kind() == lv.Kind() && rv.Kind() == reflect.Bool {
		if expr.LogicalOperator == LogicalOperatorAnd {
			return reflect.ValueOf(lv.Bool() && rv.Bool()), nil
		}
		return reflect.ValueOf(lv.Bool() || rv.Bool()), nil
	} else {
		return reflect.ValueOf(nil), errors.Errorf("cannot apply logical for non boolean expression")
	}
}
