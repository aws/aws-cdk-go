package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainConfiguration.
// Experimental.
type IDomainConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDomainConfigurationRef
type jsiiProxy_IDomainConfigurationRef struct {
	internal.Type__constructsIConstruct
}

