package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Archive.
// Experimental.
type IArchiveRef interface {
	constructs.IConstruct
	// A reference to a Archive resource.
	// Experimental.
	ArchiveRef() *ArchiveReference
}

// The jsii proxy for IArchiveRef
type jsiiProxy_IArchiveRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IArchiveRef) ArchiveRef() *ArchiveReference {
	var returns *ArchiveReference
	_jsii_.Get(
		j,
		"archiveRef",
		&returns,
	)
	return returns
}

