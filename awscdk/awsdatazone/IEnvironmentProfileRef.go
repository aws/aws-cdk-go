package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentProfile.
// Experimental.
type IEnvironmentProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnvironmentProfileRef
type jsiiProxy_IEnvironmentProfileRef struct {
	internal.Type__constructsIConstruct
}

