package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentUtils := awscdk.Cx_api.NewEnvironmentUtils()
//
// Experimental.
type EnvironmentUtils interface {
}

// The jsii proxy struct for EnvironmentUtils
type jsiiProxy_EnvironmentUtils struct {
	_ byte // padding
}

// Experimental.
func NewEnvironmentUtils() EnvironmentUtils {
	_init_.Initialize()

	j := jsiiProxy_EnvironmentUtils{}

	_jsii_.Create(
		"monocdk.cx_api.EnvironmentUtils",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEnvironmentUtils_Override(e EnvironmentUtils) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.cx_api.EnvironmentUtils",
		nil, // no parameters
		e,
	)
}

// Format an environment string from an account and region.
// Experimental.
func EnvironmentUtils_Format(account *string, region *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.cx_api.EnvironmentUtils",
		"format",
		[]interface{}{account, region},
		&returns,
	)

	return returns
}

// Build an environment object from an account and region.
// Experimental.
func EnvironmentUtils_Make(account *string, region *string) *Environment {
	_init_.Initialize()

	var returns *Environment

	_jsii_.StaticInvoke(
		"monocdk.cx_api.EnvironmentUtils",
		"make",
		[]interface{}{account, region},
		&returns,
	)

	return returns
}

// Experimental.
func EnvironmentUtils_Parse(environment *string) *Environment {
	_init_.Initialize()

	var returns *Environment

	_jsii_.StaticInvoke(
		"monocdk.cx_api.EnvironmentUtils",
		"parse",
		[]interface{}{environment},
		&returns,
	)

	return returns
}

