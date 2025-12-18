package previewawsobservabilityadminmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsobservabilityadmin/previewawsobservabilityadminmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an integration between CloudWatch and S3 Tables for analytics.
//
// This integration enables querying CloudWatch telemetry data using analytics engines like Amazon Athena, Amazon Redshift, and Apache Spark.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnS3TableIntegrationPropsMixin := awscdkmixinspreview.Mixins.NewCfnS3TableIntegrationPropsMixin(&CfnS3TableIntegrationMixinProps{
//   	Encryption: &EncryptionConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//   	},
//   	LogSources: []interface{}{
//   		&LogSourceProperty{
//   			Identifier: jsii.String("identifier"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html
//
type CfnS3TableIntegrationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnS3TableIntegrationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnS3TableIntegrationPropsMixin
type jsiiProxy_CfnS3TableIntegrationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnS3TableIntegrationPropsMixin) Props() *CfnS3TableIntegrationMixinProps {
	var returns *CfnS3TableIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnS3TableIntegrationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::S3TableIntegration`.
func NewCfnS3TableIntegrationPropsMixin(props *CfnS3TableIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnS3TableIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnS3TableIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnS3TableIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::S3TableIntegration`.
func NewCfnS3TableIntegrationPropsMixin_Override(c CfnS3TableIntegrationPropsMixin, props *CfnS3TableIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnS3TableIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnS3TableIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnS3TableIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnS3TableIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnS3TableIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

