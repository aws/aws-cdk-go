package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InfrastructureConfiguration.
// Experimental.
type IInfrastructureConfigurationRef interface {
	constructs.IConstruct
	// A reference to a InfrastructureConfiguration resource.
	// Experimental.
	InfrastructureConfigurationRef() *InfrastructureConfigurationReference
}

// The jsii proxy for IInfrastructureConfigurationRef
type jsiiProxy_IInfrastructureConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInfrastructureConfigurationRef) InfrastructureConfigurationRef() *InfrastructureConfigurationReference {
	var returns *InfrastructureConfigurationReference
	_jsii_.Get(
		j,
		"infrastructureConfigurationRef",
		&returns,
	)
	return returns
}

