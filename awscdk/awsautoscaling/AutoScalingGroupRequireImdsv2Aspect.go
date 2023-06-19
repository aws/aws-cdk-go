package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling/internal"
)

// Aspect that makes IMDSv2 required on instances deployed by AutoScalingGroups.
//
// Example:
//   aspect := autoscaling.NewAutoScalingGroupRequireImdsv2Aspect()
//
//   awscdk.Aspects_Of(this).Add(aspect)
//
// Experimental.
type AutoScalingGroupRequireImdsv2Aspect interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node awscdk.IConstruct)
	// Adds a warning annotation to a node.
	// Experimental.
	Warn(node awscdk.IConstruct, message *string)
}

// The jsii proxy struct for AutoScalingGroupRequireImdsv2Aspect
type jsiiProxy_AutoScalingGroupRequireImdsv2Aspect struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewAutoScalingGroupRequireImdsv2Aspect() AutoScalingGroupRequireImdsv2Aspect {
	_init_.Initialize()

	j := jsiiProxy_AutoScalingGroupRequireImdsv2Aspect{}

	_jsii_.Create(
		"monocdk.aws_autoscaling.AutoScalingGroupRequireImdsv2Aspect",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingGroupRequireImdsv2Aspect_Override(a AutoScalingGroupRequireImdsv2Aspect) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.AutoScalingGroupRequireImdsv2Aspect",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) Visit(node awscdk.IConstruct) {
	if err := a.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"visit",
		[]interface{}{node},
	)
}

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) Warn(node awscdk.IConstruct, message *string) {
	if err := a.validateWarnParameters(node, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"warn",
		[]interface{}{node, message},
	)
}

