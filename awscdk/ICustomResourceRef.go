package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomResource.
// Experimental.
type ICustomResourceRef interface {
	constructs.IConstruct
	// A reference to a CustomResource resource.
	// Experimental.
	CustomResourceRef() *CustomResourceReference
}

// The jsii proxy for ICustomResourceRef
type jsiiProxy_ICustomResourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomResourceRef) CustomResourceRef() *CustomResourceReference {
	var returns *CustomResourceReference
	_jsii_.Get(
		j,
		"customResourceRef",
		&returns,
	)
	return returns
}

