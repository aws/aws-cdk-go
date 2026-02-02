package previewawsfsxmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfsx/previewawsfsxmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An S3 access point attached to an Amazon FSx volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnS3AccessPointAttachmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnS3AccessPointAttachmentPropsMixin(&CfnS3AccessPointAttachmentMixinProps{
//   	Name: jsii.String("name"),
//   	OntapConfiguration: &S3AccessPointOntapConfigurationProperty{
//   		FileSystemIdentity: &OntapFileSystemIdentityProperty{
//   			Type: jsii.String("type"),
//   			UnixUser: &OntapUnixFileSystemUserProperty{
//   				Name: jsii.String("name"),
//   			},
//   			WindowsUser: &OntapWindowsFileSystemUserProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		VolumeId: jsii.String("volumeId"),
//   	},
//   	OpenZfsConfiguration: &S3AccessPointOpenZFSConfigurationProperty{
//   		FileSystemIdentity: &OpenZFSFileSystemIdentityProperty{
//   			PosixUser: &OpenZFSPosixFileSystemUserProperty{
//   				Gid: jsii.Number(123),
//   				SecondaryGids: []interface{}{
//   					&FileSystemGIDProperty{
//   						Gid: jsii.Number(123),
//   					},
//   				},
//   				Uid: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		VolumeId: jsii.String("volumeId"),
//   	},
//   	S3AccessPoint: &S3AccessPointProperty{
//   		Alias: jsii.String("alias"),
//   		Policy: policy,
//   		ResourceArn: jsii.String("resourceArn"),
//   		VpcConfiguration: &S3AccessPointVpcConfigurationProperty{
//   			VpcId: jsii.String("vpcId"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html
//
type CfnS3AccessPointAttachmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnS3AccessPointAttachmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnS3AccessPointAttachmentPropsMixin
type jsiiProxy_CfnS3AccessPointAttachmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnS3AccessPointAttachmentPropsMixin) Props() *CfnS3AccessPointAttachmentMixinProps {
	var returns *CfnS3AccessPointAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnS3AccessPointAttachmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FSx::S3AccessPointAttachment`.
func NewCfnS3AccessPointAttachmentPropsMixin(props *CfnS3AccessPointAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnS3AccessPointAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnS3AccessPointAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnS3AccessPointAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnS3AccessPointAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FSx::S3AccessPointAttachment`.
func NewCfnS3AccessPointAttachmentPropsMixin_Override(c CfnS3AccessPointAttachmentPropsMixin, props *CfnS3AccessPointAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnS3AccessPointAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnS3AccessPointAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnS3AccessPointAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnS3AccessPointAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnS3AccessPointAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnS3AccessPointAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnS3AccessPointAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnS3AccessPointAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

