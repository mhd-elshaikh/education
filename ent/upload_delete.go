// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"education/ent/predicate"
	"education/ent/upload"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UploadDelete is the builder for deleting a Upload entity.
type UploadDelete struct {
	config
	hooks    []Hook
	mutation *UploadMutation
}

// Where appends a list predicates to the UploadDelete builder.
func (ud *UploadDelete) Where(ps ...predicate.Upload) *UploadDelete {
	ud.mutation.Where(ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UploadDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ud.hooks) == 0 {
		affected, err = ud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UploadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ud.mutation = mutation
			affected, err = ud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ud.hooks) - 1; i >= 0; i-- {
			if ud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UploadDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UploadDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: upload.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upload.FieldID,
			},
		},
	}
	if ps := ud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UploadDeleteOne is the builder for deleting a single Upload entity.
type UploadDeleteOne struct {
	ud *UploadDelete
}

// Exec executes the deletion query.
func (udo *UploadDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{upload.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UploadDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
