package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomActionType.
// Experimental.
type ICustomActionTypeRef interface {
	constructs.IConstruct
	// A reference to a CustomActionType resource.
	// Experimental.
	CustomActionTypeRef() *CustomActionTypeReference
}

// The jsii proxy for ICustomActionTypeRef
type jsiiProxy_ICustomActionTypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomActionTypeRef) CustomActionTypeRef() *CustomActionTypeReference {
	var returns *CustomActionTypeReference
	_jsii_.Get(
		j,
		"customActionTypeRef",
		&returns,
	)
	return returns
}

