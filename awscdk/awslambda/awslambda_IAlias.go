package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IAlias interface {
	IFunction
	// Name of this alias.
	// Experimental.
	AliasName() *string
	// The underlying Lambda function version.
	// Experimental.
	Version() IVersion
}

// The jsii proxy for IAlias
type jsiiProxy_IAlias struct {
	jsiiProxy_IFunction
}

func (j *jsiiProxy_IAlias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Version() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

