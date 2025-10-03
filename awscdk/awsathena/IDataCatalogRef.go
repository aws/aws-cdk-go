package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCatalog.
// Experimental.
type IDataCatalogRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataCatalogRef
type jsiiProxy_IDataCatalogRef struct {
	internal.Type__constructsIConstruct
}

