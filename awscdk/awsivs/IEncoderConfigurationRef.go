package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EncoderConfiguration.
// Experimental.
type IEncoderConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEncoderConfigurationRef
type jsiiProxy_IEncoderConfigurationRef struct {
	internal.Type__constructsIConstruct
}

