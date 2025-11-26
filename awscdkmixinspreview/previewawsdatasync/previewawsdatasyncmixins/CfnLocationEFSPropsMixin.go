package previewawsdatasyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatasync/previewawsdatasyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::LocationEFS` resource creates an endpoint for an Amazon EFS file system.
//
// AWS DataSync can access this endpoint as a source or destination location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationEFSPropsMixin := awscdkmixinspreview.Mixins.NewCfnLocationEFSPropsMixin(&CfnLocationEFSMixinProps{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	Ec2Config: &Ec2ConfigProperty{
//   		SecurityGroupArns: []*string{
//   			jsii.String("securityGroupArns"),
//   		},
//   		SubnetArn: jsii.String("subnetArn"),
//   	},
//   	EfsFilesystemArn: jsii.String("efsFilesystemArn"),
//   	FileSystemAccessRoleArn: jsii.String("fileSystemAccessRoleArn"),
//   	InTransitEncryption: jsii.String("inTransitEncryption"),
//   	Subdirectory: jsii.String("subdirectory"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html
//
type CfnLocationEFSPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLocationEFSMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationEFSPropsMixin
type jsiiProxy_CfnLocationEFSPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLocationEFSPropsMixin) Props() *CfnLocationEFSMixinProps {
	var returns *CfnLocationEFSMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFSPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationEFS`.
func NewCfnLocationEFSPropsMixin(props *CfnLocationEFSMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLocationEFSPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationEFSPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationEFSPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationEFS`.
func NewCfnLocationEFSPropsMixin_Override(c CfnLocationEFSPropsMixin, props *CfnLocationEFSMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLocationEFSPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationEFSPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationEFSPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationEFSPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLocationEFSPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

