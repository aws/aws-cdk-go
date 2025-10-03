package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SequenceStore.
// Experimental.
type ISequenceStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISequenceStoreRef
type jsiiProxy_ISequenceStoreRef struct {
	internal.Type__constructsIConstruct
}

