package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SnapshotBlockPublicAccess.
// Experimental.
type ISnapshotBlockPublicAccessRef interface {
	constructs.IConstruct
	// A reference to a SnapshotBlockPublicAccess resource.
	// Experimental.
	SnapshotBlockPublicAccessRef() *SnapshotBlockPublicAccessReference
}

// The jsii proxy for ISnapshotBlockPublicAccessRef
type jsiiProxy_ISnapshotBlockPublicAccessRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISnapshotBlockPublicAccessRef) SnapshotBlockPublicAccessRef() *SnapshotBlockPublicAccessReference {
	var returns *SnapshotBlockPublicAccessReference
	_jsii_.Get(
		j,
		"snapshotBlockPublicAccessRef",
		&returns,
	)
	return returns
}

