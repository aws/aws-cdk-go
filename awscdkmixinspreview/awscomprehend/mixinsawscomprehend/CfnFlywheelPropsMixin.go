package mixinsawscomprehend

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscomprehend/mixinsawscomprehend/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// A flywheel is an AWS resource that orchestrates the ongoing training of a model for custom classification or custom entity recognition.
//
// You can create a flywheel to start with an existing trained model, or Comprehend can create and train a new model.
//
// When you create the flywheel, Comprehend creates a data lake in your account. The data lake holds the training data and test data for all versions of the model.
//
// To use a flywheel with an existing trained model, you specify the active model version. Comprehend copies the model's training data and test data into the flywheel's data lake.
//
// To use the flywheel with a new model, you need to provide a dataset for training data (and optional test data) when you create the flywheel.
//
// For more information about flywheels, see [Flywheel overview](https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html) in the *Amazon Comprehend Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlywheelPropsMixin := awscdkmixinspreview.Mixins.NewCfnFlywheelPropsMixin(&CfnFlywheelMixinProps{
//   	ActiveModelArn: jsii.String("activeModelArn"),
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DataLakeS3Uri: jsii.String("dataLakeS3Uri"),
//   	DataSecurityConfig: &DataSecurityConfigProperty{
//   		DataLakeKmsKeyId: jsii.String("dataLakeKmsKeyId"),
//   		ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	FlywheelName: jsii.String("flywheelName"),
//   	ModelType: jsii.String("modelType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskConfig: &TaskConfigProperty{
//   		DocumentClassificationConfig: &DocumentClassificationConfigProperty{
//   			Labels: []*string{
//   				jsii.String("labels"),
//   			},
//   			Mode: jsii.String("mode"),
//   		},
//   		EntityRecognitionConfig: &EntityRecognitionConfigProperty{
//   			EntityTypes: []interface{}{
//   				&EntityTypesListItemProperty{
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		LanguageCode: jsii.String("languageCode"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html
//
type CfnFlywheelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFlywheelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlywheelPropsMixin
type jsiiProxy_CfnFlywheelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFlywheelPropsMixin) Props() *CfnFlywheelMixinProps {
	var returns *CfnFlywheelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlywheelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Comprehend::Flywheel`.
func NewCfnFlywheelPropsMixin(props *CfnFlywheelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFlywheelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlywheelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlywheelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Comprehend::Flywheel`.
func NewCfnFlywheelPropsMixin_Override(c CfnFlywheelPropsMixin, props *CfnFlywheelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFlywheelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlywheelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlywheelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlywheelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFlywheelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

