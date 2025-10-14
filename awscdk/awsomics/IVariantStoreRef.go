package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VariantStore.
// Experimental.
type IVariantStoreRef interface {
	constructs.IConstruct
	// A reference to a VariantStore resource.
	// Experimental.
	VariantStoreRef() *VariantStoreReference
}

// The jsii proxy for IVariantStoreRef
type jsiiProxy_IVariantStoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVariantStoreRef) VariantStoreRef() *VariantStoreReference {
	var returns *VariantStoreReference
	_jsii_.Get(
		j,
		"variantStoreRef",
		&returns,
	)
	return returns
}

