package awsdatasync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::LocationFSxWindows` resource specifies an endpoint for an Amazon FSx for Windows Server file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLocationFSxWindowsPropsMixin := awscdkcfnpropertymixins.Aws_datasync.NewCfnLocationFSxWindowsPropsMixin(&CfnLocationFSxWindowsMixinProps{
//   	CmkSecretConfig: &CmkSecretConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CustomSecretConfig: &CustomSecretConfigProperty{
//   		SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	Domain: jsii.String("domain"),
//   	FsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	Password: jsii.String("password"),
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	User: jsii.String("user"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html
//
type CfnLocationFSxWindowsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLocationFSxWindowsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationFSxWindowsPropsMixin
type jsiiProxy_CfnLocationFSxWindowsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLocationFSxWindowsPropsMixin) Props() *CfnLocationFSxWindowsMixinProps {
	var returns *CfnLocationFSxWindowsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindowsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationFSxWindows`.
func NewCfnLocationFSxWindowsPropsMixin(props *CfnLocationFSxWindowsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLocationFSxWindowsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationFSxWindowsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationFSxWindowsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationFSxWindows`.
func NewCfnLocationFSxWindowsPropsMixin_Override(c CfnLocationFSxWindowsPropsMixin, props *CfnLocationFSxWindowsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLocationFSxWindowsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationFSxWindowsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationFSxWindowsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationFSxWindowsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLocationFSxWindowsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

