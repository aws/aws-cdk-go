package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Disk.
// Experimental.
type IDiskRef interface {
	constructs.IConstruct
	// A reference to a Disk resource.
	// Experimental.
	DiskRef() *DiskReference
}

// The jsii proxy for IDiskRef
type jsiiProxy_IDiskRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDiskRef) DiskRef() *DiskReference {
	var returns *DiskReference
	_jsii_.Get(
		j,
		"diskRef",
		&returns,
	)
	return returns
}

