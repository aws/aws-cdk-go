package awspersonalize

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DatasetGroup.
// Experimental.
type IDatasetGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDatasetGroupRef
type jsiiProxy_IDatasetGroupRef struct {
	internal.Type__constructsIConstruct
}

