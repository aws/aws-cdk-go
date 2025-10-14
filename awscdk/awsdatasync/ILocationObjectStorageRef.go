package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationObjectStorage.
// Experimental.
type ILocationObjectStorageRef interface {
	constructs.IConstruct
	// A reference to a LocationObjectStorage resource.
	// Experimental.
	LocationObjectStorageRef() *LocationObjectStorageReference
}

// The jsii proxy for ILocationObjectStorageRef
type jsiiProxy_ILocationObjectStorageRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocationObjectStorageRef) LocationObjectStorageRef() *LocationObjectStorageReference {
	var returns *LocationObjectStorageReference
	_jsii_.Get(
		j,
		"locationObjectStorageRef",
		&returns,
	)
	return returns
}

