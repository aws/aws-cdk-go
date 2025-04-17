package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Aspect that applies IMDS configuration on EC2 Launch Template constructs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateRequireImdsv2Aspect := awscdk.Aws_ec2.NewLaunchTemplateRequireImdsv2Aspect(&LaunchTemplateRequireImdsv2AspectProps{
//   	SuppressWarnings: jsii.Boolean(false),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html
//
type LaunchTemplateRequireImdsv2Aspect interface {
	awscdk.IAspect
	SuppressWarnings() *bool
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
	// Adds a warning annotation to a node, unless `suppressWarnings` is true.
	Warn(node constructs.IConstruct, message *string)
}

// The jsii proxy struct for LaunchTemplateRequireImdsv2Aspect
type jsiiProxy_LaunchTemplateRequireImdsv2Aspect struct {
	internal.Type__awscdkIAspect
}

func (j *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) SuppressWarnings() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suppressWarnings",
		&returns,
	)
	return returns
}


func NewLaunchTemplateRequireImdsv2Aspect(props *LaunchTemplateRequireImdsv2AspectProps) LaunchTemplateRequireImdsv2Aspect {
	_init_.Initialize()

	if err := validateNewLaunchTemplateRequireImdsv2AspectParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchTemplateRequireImdsv2Aspect{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.LaunchTemplateRequireImdsv2Aspect",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLaunchTemplateRequireImdsv2Aspect_Override(l LaunchTemplateRequireImdsv2Aspect, props *LaunchTemplateRequireImdsv2AspectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.LaunchTemplateRequireImdsv2Aspect",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) Visit(node constructs.IConstruct) {
	if err := l.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"visit",
		[]interface{}{node},
	)
}

func (l *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) Warn(node constructs.IConstruct, message *string) {
	if err := l.validateWarnParameters(node, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"warn",
		[]interface{}{node, message},
	)
}

