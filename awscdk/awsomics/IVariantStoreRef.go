package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VariantStore.
// Experimental.
type IVariantStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVariantStoreRef
type jsiiProxy_IVariantStoreRef struct {
	internal.Type__constructsIConstruct
}

