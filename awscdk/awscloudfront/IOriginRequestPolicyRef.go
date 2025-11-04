package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginRequestPolicy.
// Experimental.
type IOriginRequestPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OriginRequestPolicy resource.
	// Experimental.
	OriginRequestPolicyRef() *OriginRequestPolicyReference
}

// The jsii proxy for IOriginRequestPolicyRef
type jsiiProxy_IOriginRequestPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOriginRequestPolicyRef) OriginRequestPolicyRef() *OriginRequestPolicyReference {
	var returns *OriginRequestPolicyReference
	_jsii_.Get(
		j,
		"originRequestPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginRequestPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginRequestPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

