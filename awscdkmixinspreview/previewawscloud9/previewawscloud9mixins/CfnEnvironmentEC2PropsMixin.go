package previewawscloud9mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloud9/previewawscloud9mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cloud9::EnvironmentEC2` resource creates an Amazon EC2 development environment in AWS Cloud9 .
//
// For more information, see [Creating an Environment](https://docs.aws.amazon.com/cloud9/latest/user-guide/create-environment.html) in the *AWS Cloud9 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentEC2PropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentEC2PropsMixin(&CfnEnvironmentEC2MixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html
//
type CfnEnvironmentEC2PropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentEC2MixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentEC2PropsMixin
type jsiiProxy_CfnEnvironmentEC2PropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnEnvironmentEC2PropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cloud9::EnvironmentEC2`.
func NewCfnEnvironmentEC2PropsMixin(props *CfnEnvironmentEC2MixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentEC2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentEC2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentEC2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cloud9::EnvironmentEC2`.
func NewCfnEnvironmentEC2PropsMixin_Override(c CfnEnvironmentEC2PropsMixin, props *CfnEnvironmentEC2MixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentEC2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentEC2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2PropsMixin",
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
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2PropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
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

