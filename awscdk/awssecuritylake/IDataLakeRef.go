package awssecuritylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLake.
// Experimental.
type IDataLakeRef interface {
	constructs.IConstruct
	// A reference to a DataLake resource.
	// Experimental.
	DataLakeRef() *DataLakeReference
}

// The jsii proxy for IDataLakeRef
type jsiiProxy_IDataLakeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataLakeRef) DataLakeRef() *DataLakeReference {
	var returns *DataLakeReference
	_jsii_.Get(
		j,
		"dataLakeRef",
		&returns,
	)
	return returns
}

