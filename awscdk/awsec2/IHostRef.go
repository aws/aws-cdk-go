package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Host.
// Experimental.
type IHostRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Host resource.
	// Experimental.
	HostRef() *HostReference
}

// The jsii proxy for IHostRef
type jsiiProxy_IHostRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IHostRef) HostRef() *HostReference {
	var returns *HostReference
	_jsii_.Get(
		j,
		"hostRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

