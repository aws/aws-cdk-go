package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MLTransform.
// Experimental.
type IMLTransformRef interface {
	constructs.IConstruct
	// A reference to a MLTransform resource.
	// Experimental.
	MlTransformRef() *MLTransformReference
}

// The jsii proxy for IMLTransformRef
type jsiiProxy_IMLTransformRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMLTransformRef) MlTransformRef() *MLTransformReference {
	var returns *MLTransformReference
	_jsii_.Get(
		j,
		"mlTransformRef",
		&returns,
	)
	return returns
}

