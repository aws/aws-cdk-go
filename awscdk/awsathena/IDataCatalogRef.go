package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCatalog.
// Experimental.
type IDataCatalogRef interface {
	constructs.IConstruct
	// A reference to a DataCatalog resource.
	// Experimental.
	DataCatalogRef() *DataCatalogReference
}

// The jsii proxy for IDataCatalogRef
type jsiiProxy_IDataCatalogRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataCatalogRef) DataCatalogRef() *DataCatalogReference {
	var returns *DataCatalogReference
	_jsii_.Get(
		j,
		"dataCatalogRef",
		&returns,
	)
	return returns
}

