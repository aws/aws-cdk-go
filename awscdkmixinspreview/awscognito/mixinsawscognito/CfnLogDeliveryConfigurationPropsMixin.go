package mixinsawscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscognito/mixinsawscognito/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Sets up or modifies the logging configuration of a user pool.
//
// User pools can export user notification logs and, when threat protection is active, user-activity logs. For more information, see [Exporting user pool logs](https://docs.aws.amazon.com/cognito/latest/developerguide/exporting-quotas-and-usage.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogDeliveryConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnLogDeliveryConfigurationPropsMixin(&CfnLogDeliveryConfigurationMixinProps{
//   	LogConfigurations: []interface{}{
//   		&LogConfigurationProperty{
//   			CloudWatchLogsConfiguration: &CloudWatchLogsConfigurationProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   			EventSource: jsii.String("eventSource"),
//   			FirehoseConfiguration: &FirehoseConfigurationProperty{
//   				StreamArn: jsii.String("streamArn"),
//   			},
//   			LogLevel: jsii.String("logLevel"),
//   			S3Configuration: &S3ConfigurationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   			},
//   		},
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html
//
type CfnLogDeliveryConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLogDeliveryConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLogDeliveryConfigurationPropsMixin
type jsiiProxy_CfnLogDeliveryConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLogDeliveryConfigurationPropsMixin) Props() *CfnLogDeliveryConfigurationMixinProps {
	var returns *CfnLogDeliveryConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogDeliveryConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::LogDeliveryConfiguration`.
func NewCfnLogDeliveryConfigurationPropsMixin(props *CfnLogDeliveryConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLogDeliveryConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLogDeliveryConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLogDeliveryConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::LogDeliveryConfiguration`.
func NewCfnLogDeliveryConfigurationPropsMixin_Override(c CfnLogDeliveryConfigurationPropsMixin, props *CfnLogDeliveryConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLogDeliveryConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogDeliveryConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogDeliveryConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogDeliveryConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLogDeliveryConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

