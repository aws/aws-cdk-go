package interfacesawsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FileCache.
// Experimental.
type IFileCacheRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FileCache resource.
	// Experimental.
	FileCacheRef() *FileCacheReference
}

// The jsii proxy for IFileCacheRef
type jsiiProxy_IFileCacheRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFileCacheRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IFileCacheRef) FileCacheRef() *FileCacheReference {
	var returns *FileCacheReference
	_jsii_.Get(
		j,
		"fileCacheRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileCacheRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileCacheRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

