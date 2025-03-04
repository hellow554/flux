package join

import (
	"context"

	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/stdlib/universe"
)

const SortMergeJoinKind = "sortmergejoin"

func init() {
	plan.RegisterPhysicalRules(SortMergeJoinPredicateRule{})
	execute.RegisterTransformation(SortMergeJoinKind, createJoinTransformation)
}

type SortMergeJoinProcedureSpec EquiJoinProcedureSpec

func (p *SortMergeJoinProcedureSpec) Kind() plan.ProcedureKind {
	return plan.ProcedureKind(SortMergeJoinKind)
}

func (p *SortMergeJoinProcedureSpec) Copy() plan.ProcedureSpec {
	return &SortMergeJoinProcedureSpec{
		On:     p.On,
		As:     p.As,
		Left:   p.Left,
		Right:  p.Right,
		Method: p.Method,
	}
}

// RequiredAttributes says that merge join must have its left input
// sorted by the left side join keys, and the right input must be sorted
// by the right side join keys.
func (p *SortMergeJoinProcedureSpec) RequiredAttributes() []plan.PhysicalAttributes {
	return []plan.PhysicalAttributes{
		{
			plan.CollationKey: &plan.CollationAttr{
				Columns: getJoinKeyCols(p.On, true),
			},
		},
		{
			plan.CollationKey: &plan.CollationAttr{
				Columns: getJoinKeyCols(p.On, false),
			},
		},
	}
}

func (p *SortMergeJoinProcedureSpec) Cost(inStats []plan.Statistics) (cost plan.Cost, outStats plan.Statistics) {
	return plan.Cost{}, plan.Statistics{}
}

type SortMergeJoinPredicateRule struct{}

func (SortMergeJoinPredicateRule) Name() string {
	return "sortMergeJoinPredicate"
}

func (SortMergeJoinPredicateRule) Pattern() plan.Pattern {
	return plan.Pat(EquiJoinKind, plan.Any(), plan.Any())
}

func (SortMergeJoinPredicateRule) Rewrite(ctx context.Context, n plan.Node) (plan.Node, bool, error) {
	s := n.ProcedureSpec()
	spec, ok := s.(*EquiJoinProcedureSpec)
	if !ok {
		return nil, false, errors.New(codes.Internal, "invalid spec type on join node")
	}

	predecessors := n.Predecessors()
	n.ClearPredecessors()

	makeSortNode := func(parentNode plan.Node, columns []string) *plan.PhysicalPlanNode {
		sortProc := universe.SortProcedureSpec{
			Columns: columns,
		}
		sortNode := plan.CreateUniquePhysicalNode(ctx, "sortMergeJoin", &sortProc)

		sortNode.AddPredecessors(parentNode)
		sortNode.AddSuccessors(n)
		n.AddPredecessors(sortNode)

		return sortNode
	}

	successors := predecessors[0].Successors()

	columns := make([]string, 0, len(spec.On))
	for _, pair := range spec.On {
		columns = append(columns, pair.Left)
	}
	successors[0] = makeSortNode(predecessors[0], columns)

	successors = predecessors[1].Successors()

	columns = make([]string, 0, len(spec.On))
	for _, pair := range spec.On {
		columns = append(columns, pair.Right)
	}
	successors[0] = makeSortNode(predecessors[1], columns)

	// Replace the spec so we don't end up trying to apply this rewrite forever
	x := SortMergeJoinProcedureSpec(*spec)
	if err := n.ReplaceSpec(&x); err != nil {
		return n, false, err
	}

	return n, true, nil
}
