package awscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledBaseline.
// Experimental.
type IEnabledBaselineRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnabledBaselineRef
type jsiiProxy_IEnabledBaselineRef struct {
	internal.Type__constructsIConstruct
}

