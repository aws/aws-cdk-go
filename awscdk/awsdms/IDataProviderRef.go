package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProvider.
// Experimental.
type IDataProviderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataProviderRef
type jsiiProxy_IDataProviderRef struct {
	internal.Type__constructsIConstruct
}

