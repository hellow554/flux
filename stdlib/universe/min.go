package universe

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/array"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/runtime"
)

const MinKind = "min"

type MinOpSpec struct {
	execute.SelectorConfig
}

func init() {
	minSignature := runtime.MustLookupBuiltinType("universe", "min")

	runtime.RegisterPackageValue("universe", MinKind, flux.MustValue(flux.FunctionValue(MinKind, CreateMinOpSpec, minSignature)))
	flux.RegisterOpSpec(MinKind, newMinOp)
	plan.RegisterProcedureSpec(MinKind, newMinProcedure, MinKind)
	execute.RegisterTransformation(MinKind, createMinTransformation)
}

func CreateMinOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	if err := a.AddParentFromArgs(args); err != nil {
		return nil, err
	}

	spec := new(MinOpSpec)
	if err := spec.SelectorConfig.ReadArgs(args); err != nil {
		return nil, err
	}

	return spec, nil
}

func newMinOp() flux.OperationSpec {
	return new(MinOpSpec)
}

func (s *MinOpSpec) Kind() flux.OperationKind {
	return MinKind
}

type MinProcedureSpec struct {
	execute.SelectorConfig
}

func newMinProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	spec, ok := qs.(*MinOpSpec)
	if !ok {
		return nil, errors.Newf(codes.Internal, "invalid spec type %T", qs)
	}
	return &MinProcedureSpec{
		SelectorConfig: spec.SelectorConfig,
	}, nil
}

func (s *MinProcedureSpec) Kind() plan.ProcedureKind {
	return MinKind
}
func (s *MinProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(MinProcedureSpec)
	ns.SelectorConfig = s.SelectorConfig
	return ns
}

// TriggerSpec implements plan.TriggerAwareProcedureSpec
func (s *MinProcedureSpec) TriggerSpec() plan.TriggerSpec {
	return plan.NarrowTransformationTriggerSpec{}
}

type MinSelector struct {
	set  bool
	rows []execute.Row
}

func createMinTransformation(id execute.DatasetID, mode execute.AccumulationMode, spec plan.ProcedureSpec, a execute.Administration) (execute.Transformation, execute.Dataset, error) {
	ps, ok := spec.(*MinProcedureSpec)
	if !ok {
		return nil, nil, errors.Newf(codes.Internal, "invalid spec type %T", ps)
	}
	t, d := execute.NewRowSelectorTransformationAndDataset(id, mode, new(MinSelector), ps.SelectorConfig, a.Allocator())
	return t, d, nil
}

type MinIntSelector struct {
	MinSelector
	min int64
}
type MinUIntSelector struct {
	MinSelector
	min uint64
}
type MinFloatSelector struct {
	MinSelector
	min float64
}
type MinTimeSelector struct {
	MinIntSelector
}

func (s *MinSelector) NewTimeSelector() execute.DoTimeRowSelector {
	return new(MinTimeSelector)
}

func (s *MinSelector) NewBoolSelector() execute.DoBoolRowSelector {
	return nil
}

func (s *MinSelector) NewIntSelector() execute.DoIntRowSelector {
	return new(MinIntSelector)
}

func (s *MinSelector) NewUIntSelector() execute.DoUIntRowSelector {
	return new(MinUIntSelector)
}

func (s *MinSelector) NewFloatSelector() execute.DoFloatRowSelector {
	return new(MinFloatSelector)
}

func (s *MinSelector) NewStringSelector() execute.DoStringRowSelector {
	return nil
}

func (s *MinSelector) Rows() []execute.Row {
	if !s.set {
		return nil
	}
	return s.rows
}

func (s *MinSelector) selectRow(idx int, cr flux.ColReader) {
	// Capture row
	if idx >= 0 {
		s.rows = []execute.Row{execute.ReadRow(idx, cr)}
	}
}

func (s *MinTimeSelector) DoTime(vs *array.Int, cr flux.ColReader) {
	s.MinIntSelector.DoInt(vs, cr)
}
func (s *MinIntSelector) DoInt(vs *array.Int, cr flux.ColReader) {
	minIdx := -1
	for i := 0; i < vs.Len(); i++ {
		if vs.IsValid(i) {
			if v := vs.Value(i); !s.set || v < s.min {
				s.set = true
				s.min = v
				minIdx = i
			}
		}
	}
	s.selectRow(minIdx, cr)
}
func (s *MinUIntSelector) DoUInt(vs *array.Uint, cr flux.ColReader) {
	minIdx := -1
	for i := 0; i < vs.Len(); i++ {
		if vs.IsValid(i) {
			if v := vs.Value(i); !s.set || v < s.min {
				s.set = true
				s.min = v
				minIdx = i
			}
		}
	}
	s.selectRow(minIdx, cr)
}
func (s *MinFloatSelector) DoFloat(vs *array.Float, cr flux.ColReader) {
	minIdx := -1
	for i := 0; i < vs.Len(); i++ {
		if vs.IsValid(i) {
			if v := vs.Value(i); !s.set || v < s.min {
				s.set = true
				s.min = v
				minIdx = i
			}
		}
	}
	s.selectRow(minIdx, cr)
}
