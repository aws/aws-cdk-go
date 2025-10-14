package awsdatabrew

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dataset.
// Experimental.
type IDatasetRef interface {
	constructs.IConstruct
	// A reference to a Dataset resource.
	// Experimental.
	DatasetRef() *DatasetReference
}

// The jsii proxy for IDatasetRef
type jsiiProxy_IDatasetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDatasetRef) DatasetRef() *DatasetReference {
	var returns *DatasetReference
	_jsii_.Get(
		j,
		"datasetRef",
		&returns,
	)
	return returns
}

