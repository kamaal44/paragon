// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/credential"
)

// CredentialCreate is the builder for creating a Credential entity.
type CredentialCreate struct {
	config
	principal *string
	secret    *string
	fails     *int
}

// SetPrincipal sets the principal field.
func (cc *CredentialCreate) SetPrincipal(s string) *CredentialCreate {
	cc.principal = &s
	return cc
}

// SetSecret sets the secret field.
func (cc *CredentialCreate) SetSecret(s string) *CredentialCreate {
	cc.secret = &s
	return cc
}

// SetFails sets the fails field.
func (cc *CredentialCreate) SetFails(i int) *CredentialCreate {
	cc.fails = &i
	return cc
}

// SetNillableFails sets the fails field if the given value is not nil.
func (cc *CredentialCreate) SetNillableFails(i *int) *CredentialCreate {
	if i != nil {
		cc.SetFails(*i)
	}
	return cc
}

// Save creates the Credential in the database.
func (cc *CredentialCreate) Save(ctx context.Context) (*Credential, error) {
	if cc.principal == nil {
		return nil, errors.New("ent: missing required field \"principal\"")
	}
	if cc.secret == nil {
		return nil, errors.New("ent: missing required field \"secret\"")
	}
	if cc.fails == nil {
		v := credential.DefaultFails
		cc.fails = &v
	}
	if err := credential.FailsValidator(*cc.fails); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"fails\": %v", err)
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CredentialCreate) SaveX(ctx context.Context) *Credential {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CredentialCreate) sqlSave(ctx context.Context) (*Credential, error) {
	var (
		res sql.Result
		c   = &Credential{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Dialect(cc.driver.Dialect()).
		Insert(credential.Table).
		Default()
	if value := cc.principal; value != nil {
		builder.Set(credential.FieldPrincipal, *value)
		c.Principal = *value
	}
	if value := cc.secret; value != nil {
		builder.Set(credential.FieldSecret, *value)
		c.Secret = *value
	}
	if value := cc.fails; value != nil {
		builder.Set(credential.FieldFails, *value)
		c.Fails = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
