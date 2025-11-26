package previewawscomprehendmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscomprehend/previewawscomprehendmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource creates and trains a document classifier to categorize documents.
//
// You provide a set of training documents that are labeled with the categories that you want to identify. After the classifier is trained you can use it to categorize a set of labeled documents into the categories. For more information, see [Document Classification](https://docs.aws.amazon.com/comprehend/latest/dg/how-document-classification.html) in the Comprehend Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDocumentClassifierPropsMixin := awscdkmixinspreview.Mixins.NewCfnDocumentClassifierPropsMixin(&CfnDocumentClassifierMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DocumentClassifierName: jsii.String("documentClassifierName"),
//   	InputDataConfig: &DocumentClassifierInputDataConfigProperty{
//   		AugmentedManifests: []interface{}{
//   			&AugmentedManifestsListItemProperty{
//   				AttributeNames: []*string{
//   					jsii.String("attributeNames"),
//   				},
//   				S3Uri: jsii.String("s3Uri"),
//   				Split: jsii.String("split"),
//   			},
//   		},
//   		DataFormat: jsii.String("dataFormat"),
//   		DocumentReaderConfig: &DocumentReaderConfigProperty{
//   			DocumentReadAction: jsii.String("documentReadAction"),
//   			DocumentReadMode: jsii.String("documentReadMode"),
//   			FeatureTypes: []*string{
//   				jsii.String("featureTypes"),
//   			},
//   		},
//   		Documents: &DocumentClassifierDocumentsProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   			TestS3Uri: jsii.String("testS3Uri"),
//   		},
//   		DocumentType: jsii.String("documentType"),
//   		LabelDelimiter: jsii.String("labelDelimiter"),
//   		S3Uri: jsii.String("s3Uri"),
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   	Mode: jsii.String("mode"),
//   	ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   	ModelPolicy: jsii.String("modelPolicy"),
//   	OutputDataConfig: &DocumentClassifierOutputDataConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersionName: jsii.String("versionName"),
//   	VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html
//
type CfnDocumentClassifierPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDocumentClassifierMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDocumentClassifierPropsMixin
type jsiiProxy_CfnDocumentClassifierPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDocumentClassifierPropsMixin) Props() *CfnDocumentClassifierMixinProps {
	var returns *CfnDocumentClassifierMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentClassifierPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Comprehend::DocumentClassifier`.
func NewCfnDocumentClassifierPropsMixin(props *CfnDocumentClassifierMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDocumentClassifierPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDocumentClassifierPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDocumentClassifierPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Comprehend::DocumentClassifier`.
func NewCfnDocumentClassifierPropsMixin_Override(c CfnDocumentClassifierPropsMixin, props *CfnDocumentClassifierMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDocumentClassifierPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDocumentClassifierPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDocumentClassifierPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDocumentClassifierPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDocumentClassifierPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

