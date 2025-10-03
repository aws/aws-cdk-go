package awscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledControl.
// Experimental.
type IEnabledControlRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnabledControlRef
type jsiiProxy_IEnabledControlRef struct {
	internal.Type__constructsIConstruct
}

