package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceGroup.
// Experimental.
type IResourceGroupRef interface {
	constructs.IConstruct
	// A reference to a ResourceGroup resource.
	// Experimental.
	ResourceGroupRef() *ResourceGroupReference
}

// The jsii proxy for IResourceGroupRef
type jsiiProxy_IResourceGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceGroupRef) ResourceGroupRef() *ResourceGroupReference {
	var returns *ResourceGroupReference
	_jsii_.Get(
		j,
		"resourceGroupRef",
		&returns,
	)
	return returns
}

