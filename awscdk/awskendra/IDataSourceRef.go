package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataSource.
// Experimental.
type IDataSourceRef interface {
	constructs.IConstruct
	// A reference to a DataSource resource.
	// Experimental.
	DataSourceRef() *DataSourceReference
}

// The jsii proxy for IDataSourceRef
type jsiiProxy_IDataSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataSourceRef) DataSourceRef() *DataSourceReference {
	var returns *DataSourceReference
	_jsii_.Get(
		j,
		"dataSourceRef",
		&returns,
	)
	return returns
}

