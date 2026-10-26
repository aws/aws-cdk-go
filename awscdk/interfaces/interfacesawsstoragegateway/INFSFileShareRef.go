package interfacesawsstoragegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstoragegateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NFSFileShare.
// Experimental.
type INFSFileShareRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NFSFileShare resource.
	// Experimental.
	NfsFileShareRef() *NFSFileShareReference
}

// The jsii proxy for INFSFileShareRef
type jsiiProxy_INFSFileShareRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_INFSFileShareRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INFSFileShareRef) NfsFileShareRef() *NFSFileShareReference {
	var returns *NFSFileShareReference
	_jsii_.Get(
		j,
		"nfsFileShareRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INFSFileShareRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INFSFileShareRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

