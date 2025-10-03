package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalCatalog.
// Experimental.
type ISignalCatalogRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISignalCatalogRef
type jsiiProxy_ISignalCatalogRef struct {
	internal.Type__constructsIConstruct
}

