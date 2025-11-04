package awsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TargetAccountConfiguration.
// Experimental.
type ITargetAccountConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TargetAccountConfiguration resource.
	// Experimental.
	TargetAccountConfigurationRef() *TargetAccountConfigurationReference
}

// The jsii proxy for ITargetAccountConfigurationRef
type jsiiProxy_ITargetAccountConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITargetAccountConfigurationRef) TargetAccountConfigurationRef() *TargetAccountConfigurationReference {
	var returns *TargetAccountConfigurationReference
	_jsii_.Get(
		j,
		"targetAccountConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetAccountConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetAccountConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

