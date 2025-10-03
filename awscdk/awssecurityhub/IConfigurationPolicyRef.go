package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationPolicy.
// Experimental.
type IConfigurationPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationPolicyRef
type jsiiProxy_IConfigurationPolicyRef struct {
	internal.Type__constructsIConstruct
}

