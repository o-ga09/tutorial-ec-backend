package owner

import (
	"github.com/google/uuid"
	"github.com/o-ga09/tutorial-ec-backend/pkg/strings"
)

type Owner struct {
	id string
	name string
	email string
}

func Recontract(
	id string,
	name string,
	email string,
) (*Owner, error) {
	return newOwner(id,name,email)
}

func NewOwner(
	name string,
	email string,
) (*Owner, error) {
	return newOwner(strings.RemoveHyphen(uuid.New().String()),name,email)
}

func newOwner(
	id string,
	name string,
	email string,
) (*Owner, error) {
	return &Owner{
		id: id,
		email: email,
		name: name,
	}, nil
}

func(o *Owner) ID() string {return o.id}
func(o *Owner) Email() string {return o.email}
func(o *Owner) Name() string {return o.name}