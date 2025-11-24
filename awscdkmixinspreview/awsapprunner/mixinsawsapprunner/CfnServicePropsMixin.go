package mixinsawsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsapprunner/mixinsawsapprunner/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify an AWS App Runner service by using the `AWS::AppRunner::Service` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::Service` resource is an AWS App Runner resource type that specifies an App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServicePropsMixin := awscdkmixinspreview.Mixins.NewCfnServicePropsMixin(&CfnServiceMixinProps{
//   	AutoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	HealthCheckConfiguration: &HealthCheckConfigurationProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		Interval: jsii.Number(123),
//   		Path: jsii.String("path"),
//   		Protocol: jsii.String("protocol"),
//   		Timeout: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//   	},
//   	InstanceConfiguration: &InstanceConfigurationProperty{
//   		Cpu: jsii.String("cpu"),
//   		InstanceRoleArn: jsii.String("instanceRoleArn"),
//   		Memory: jsii.String("memory"),
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		EgressConfiguration: &EgressConfigurationProperty{
//   			EgressType: jsii.String("egressType"),
//   			VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   		IngressConfiguration: &IngressConfigurationProperty{
//   			IsPubliclyAccessible: jsii.Boolean(false),
//   		},
//   		IpAddressType: jsii.String("ipAddressType"),
//   	},
//   	ObservabilityConfiguration: &ServiceObservabilityConfigurationProperty{
//   		ObservabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   		ObservabilityEnabled: jsii.Boolean(false),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   			AccessRoleArn: jsii.String("accessRoleArn"),
//   			ConnectionArn: jsii.String("connectionArn"),
//   		},
//   		AutoDeploymentsEnabled: jsii.Boolean(false),
//   		CodeRepository: &CodeRepositoryProperty{
//   			CodeConfiguration: &CodeConfigurationProperty{
//   				CodeConfigurationValues: &CodeConfigurationValuesProperty{
//   					BuildCommand: jsii.String("buildCommand"),
//   					Port: jsii.String("port"),
//   					Runtime: jsii.String("runtime"),
//   					RuntimeEnvironmentSecrets: []interface{}{
//   						&KeyValuePairProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					RuntimeEnvironmentVariables: []interface{}{
//   						&KeyValuePairProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					StartCommand: jsii.String("startCommand"),
//   				},
//   				ConfigurationSource: jsii.String("configurationSource"),
//   			},
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   			SourceCodeVersion: &SourceCodeVersionProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			SourceDirectory: jsii.String("sourceDirectory"),
//   		},
//   		ImageRepository: &ImageRepositoryProperty{
//   			ImageConfiguration: &ImageConfigurationProperty{
//   				Port: jsii.String("port"),
//   				RuntimeEnvironmentSecrets: []interface{}{
//   					&KeyValuePairProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RuntimeEnvironmentVariables: []interface{}{
//   					&KeyValuePairProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				StartCommand: jsii.String("startCommand"),
//   			},
//   			ImageIdentifier: jsii.String("imageIdentifier"),
//   			ImageRepositoryType: jsii.String("imageRepositoryType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html
//
type CfnServicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServicePropsMixin
type jsiiProxy_CfnServicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServicePropsMixin) Props() *CfnServiceMixinProps {
	var returns *CfnServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppRunner::Service`.
func NewCfnServicePropsMixin(props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppRunner::Service`.
func NewCfnServicePropsMixin_Override(c CfnServicePropsMixin, props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServicePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

