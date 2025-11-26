package previewawscodebuildmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodebuild/previewawscodebuildmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a report group.
//
// A report group contains a collection of reports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReportGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnReportGroupPropsMixin(&CfnReportGroupMixinProps{
//   	DeleteReports: jsii.Boolean(false),
//   	ExportConfig: &ReportExportConfigProperty{
//   		ExportConfigType: jsii.String("exportConfigType"),
//   		S3Destination: &S3ReportExportConfigProperty{
//   			Bucket: jsii.String("bucket"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			EncryptionDisabled: jsii.Boolean(false),
//   			EncryptionKey: jsii.String("encryptionKey"),
//   			Packaging: jsii.String("packaging"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-reportgroup.html
//
type CfnReportGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReportGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReportGroupPropsMixin
type jsiiProxy_CfnReportGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReportGroupPropsMixin) Props() *CfnReportGroupMixinProps {
	var returns *CfnReportGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeBuild::ReportGroup`.
func NewCfnReportGroupPropsMixin(props *CfnReportGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReportGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReportGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReportGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeBuild::ReportGroup`.
func NewCfnReportGroupPropsMixin_Override(c CfnReportGroupPropsMixin, props *CfnReportGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReportGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReportGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReportGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReportGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnReportGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

