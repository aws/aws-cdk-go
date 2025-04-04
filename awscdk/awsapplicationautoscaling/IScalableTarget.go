package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
)

type IScalableTarget interface {
	awscdk.IResource
	ScalableTargetId() *string
}

// The jsii proxy for IScalableTarget
type jsiiProxy_IScalableTarget struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IScalableTarget) ScalableTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableTargetId",
		&returns,
	)
	return returns
}

