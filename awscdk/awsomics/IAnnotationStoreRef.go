package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnnotationStore.
// Experimental.
type IAnnotationStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAnnotationStoreRef
type jsiiProxy_IAnnotationStoreRef struct {
	internal.Type__constructsIConstruct
}

