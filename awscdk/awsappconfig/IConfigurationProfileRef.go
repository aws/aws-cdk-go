package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationProfile.
// Experimental.
type IConfigurationProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationProfileRef
type jsiiProxy_IConfigurationProfileRef struct {
	internal.Type__constructsIConstruct
}

