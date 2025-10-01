package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DistributionConfiguration.
// Experimental.
type IDistributionConfigurationRef interface {
	constructs.IConstruct
	// A reference to a DistributionConfiguration resource.
	// Experimental.
	DistributionConfigurationRef() *DistributionConfigurationReference
}

// The jsii proxy for IDistributionConfigurationRef
type jsiiProxy_IDistributionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDistributionConfigurationRef) DistributionConfigurationRef() *DistributionConfigurationReference {
	var returns *DistributionConfigurationReference
	_jsii_.Get(
		j,
		"distributionConfigurationRef",
		&returns,
	)
	return returns
}

