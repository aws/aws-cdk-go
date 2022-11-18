package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
)

// Aspect that applies IMDS configuration on EC2 Instance constructs.
//
// This aspect configures IMDS on an EC2 instance by creating a Launch Template with the
// IMDS configuration and associating that Launch Template with the instance. If an Instance
// is already associated with a Launch Template, a warning will (optionally) be added to the
// construct node and it will be skipped.
//
// To cover Instances already associated with Launch Templates, use `LaunchTemplateImdsAspect`.
//
// Example:
//   aspect := ec2.NewInstanceRequireImdsv2Aspect()
//   awscdk.Aspects.of(this).add(aspect)
//
// Experimental.
type InstanceRequireImdsv2Aspect interface {
	awscdk.IAspect
	// Experimental.
	SuppressWarnings() *bool
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node awscdk.IConstruct)
	// Adds a warning annotation to a node, unless `suppressWarnings` is true.
	// Experimental.
	Warn(node awscdk.IConstruct, message *string)
}

// The jsii proxy struct for InstanceRequireImdsv2Aspect
type jsiiProxy_InstanceRequireImdsv2Aspect struct {
	internal.Type__awscdkIAspect
}

func (j *jsiiProxy_InstanceRequireImdsv2Aspect) SuppressWarnings() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suppressWarnings",
		&returns,
	)
	return returns
}


// Experimental.
func NewInstanceRequireImdsv2Aspect(props *InstanceRequireImdsv2AspectProps) InstanceRequireImdsv2Aspect {
	_init_.Initialize()

	if err := validateNewInstanceRequireImdsv2AspectParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceRequireImdsv2Aspect{}

	_jsii_.Create(
		"monocdk.aws_ec2.InstanceRequireImdsv2Aspect",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceRequireImdsv2Aspect_Override(i InstanceRequireImdsv2Aspect, props *InstanceRequireImdsv2AspectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.InstanceRequireImdsv2Aspect",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_InstanceRequireImdsv2Aspect) Visit(node awscdk.IConstruct) {
	if err := i.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"visit",
		[]interface{}{node},
	)
}

func (i *jsiiProxy_InstanceRequireImdsv2Aspect) Warn(node awscdk.IConstruct, message *string) {
	if err := i.validateWarnParameters(node, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"warn",
		[]interface{}{node, message},
	)
}

