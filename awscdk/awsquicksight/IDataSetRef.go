package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataSet.
// Experimental.
type IDataSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataSetRef
type jsiiProxy_IDataSetRef struct {
	internal.Type__constructsIConstruct
}

