package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalableTarget.
// Experimental.
type IScalableTargetRef interface {
	constructs.IConstruct
	// A reference to a ScalableTarget resource.
	// Experimental.
	ScalableTargetRef() *ScalableTargetReference
}

// The jsii proxy for IScalableTargetRef
type jsiiProxy_IScalableTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScalableTargetRef) ScalableTargetRef() *ScalableTargetReference {
	var returns *ScalableTargetReference
	_jsii_.Get(
		j,
		"scalableTargetRef",
		&returns,
	)
	return returns
}

