package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Partition.
// Experimental.
type IPartitionRef interface {
	constructs.IConstruct
	// A reference to a Partition resource.
	// Experimental.
	PartitionRef() *PartitionReference
}

// The jsii proxy for IPartitionRef
type jsiiProxy_IPartitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPartitionRef) PartitionRef() *PartitionReference {
	var returns *PartitionReference
	_jsii_.Get(
		j,
		"partitionRef",
		&returns,
	)
	return returns
}

