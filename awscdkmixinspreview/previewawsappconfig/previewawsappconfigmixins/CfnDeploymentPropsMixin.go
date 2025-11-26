package previewawsappconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappconfig/previewawsappconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppConfig::Deployment` resource starts a deployment.
//
// Starting a deployment in AWS AppConfig calls the `StartDeployment` API action. This call includes the IDs of the AWS AppConfig application, the environment, the configuration profile, and (optionally) the configuration data version to deploy. The call also includes the ID of the deployment strategy to use, which determines how the configuration data is deployed.
//
// AWS AppConfig monitors the distribution to all hosts and reports status. If a distribution fails, then AWS AppConfig rolls back the configuration.
//
// AWS AppConfig requires that you create resources and deploy a configuration in the following order:
//
// - Create an application
// - Create an environment
// - Create a configuration profile
// - Choose a pre-defined deployment strategy or create your own
// - Deploy the configuration
//
// For more information, see [AWS AppConfig](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html) in the *AWS AppConfig User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeploymentPropsMixin(&CfnDeploymentMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   	ConfigurationVersion: jsii.String("configurationVersion"),
//   	DeploymentStrategyId: jsii.String("deploymentStrategyId"),
//   	Description: jsii.String("description"),
//   	DynamicExtensionParameters: []interface{}{
//   		&DynamicExtensionParametersProperty{
//   			ExtensionReference: jsii.String("extensionReference"),
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	EnvironmentId: jsii.String("environmentId"),
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html
//
type CfnDeploymentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeploymentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeploymentPropsMixin
type jsiiProxy_CfnDeploymentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeploymentPropsMixin) Props() *CfnDeploymentMixinProps {
	var returns *CfnDeploymentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppConfig::Deployment`.
func NewCfnDeploymentPropsMixin(props *CfnDeploymentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeploymentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeploymentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeploymentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnDeploymentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppConfig::Deployment`.
func NewCfnDeploymentPropsMixin_Override(c CfnDeploymentPropsMixin, props *CfnDeploymentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnDeploymentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeploymentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnDeploymentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnDeploymentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeploymentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDeploymentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

