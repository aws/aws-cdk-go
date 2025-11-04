package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Hub.
// Experimental.
type IHubRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Hub resource.
	// Experimental.
	HubRef() *HubReference
}

// The jsii proxy for IHubRef
type jsiiProxy_IHubRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IHubRef) HubRef() *HubReference {
	var returns *HubReference
	_jsii_.Get(
		j,
		"hubRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHubRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHubRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

