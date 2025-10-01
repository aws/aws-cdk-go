package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MigrationProject.
// Experimental.
type IMigrationProjectRef interface {
	constructs.IConstruct
	// A reference to a MigrationProject resource.
	// Experimental.
	MigrationProjectRef() *MigrationProjectReference
}

// The jsii proxy for IMigrationProjectRef
type jsiiProxy_IMigrationProjectRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMigrationProjectRef) MigrationProjectRef() *MigrationProjectReference {
	var returns *MigrationProjectReference
	_jsii_.Get(
		j,
		"migrationProjectRef",
		&returns,
	)
	return returns
}

