package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsageProfile.
// Experimental.
type IUsageProfileRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UsageProfile resource.
	// Experimental.
	UsageProfileRef() *UsageProfileReference
}

// The jsii proxy for IUsageProfileRef
type jsiiProxy_IUsageProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUsageProfileRef) UsageProfileRef() *UsageProfileReference {
	var returns *UsageProfileReference
	_jsii_.Get(
		j,
		"usageProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsageProfileRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsageProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

