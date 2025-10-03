package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MigrationProject.
// Experimental.
type IMigrationProjectRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMigrationProjectRef
type jsiiProxy_IMigrationProjectRef struct {
	internal.Type__constructsIConstruct
}

