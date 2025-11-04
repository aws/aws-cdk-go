package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VariantStore.
// Experimental.
type IVariantStoreRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VariantStore resource.
	// Experimental.
	VariantStoreRef() *VariantStoreReference
}

// The jsii proxy for IVariantStoreRef
type jsiiProxy_IVariantStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVariantStoreRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVariantStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

