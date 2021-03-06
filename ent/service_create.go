// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/event"
	"github.com/kcarretto/paragon/ent/service"
	"github.com/kcarretto/paragon/ent/tag"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	Name        *string
	PubKey      *string
	Config      *string
	IsActivated *bool
	tag         map[int]struct{}
	events      map[int]struct{}
}

// SetName sets the Name field.
func (sc *ServiceCreate) SetName(s string) *ServiceCreate {
	sc.Name = &s
	return sc
}

// SetPubKey sets the PubKey field.
func (sc *ServiceCreate) SetPubKey(s string) *ServiceCreate {
	sc.PubKey = &s
	return sc
}

// SetConfig sets the Config field.
func (sc *ServiceCreate) SetConfig(s string) *ServiceCreate {
	sc.Config = &s
	return sc
}

// SetNillableConfig sets the Config field if the given value is not nil.
func (sc *ServiceCreate) SetNillableConfig(s *string) *ServiceCreate {
	if s != nil {
		sc.SetConfig(*s)
	}
	return sc
}

// SetIsActivated sets the IsActivated field.
func (sc *ServiceCreate) SetIsActivated(b bool) *ServiceCreate {
	sc.IsActivated = &b
	return sc
}

// SetNillableIsActivated sets the IsActivated field if the given value is not nil.
func (sc *ServiceCreate) SetNillableIsActivated(b *bool) *ServiceCreate {
	if b != nil {
		sc.SetIsActivated(*b)
	}
	return sc
}

// SetTagID sets the tag edge to Tag by id.
func (sc *ServiceCreate) SetTagID(id int) *ServiceCreate {
	if sc.tag == nil {
		sc.tag = make(map[int]struct{})
	}
	sc.tag[id] = struct{}{}
	return sc
}

// SetTag sets the tag edge to Tag.
func (sc *ServiceCreate) SetTag(t *Tag) *ServiceCreate {
	return sc.SetTagID(t.ID)
}

// AddEventIDs adds the events edge to Event by ids.
func (sc *ServiceCreate) AddEventIDs(ids ...int) *ServiceCreate {
	if sc.events == nil {
		sc.events = make(map[int]struct{})
	}
	for i := range ids {
		sc.events[ids[i]] = struct{}{}
	}
	return sc
}

// AddEvents adds the events edges to Event.
func (sc *ServiceCreate) AddEvents(e ...*Event) *ServiceCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return sc.AddEventIDs(ids...)
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	if sc.Name == nil {
		return nil, errors.New("ent: missing required field \"Name\"")
	}
	if err := service.NameValidator(*sc.Name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
	}
	if sc.PubKey == nil {
		return nil, errors.New("ent: missing required field \"PubKey\"")
	}
	if err := service.PubKeyValidator(*sc.PubKey); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"PubKey\": %v", err)
	}
	if sc.Config == nil {
		v := service.DefaultConfig
		sc.Config = &v
	}
	if sc.IsActivated == nil {
		v := service.DefaultIsActivated
		sc.IsActivated = &v
	}
	if len(sc.tag) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"tag\"")
	}
	if sc.tag == nil {
		return nil, errors.New("ent: missing required edge \"tag\"")
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	var (
		s     = &Service{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: service.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: service.FieldID,
			},
		}
	)
	if value := sc.Name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldName,
		})
		s.Name = *value
	}
	if value := sc.PubKey; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldPubKey,
		})
		s.PubKey = *value
	}
	if value := sc.Config; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldConfig,
		})
		s.Config = *value
	}
	if value := sc.IsActivated; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: service.FieldIsActivated,
		})
		s.IsActivated = *value
	}
	if nodes := sc.tag; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.TagTable,
			Columns: []string{service.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.events; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.EventsTable,
			Columns: []string{service.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
