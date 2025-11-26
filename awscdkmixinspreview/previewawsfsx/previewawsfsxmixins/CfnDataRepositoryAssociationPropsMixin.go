package previewawsfsxmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfsx/previewawsfsxmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon FSx for Lustre data repository association (DRA).
//
// A data repository association is a link between a directory on the file system and an Amazon S3 bucket or prefix. You can have a maximum of 8 data repository associations on a file system. Data repository associations are supported on all FSx for Lustre 2.12 and newer file systems, excluding `scratch_1` deployment type.
//
// Each data repository association must have a unique Amazon FSx file system directory and a unique S3 bucket or prefix associated with it. You can configure a data repository association for automatic import only, for automatic export only, or for both. To learn more about linking a data repository to your file system, see [Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataRepositoryAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataRepositoryAssociationPropsMixin(&CfnDataRepositoryAssociationMixinProps{
//   	BatchImportMetaDataOnCreate: jsii.Boolean(false),
//   	DataRepositoryPath: jsii.String("dataRepositoryPath"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	FileSystemPath: jsii.String("fileSystemPath"),
//   	ImportedFileChunkSize: jsii.Number(123),
//   	S3: &S3Property{
//   		AutoExportPolicy: &AutoExportPolicyProperty{
//   			Events: []*string{
//   				jsii.String("events"),
//   			},
//   		},
//   		AutoImportPolicy: &AutoImportPolicyProperty{
//   			Events: []*string{
//   				jsii.String("events"),
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html
//
type CfnDataRepositoryAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataRepositoryAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataRepositoryAssociationPropsMixin
type jsiiProxy_CfnDataRepositoryAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataRepositoryAssociationPropsMixin) Props() *CfnDataRepositoryAssociationMixinProps {
	var returns *CfnDataRepositoryAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataRepositoryAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FSx::DataRepositoryAssociation`.
func NewCfnDataRepositoryAssociationPropsMixin(props *CfnDataRepositoryAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataRepositoryAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataRepositoryAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataRepositoryAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnDataRepositoryAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FSx::DataRepositoryAssociation`.
func NewCfnDataRepositoryAssociationPropsMixin_Override(c CfnDataRepositoryAssociationPropsMixin, props *CfnDataRepositoryAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnDataRepositoryAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataRepositoryAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataRepositoryAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnDataRepositoryAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataRepositoryAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnDataRepositoryAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataRepositoryAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataRepositoryAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

