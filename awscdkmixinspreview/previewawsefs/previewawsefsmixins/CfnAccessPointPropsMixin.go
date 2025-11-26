package previewawsefsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsefs/previewawsefsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::EFS::AccessPoint` resource creates an EFS access point.
//
// An access point is an application-specific view into an EFS file system that applies an operating system user and group, and a file system path, to any file system request made through the access point. The operating system user and group override any identity information provided by the NFS client. The file system path is exposed as the access point's root directory. Applications using the access point can only access data in its own directory and below. To learn more, see [Mounting a file system using EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) .
//
// This operation requires permissions for the `elasticfilesystem:CreateAccessPoint` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessPointPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessPointPropsMixin(&CfnAccessPointMixinProps{
//   	AccessPointTags: []AccessPointTagProperty{
//   		&AccessPointTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ClientToken: jsii.String("clientToken"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	PosixUser: &PosixUserProperty{
//   		Gid: jsii.String("gid"),
//   		SecondaryGids: []*string{
//   			jsii.String("secondaryGids"),
//   		},
//   		Uid: jsii.String("uid"),
//   	},
//   	RootDirectory: &RootDirectoryProperty{
//   		CreationInfo: &CreationInfoProperty{
//   			OwnerGid: jsii.String("ownerGid"),
//   			OwnerUid: jsii.String("ownerUid"),
//   			Permissions: jsii.String("permissions"),
//   		},
//   		Path: jsii.String("path"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html
//
type CfnAccessPointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessPointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessPointPropsMixin
type jsiiProxy_CfnAccessPointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessPointPropsMixin) Props() *CfnAccessPointMixinProps {
	var returns *CfnAccessPointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EFS::AccessPoint`.
func NewCfnAccessPointPropsMixin(props *CfnAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessPointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessPointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessPointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EFS::AccessPoint`.
func NewCfnAccessPointPropsMixin_Override(c CfnAccessPointPropsMixin, props *CfnAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessPointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessPointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessPointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessPointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccessPointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

