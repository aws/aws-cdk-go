package awscloud9

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloud9/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cloud9::EnvironmentEC2` resource creates an Amazon EC2 development environment in AWS Cloud9 .
//
// For more information, see [Creating an Environment](https://docs.aws.amazon.com/cloud9/latest/user-guide/create-environment.html) in the *AWS Cloud9 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEnvironmentEC2PropsMixin := awscdkcfnpropertymixins.Aws_cloud9.NewCfnEnvironmentEC2PropsMixin(&CfnEnvironmentEC2MixinProps{
//   	AutomaticStopTimeMinutes: jsii.Number(123),
//   	ConnectionType: jsii.String("connectionType"),
//   	Description: jsii.String("description"),
//   	ImageId: jsii.String("imageId"),
//   	InstanceType: jsii.String("instanceType"),
//   	Name: jsii.String("name"),
//   	OwnerArn: jsii.String("ownerArn"),
//   	Repositories: []interface{}{
//   		&RepositoryProperty{
//   			PathComponent: jsii.String("pathComponent"),
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html
//
type CfnEnvironmentEC2PropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEnvironmentEC2MixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentEC2PropsMixin
type jsiiProxy_CfnEnvironmentEC2PropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEnvironmentEC2PropsMixin) Props() *CfnEnvironmentEC2MixinProps {
	var returns *CfnEnvironmentEC2MixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2PropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cloud9::EnvironmentEC2`.
func NewCfnEnvironmentEC2PropsMixin(props *CfnEnvironmentEC2MixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEnvironmentEC2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentEC2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentEC2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloud9.CfnEnvironmentEC2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cloud9::EnvironmentEC2`.
func NewCfnEnvironmentEC2PropsMixin_Override(c CfnEnvironmentEC2PropsMixin, props *CfnEnvironmentEC2MixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloud9.CfnEnvironmentEC2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEnvironmentEC2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentEC2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloud9.CfnEnvironmentEC2PropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentEC2PropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloud9.CfnEnvironmentEC2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2PropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2PropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

