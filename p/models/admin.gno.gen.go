// Code generated by github.com/gnolang/gno. DO NOT EDIT.

//go:build gno
// +build gno

// Package models contains the Admin and IAdmin types for managing
// admin and controller access in a Golang blockchain smart contract.
package models

import (
	"errors"
	"github.com/gnolang/gno/stdlibs/stdshim" // gnolang stdlib
)

// IAdmin defines the interface for the Admin contract, including methods
// for setting various properties and managing admins and controllers.
type IAdmin interface {
	SetInPause(inPause bool) string
	SetMessage(message string) string
	SetTransferLimit(amount int64) string
	SetAdminAddr(addr std.Address) string
	ControllerAdd(addr std.Address) string
	ControllerRemove(addr std.Address) string
	AdminAdd(addr std.Address) string
	AdminRemove(addr std.Address) string
}

// Admin is a struct that stores the admins, controllers, and other properties
// for a Golang blockchain smart contract.
type Admin struct {
	admins      []std.Address
	controllers []std.Address
	gInPause    bool
	gMessage    string
	gLimit      std.Coin
}

// NewAdmin creates a new Admin contract with default values.
// It returns a pointer to the new Admin instance.
func NewAdmin() *Admin {
	return &Admin{
		admins:      []std.Address{},
		controllers: []std.Address{},
		gInPause:    false,
		gMessage:    "",
		gLimit:      std.Coin{Denom: "ugnot", Amount: 0},
	}
}

// SetInPause sets the contract pause state. It returns an empty string
// if successful or an error message if the caller is not an admin or controller.
func (a *Admin) SetInPause(inPause bool) string {
	if err := a.assertIsAdminOrController(); err != nil {
		return err.Error()
	}
	a.gInPause = inPause
	return ""
}

// SetMessage sets the contract message. It returns an empty string
// if successful or an error message if the caller is not an admin or controller.
func (a *Admin) SetMessage(message string) string {
	if err := a.assertIsAdminOrController(); err != nil {
		return err.Error()
	}
	a.gMessage = message
	return ""
}

// SetTransferLimit sets the contract transfer limit. It returns an empty string
// if successful or an error message if the caller is not an admin or controller.
func (a *Admin) SetTransferLimit(amount int64) string {
	if err := a.assertIsAdminOrController(); err != nil {
		return err.Error()
	}
	a.gLimit = std.Coin{Denom: "ugnot", Amount: amount}
	return ""
}

// assertIsAdminOrController checks if the caller is an admin or controller.
// It returns nil if the caller is an admin or controller, or an error otherwise.
func (a *Admin) assertIsAdminOrController() error {
	if err := a.assertIsAdmin(); err == nil {
		return nil
	}
	if err := a.assertIsController(); err == nil {
		return nil
	}
	return errors.New("restricted for admin or controller")
}

// assertIsAdmin checks if the caller is an admin.
// It returns nil if the caller is an admin, or an error otherwise.
func (a *Admin) assertIsAdmin() error {
	caller := std.GetOrigCaller()

	for _, admin := range a.admins {
		if caller.Equals(admin) {
			return nil
		}
	}

	return errors.New("restricted for admin")
}

// assertIsController checks if the caller is a controller.
// It returns nil if the caller is a controller, or an error otherwise.
func (a *Admin) assertIsController() error {
	caller := std.GetOrigCaller()

	for _, controller := range a.controllers {
		if caller.Equals(controller) {
			return nil
		}
	}

	return errors.New("restricted for controller")
}
