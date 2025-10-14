package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageProfile.
// Experimental.
type IStorageProfileRef interface {
	constructs.IConstruct
	// A reference to a StorageProfile resource.
	// Experimental.
	StorageProfileRef() *StorageProfileReference
}

// The jsii proxy for IStorageProfileRef
type jsiiProxy_IStorageProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStorageProfileRef) StorageProfileRef() *StorageProfileReference {
	var returns *StorageProfileReference
	_jsii_.Get(
		j,
		"storageProfileRef",
		&returns,
	)
	return returns
}

