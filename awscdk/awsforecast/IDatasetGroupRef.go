package awsforecast

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsforecast/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DatasetGroup.
// Experimental.
type IDatasetGroupRef interface {
	constructs.IConstruct
	// A reference to a DatasetGroup resource.
	// Experimental.
	DatasetGroupRef() *DatasetGroupReference
}

// The jsii proxy for IDatasetGroupRef
type jsiiProxy_IDatasetGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDatasetGroupRef) DatasetGroupRef() *DatasetGroupReference {
	var returns *DatasetGroupReference
	_jsii_.Get(
		j,
		"datasetGroupRef",
		&returns,
	)
	return returns
}

