package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/o-ga09/tutorial-ec-backend/pkg/strings"
)

type User struct {
	id string
	email string
	phoneNumber string
	lastName string
	firstName string
	address address
}

type address struct {
	pref string
	city string
	extra string
}

func Reconstract(
	id string,
	email string,
	phonenumber string,
	lastname string,
	firstname string,
	pref string,
	city string,
	extra string,
) (*User, error) {
	return newUser(email,phonenumber,lastname,firstname,pref,city,extra)
}

func NewUser(
	email string,
	phonenumber string,
	lastname string,
	firstname string,
	pref string,
	city string,
	extra string,
) (*User, error) {
	return newUser(email,phonenumber,lastname,firstname,pref,city,extra)
}

func newUser(
	email string,
	phonenumber string,
	lastname string,
	firstname string,
	pref string,
	city string,
	extra string,
) (*User, error) {
	address, err := newAddress(pref,city,extra)
	if err != nil {
		return nil, err
	}

	return &User{
		id: strings.RemoveHyphen(uuid.New().String()),
		email: email,
		phoneNumber: phonenumber,
		lastName: lastname,
		firstName: firstname,
		address: *address,
	}, nil
}

func newAddress(
	pref string,
	city string,
	extra string,
) (*address, error) {
	if city == "" || pref == "" || extra == "" {
		return nil, errors.New("住所の値が不正です")
	}
	
	return &address{
		city: city,
		pref: pref,
		extra: extra,
	}, nil
}

func(u *User) ID()string {return u.id}
func(u *User) Email()string {return u.email}
func(u *User) PhoneNumber()string {return u.phoneNumber}
func(u *User) LastName()string {return u.lastName}
func(u *User) FirstName()string {return u.firstName}
func(u *User) Address()address {return u.address}
func(u *User) Pref()string {return u.address.pref}
func(u *User) City()string {return u.address.city}
func(u *User) Extra()string {return u.address.extra}