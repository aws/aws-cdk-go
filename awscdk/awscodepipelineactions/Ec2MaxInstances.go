package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or percentage of max instances for EC2 deploy action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2MaxInstances := awscdk.Aws_codepipeline_actions.Ec2MaxInstances_Percentage(jsii.Number(123))
//
type Ec2MaxInstances interface {
	// Template value.
	Value() *string
}

// The jsii proxy struct for Ec2MaxInstances
type jsiiProxy_Ec2MaxInstances struct {
	_ byte // padding
}

func (j *jsiiProxy_Ec2MaxInstances) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewEc2MaxInstances_Override(e Ec2MaxInstances) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2MaxInstances",
		nil, // no parameters
		e,
	)
}

// Max percentage of instances.
//
// Valid range: from 1 to 100.
func Ec2MaxInstances_Percentage(percentage *float64) Ec2MaxInstances {
	_init_.Initialize()

	if err := validateEc2MaxInstances_PercentageParameters(percentage); err != nil {
		panic(err)
	}
	var returns Ec2MaxInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2MaxInstances",
		"percentage",
		[]interface{}{percentage},
		&returns,
	)

	return returns
}

// Max number of instances.
//
// Valid range: from 1 to number of your instances.
func Ec2MaxInstances_Targets(targets *float64) Ec2MaxInstances {
	_init_.Initialize()

	if err := validateEc2MaxInstances_TargetsParameters(targets); err != nil {
		panic(err)
	}
	var returns Ec2MaxInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2MaxInstances",
		"targets",
		[]interface{}{targets},
		&returns,
	)

	return returns
}

