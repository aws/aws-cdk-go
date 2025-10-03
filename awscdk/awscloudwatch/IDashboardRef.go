package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dashboard.
// Experimental.
type IDashboardRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDashboardRef
type jsiiProxy_IDashboardRef struct {
	internal.Type__constructsIConstruct
}

