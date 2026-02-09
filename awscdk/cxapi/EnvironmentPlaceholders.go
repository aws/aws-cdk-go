package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Placeholders which can be used manifests.
//
// These can occur both in the Asset Manifest as well as the general
// Cloud Assembly manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentPlaceholders := awscdk.Cx_api.NewEnvironmentPlaceholders()
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type EnvironmentPlaceholders interface {
}

// The jsii proxy struct for EnvironmentPlaceholders
type jsiiProxy_EnvironmentPlaceholders struct {
	_ byte // padding
}

// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
func NewEnvironmentPlaceholders() EnvironmentPlaceholders {
	_init_.Initialize()

	j := jsiiProxy_EnvironmentPlaceholders{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		nil, // no parameters
		&j,
	)

	return &j
}

// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
func NewEnvironmentPlaceholders_Override(e EnvironmentPlaceholders) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		nil, // no parameters
		e,
	)
}

// Replace the environment placeholders in all strings found in a complex object.
//
// Duplicated between cdk-assets and aws-cdk CLI because we don't have a good single place to put it
// (they're nominally independent tools).
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
func EnvironmentPlaceholders_Replace(object interface{}, values *EnvironmentPlaceholderValues) interface{} {
	_init_.Initialize()

	if err := validateEnvironmentPlaceholders_ReplaceParameters(object, values); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"replace",
		[]interface{}{object, values},
		&returns,
	)

	return returns
}

// Like 'replace', but asynchronous.
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
func EnvironmentPlaceholders_ReplaceAsync(object interface{}, provider IEnvironmentPlaceholderProvider) interface{} {
	_init_.Initialize()

	if err := validateEnvironmentPlaceholders_ReplaceAsyncParameters(object, provider); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"replaceAsync",
		[]interface{}{object, provider},
		&returns,
	)

	return returns
}

func EnvironmentPlaceholders_CURRENT_ACCOUNT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"CURRENT_ACCOUNT",
		&returns,
	)
	return returns
}

func EnvironmentPlaceholders_CURRENT_PARTITION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"CURRENT_PARTITION",
		&returns,
	)
	return returns
}

func EnvironmentPlaceholders_CURRENT_REGION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"CURRENT_REGION",
		&returns,
	)
	return returns
}

