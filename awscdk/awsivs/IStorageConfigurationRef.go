package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageConfiguration.
// Experimental.
type IStorageConfigurationRef interface {
	constructs.IConstruct
	// A reference to a StorageConfiguration resource.
	// Experimental.
	StorageConfigurationRef() *StorageConfigurationReference
}

// The jsii proxy for IStorageConfigurationRef
type jsiiProxy_IStorageConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStorageConfigurationRef) StorageConfigurationRef() *StorageConfigurationReference {
	var returns *StorageConfigurationReference
	_jsii_.Get(
		j,
		"storageConfigurationRef",
		&returns,
	)
	return returns
}

