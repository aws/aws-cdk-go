package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dataset.
// Experimental.
type IDatasetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDatasetRef
type jsiiProxy_IDatasetRef struct {
	internal.Type__constructsIConstruct
}

