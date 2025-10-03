package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReferenceStore.
// Experimental.
type IReferenceStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReferenceStoreRef
type jsiiProxy_IReferenceStoreRef struct {
	internal.Type__constructsIConstruct
}

