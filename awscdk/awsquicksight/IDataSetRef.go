package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataSet.
// Experimental.
type IDataSetRef interface {
	constructs.IConstruct
	// A reference to a DataSet resource.
	// Experimental.
	DataSetRef() *DataSetReference
}

// The jsii proxy for IDataSetRef
type jsiiProxy_IDataSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataSetRef) DataSetRef() *DataSetReference {
	var returns *DataSetReference
	_jsii_.Get(
		j,
		"dataSetRef",
		&returns,
	)
	return returns
}

