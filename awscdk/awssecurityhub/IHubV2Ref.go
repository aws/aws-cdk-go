package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HubV2.
// Experimental.
type IHubV2Ref interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a HubV2 resource.
	// Experimental.
	HubV2Ref() *HubV2Reference
}

// The jsii proxy for IHubV2Ref
type jsiiProxy_IHubV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IHubV2Ref) HubV2Ref() *HubV2Reference {
	var returns *HubV2Reference
	_jsii_.Get(
		j,
		"hubV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHubV2Ref) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHubV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

