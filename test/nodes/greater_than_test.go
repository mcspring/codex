package nodes

import (
  "codex/tree/nodes"
  "testing"
)

func TestGreaterThan(t *testing.T) {
  left, right := 1, 2
  greaterThan := &nodes.GreaterThan{left, right}
  if left != greaterThan.Left {
    t.Errorf("Expect Left GreaterThan leaf to equal %v, got %v.", left, greaterThan.Left)
  } else if right != greaterThan.Right {
    t.Errorf("Expect Right GreaterThan leaf to equal %v, got %v.", right, greaterThan.Right)
  }
}

func TestGreaterThanOr(t *testing.T) {
  left, right := 1, 2
  greaterThan := &nodes.GreaterThan{left, right}
  other := 3
  or := greaterThan.Or(other)
  if greaterThan != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", greaterThan, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestGreaterThanAnd(t *testing.T) {
  left, right := 1, 2
  greaterThan := &nodes.GreaterThan{left, right}
  other := 3
  and := greaterThan.And(other)
  if greaterThan != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", greaterThan, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right)
  }
}