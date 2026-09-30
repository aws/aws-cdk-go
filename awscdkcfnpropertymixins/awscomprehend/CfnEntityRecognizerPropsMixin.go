package awscomprehend

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amazon Comprehend custom entity recognizer: a trained model that identifies custom entity types in text, created via an asynchronous training job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEntityRecognizerPropsMixin := awscdkcfnpropertymixins.Aws_comprehend.NewCfnEntityRecognizerPropsMixin(&CfnEntityRecognizerMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	InputDataConfig: &EntityRecognizerInputDataConfigProperty{
//   		Annotations: &EntityRecognizerAnnotationsProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   			TestS3Uri: jsii.String("testS3Uri"),
//   		},
//   		AugmentedManifests: []interface{}{
//   			&AugmentedManifestsListItemProperty{
//   				AnnotationDataS3Uri: jsii.String("annotationDataS3Uri"),
//   				AttributeNames: []*string{
//   					jsii.String("attributeNames"),
//   				},
//   				DocumentType: jsii.String("documentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   				SourceDocumentsS3Uri: jsii.String("sourceDocumentsS3Uri"),
//   				Split: jsii.String("split"),
//   			},
//   		},
//   		DataFormat: jsii.String("dataFormat"),
//   		Documents: &EntityRecognizerDocumentsProperty{
//   			InputFormat: jsii.String("inputFormat"),
//   			S3Uri: jsii.String("s3Uri"),
//   			TestS3Uri: jsii.String("testS3Uri"),
//   		},
//   		EntityList: &EntityRecognizerEntityListProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		EntityTypes: []interface{}{
//   			&EntityTypesListItemProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   	ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   	ModelPolicy: jsii.String("modelPolicy"),
//   	RecognizerName: jsii.String("recognizerName"),
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html
//
type CfnEntityRecognizerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEntityRecognizerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEntityRecognizerPropsMixin
type jsiiProxy_CfnEntityRecognizerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEntityRecognizerPropsMixin) Props() *CfnEntityRecognizerMixinProps {
	var returns *CfnEntityRecognizerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntityRecognizerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Comprehend::EntityRecognizer`.
func NewCfnEntityRecognizerPropsMixin(props *CfnEntityRecognizerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEntityRecognizerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEntityRecognizerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEntityRecognizerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnEntityRecognizerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Comprehend::EntityRecognizer`.
func NewCfnEntityRecognizerPropsMixin_Override(c CfnEntityRecognizerPropsMixin, props *CfnEntityRecognizerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnEntityRecognizerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEntityRecognizerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEntityRecognizerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnEntityRecognizerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEntityRecognizerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnEntityRecognizerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEntityRecognizerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEntityRecognizerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

