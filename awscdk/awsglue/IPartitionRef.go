package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Partition.
// Experimental.
type IPartitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Partition resource.
	// Experimental.
	PartitionRef() *PartitionReference
}

// The jsii proxy for IPartitionRef
type jsiiProxy_IPartitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IPartitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

