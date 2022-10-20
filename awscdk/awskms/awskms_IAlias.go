package awskms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A KMS Key alias.
//
// An alias can be used in all places that expect a key.
type IAlias interface {
	IKey
	// The name of the alias.
	AliasName() *string
	// The Key to which the Alias refers.
	AliasTargetKey() IKey
}

// The jsii proxy for IAlias
type jsiiProxy_IAlias struct {
	jsiiProxy_IKey
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

func (j *jsiiProxy_IAlias) AliasTargetKey() IKey {
	var returns IKey
	_jsii_.Get(
		j,
		"aliasTargetKey",
		&returns,
	)
	return returns
}

