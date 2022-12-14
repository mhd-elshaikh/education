// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"education/ent/playlist"
	"education/ent/predicate"
	"education/ent/upload"
	"education/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UploadUpdate is the builder for updating Upload entities.
type UploadUpdate struct {
	config
	hooks    []Hook
	mutation *UploadMutation
}

// Where appends a list predicates to the UploadUpdate builder.
func (uu *UploadUpdate) Where(ps ...predicate.Upload) *UploadUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UploadUpdate) SetUserID(i int) *UploadUpdate {
	uu.mutation.SetUserID(i)
	return uu
}

// SetPlaylistID sets the "playlist_id" field.
func (uu *UploadUpdate) SetPlaylistID(i int) *UploadUpdate {
	uu.mutation.SetPlaylistID(i)
	return uu
}

// SetName sets the "name" field.
func (uu *UploadUpdate) SetName(s string) *UploadUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetUUID sets the "uuid" field.
func (uu *UploadUpdate) SetUUID(s string) *UploadUpdate {
	uu.mutation.SetUUID(s)
	return uu
}

// SetMimeType sets the "mime_type" field.
func (uu *UploadUpdate) SetMimeType(s string) *UploadUpdate {
	uu.mutation.SetMimeType(s)
	return uu
}

// SetSize sets the "size" field.
func (uu *UploadUpdate) SetSize(i int) *UploadUpdate {
	uu.mutation.ResetSize()
	uu.mutation.SetSize(i)
	return uu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (uu *UploadUpdate) SetNillableSize(i *int) *UploadUpdate {
	if i != nil {
		uu.SetSize(*i)
	}
	return uu
}

// AddSize adds i to the "size" field.
func (uu *UploadUpdate) AddSize(i int) *UploadUpdate {
	uu.mutation.AddSize(i)
	return uu
}

// ClearSize clears the value of the "size" field.
func (uu *UploadUpdate) ClearSize() *UploadUpdate {
	uu.mutation.ClearSize()
	return uu
}

// SetTitle sets the "title" field.
func (uu *UploadUpdate) SetTitle(s string) *UploadUpdate {
	uu.mutation.SetTitle(s)
	return uu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uu *UploadUpdate) SetNillableTitle(s *string) *UploadUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// ClearTitle clears the value of the "title" field.
func (uu *UploadUpdate) ClearTitle() *UploadUpdate {
	uu.mutation.ClearTitle()
	return uu
}

// SetDescription sets the "description" field.
func (uu *UploadUpdate) SetDescription(s string) *UploadUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UploadUpdate) SetNillableDescription(s *string) *UploadUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// ClearDescription clears the value of the "description" field.
func (uu *UploadUpdate) ClearDescription() *UploadUpdate {
	uu.mutation.ClearDescription()
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UploadUpdate) SetUpdatedAt(t time.Time) *UploadUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetUser sets the "user" edge to the User entity.
func (uu *UploadUpdate) SetUser(u *User) *UploadUpdate {
	return uu.SetUserID(u.ID)
}

// SetPlaylist sets the "playlist" edge to the Playlist entity.
func (uu *UploadUpdate) SetPlaylist(p *Playlist) *UploadUpdate {
	return uu.SetPlaylistID(p.ID)
}

// Mutation returns the UploadMutation object of the builder.
func (uu *UploadUpdate) Mutation() *UploadMutation {
	return uu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uu *UploadUpdate) ClearUser() *UploadUpdate {
	uu.mutation.ClearUser()
	return uu
}

// ClearPlaylist clears the "playlist" edge to the Playlist entity.
func (uu *UploadUpdate) ClearPlaylist() *UploadUpdate {
	uu.mutation.ClearPlaylist()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UploadUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UploadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UploadUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UploadUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UploadUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UploadUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := upload.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UploadUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := upload.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Upload.name": %w`, err)}
		}
	}
	if _, ok := uu.mutation.UserID(); uu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upload.user"`)
	}
	if _, ok := uu.mutation.PlaylistID(); uu.mutation.PlaylistCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upload.playlist"`)
	}
	return nil
}

func (uu *UploadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upload.Table,
			Columns: upload.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upload.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldName,
		})
	}
	if value, ok := uu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldUUID,
		})
	}
	if value, ok := uu.mutation.MimeType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldMimeType,
		})
	}
	if value, ok := uu.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: upload.FieldSize,
		})
	}
	if value, ok := uu.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: upload.FieldSize,
		})
	}
	if uu.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: upload.FieldSize,
		})
	}
	if value, ok := uu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldTitle,
		})
	}
	if uu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: upload.FieldTitle,
		})
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldDescription,
		})
	}
	if uu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: upload.FieldDescription,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: upload.FieldUpdatedAt,
		})
	}
	if uu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.UserTable,
			Columns: []string{upload.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.UserTable,
			Columns: []string{upload.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PlaylistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.PlaylistTable,
			Columns: []string{upload.PlaylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PlaylistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.PlaylistTable,
			Columns: []string{upload.PlaylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upload.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UploadUpdateOne is the builder for updating a single Upload entity.
type UploadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UploadMutation
}

// SetUserID sets the "user_id" field.
func (uuo *UploadUpdateOne) SetUserID(i int) *UploadUpdateOne {
	uuo.mutation.SetUserID(i)
	return uuo
}

// SetPlaylistID sets the "playlist_id" field.
func (uuo *UploadUpdateOne) SetPlaylistID(i int) *UploadUpdateOne {
	uuo.mutation.SetPlaylistID(i)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UploadUpdateOne) SetName(s string) *UploadUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetUUID sets the "uuid" field.
func (uuo *UploadUpdateOne) SetUUID(s string) *UploadUpdateOne {
	uuo.mutation.SetUUID(s)
	return uuo
}

// SetMimeType sets the "mime_type" field.
func (uuo *UploadUpdateOne) SetMimeType(s string) *UploadUpdateOne {
	uuo.mutation.SetMimeType(s)
	return uuo
}

// SetSize sets the "size" field.
func (uuo *UploadUpdateOne) SetSize(i int) *UploadUpdateOne {
	uuo.mutation.ResetSize()
	uuo.mutation.SetSize(i)
	return uuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (uuo *UploadUpdateOne) SetNillableSize(i *int) *UploadUpdateOne {
	if i != nil {
		uuo.SetSize(*i)
	}
	return uuo
}

// AddSize adds i to the "size" field.
func (uuo *UploadUpdateOne) AddSize(i int) *UploadUpdateOne {
	uuo.mutation.AddSize(i)
	return uuo
}

// ClearSize clears the value of the "size" field.
func (uuo *UploadUpdateOne) ClearSize() *UploadUpdateOne {
	uuo.mutation.ClearSize()
	return uuo
}

// SetTitle sets the "title" field.
func (uuo *UploadUpdateOne) SetTitle(s string) *UploadUpdateOne {
	uuo.mutation.SetTitle(s)
	return uuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uuo *UploadUpdateOne) SetNillableTitle(s *string) *UploadUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// ClearTitle clears the value of the "title" field.
func (uuo *UploadUpdateOne) ClearTitle() *UploadUpdateOne {
	uuo.mutation.ClearTitle()
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UploadUpdateOne) SetDescription(s string) *UploadUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UploadUpdateOne) SetNillableDescription(s *string) *UploadUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// ClearDescription clears the value of the "description" field.
func (uuo *UploadUpdateOne) ClearDescription() *UploadUpdateOne {
	uuo.mutation.ClearDescription()
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UploadUpdateOne) SetUpdatedAt(t time.Time) *UploadUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetUser sets the "user" edge to the User entity.
func (uuo *UploadUpdateOne) SetUser(u *User) *UploadUpdateOne {
	return uuo.SetUserID(u.ID)
}

// SetPlaylist sets the "playlist" edge to the Playlist entity.
func (uuo *UploadUpdateOne) SetPlaylist(p *Playlist) *UploadUpdateOne {
	return uuo.SetPlaylistID(p.ID)
}

// Mutation returns the UploadMutation object of the builder.
func (uuo *UploadUpdateOne) Mutation() *UploadMutation {
	return uuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uuo *UploadUpdateOne) ClearUser() *UploadUpdateOne {
	uuo.mutation.ClearUser()
	return uuo
}

// ClearPlaylist clears the "playlist" edge to the Playlist entity.
func (uuo *UploadUpdateOne) ClearPlaylist() *UploadUpdateOne {
	uuo.mutation.ClearPlaylist()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UploadUpdateOne) Select(field string, fields ...string) *UploadUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Upload entity.
func (uuo *UploadUpdateOne) Save(ctx context.Context) (*Upload, error) {
	var (
		err  error
		node *Upload
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UploadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Upload)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UploadMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UploadUpdateOne) SaveX(ctx context.Context) *Upload {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UploadUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UploadUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UploadUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := upload.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UploadUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := upload.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Upload.name": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.UserID(); uuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upload.user"`)
	}
	if _, ok := uuo.mutation.PlaylistID(); uuo.mutation.PlaylistCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upload.playlist"`)
	}
	return nil
}

func (uuo *UploadUpdateOne) sqlSave(ctx context.Context) (_node *Upload, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upload.Table,
			Columns: upload.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upload.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Upload.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upload.FieldID)
		for _, f := range fields {
			if !upload.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != upload.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldName,
		})
	}
	if value, ok := uuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldUUID,
		})
	}
	if value, ok := uuo.mutation.MimeType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldMimeType,
		})
	}
	if value, ok := uuo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: upload.FieldSize,
		})
	}
	if value, ok := uuo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: upload.FieldSize,
		})
	}
	if uuo.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: upload.FieldSize,
		})
	}
	if value, ok := uuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldTitle,
		})
	}
	if uuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: upload.FieldTitle,
		})
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: upload.FieldDescription,
		})
	}
	if uuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: upload.FieldDescription,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: upload.FieldUpdatedAt,
		})
	}
	if uuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.UserTable,
			Columns: []string{upload.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.UserTable,
			Columns: []string{upload.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PlaylistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.PlaylistTable,
			Columns: []string{upload.PlaylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PlaylistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upload.PlaylistTable,
			Columns: []string{upload.PlaylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Upload{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upload.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
