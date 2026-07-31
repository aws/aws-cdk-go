package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for propagating tags to managed resources created by a capacity provider.
//
// Use static factory methods to create instances:
// ```
// propagateTags: lambda.PropagateTags.explicit({ env: 'prod', team: 'platform' })
// ```.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"))
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   capacityProvider := lambda.NewCapacityProvider(this, jsii.String("MyCapacityProvider"), &CapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	PropagateTags: lambda.PropagateTags_Explicit(map[string]*string{
//   		"CostCenter": jsii.String("Engineering"),
//   		"Project": jsii.String("MyProject"),
//   	}),
//   })
//
type PropagateTags interface {
	// The explicit tags to propagate (only for Explicit mode).
	ExplicitTags() *map[string]*string
	// The propagation mode.
	Mode() *string
}

// The jsii proxy struct for PropagateTags
type jsiiProxy_PropagateTags struct {
	_ byte // padding
}

func (j *jsiiProxy_PropagateTags) ExplicitTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"explicitTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropagateTags) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// Propagate the specified tags to all managed resources.
func PropagateTags_Explicit(tags *map[string]*string) PropagateTags {
	_init_.Initialize()

	if err := validatePropagateTags_ExplicitParameters(tags); err != nil {
		panic(err)
	}
	var returns PropagateTags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.PropagateTags",
		"explicit",
		[]interface{}{tags},
		&returns,
	)

	return returns
}

// No tag propagation to managed resources.
func PropagateTags_None() PropagateTags {
	_init_.Initialize()

	var returns PropagateTags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.PropagateTags",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

