package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnycastIpList.
// Experimental.
type IAnycastIpListRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AnycastIpList resource.
	// Experimental.
	AnycastIpListRef() *AnycastIpListReference
}

// The jsii proxy for IAnycastIpListRef
type jsiiProxy_IAnycastIpListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAnycastIpListRef) AnycastIpListRef() *AnycastIpListReference {
	var returns *AnycastIpListReference
	_jsii_.Get(
		j,
		"anycastIpListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnycastIpListRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnycastIpListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

