package interfaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Used to indicate that a particular construct has an resource environment.
type IEnvironmentAware interface {
	// The environment this resource belongs to.
	//
	// For resources that are created and managed in a Stack (those created by
	// creating new class instances like `new Role()`, `new Bucket()`, etc.), this
	// is always the same as the environment of the stack they belong to.
	//
	// For referenced resources (those obtained from referencing methods like
	// `Role.fromRoleArn()`, `Bucket.fromBucketName()`, etc.), they might be
	// different than the stack they were imported into.
	Env() *ResourceEnvironment
}

// The jsii proxy for IEnvironmentAware
type jsiiProxy_IEnvironmentAware struct {
	_ byte // padding
}

func (j *jsiiProxy_IEnvironmentAware) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

