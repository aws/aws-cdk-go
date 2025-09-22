package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceStorageConfig.
// Experimental.
type IInstanceStorageConfigRef interface {
	constructs.IConstruct
	// A reference to a InstanceStorageConfig resource.
	// Experimental.
	InstanceStorageConfigRef() *InstanceStorageConfigReference
}

// The jsii proxy for IInstanceStorageConfigRef
type jsiiProxy_IInstanceStorageConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceStorageConfigRef) InstanceStorageConfigRef() *InstanceStorageConfigReference {
	var returns *InstanceStorageConfigReference
	_jsii_.Get(
		j,
		"instanceStorageConfigRef",
		&returns,
	)
	return returns
}

