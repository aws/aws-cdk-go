package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProvider.
// Experimental.
type IDataProviderRef interface {
	constructs.IConstruct
	// A reference to a DataProvider resource.
	// Experimental.
	DataProviderRef() *DataProviderReference
}

// The jsii proxy for IDataProviderRef
type jsiiProxy_IDataProviderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataProviderRef) DataProviderRef() *DataProviderReference {
	var returns *DataProviderReference
	_jsii_.Get(
		j,
		"dataProviderRef",
		&returns,
	)
	return returns
}

