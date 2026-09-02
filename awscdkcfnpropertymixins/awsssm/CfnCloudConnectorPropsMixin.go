package awsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SSM::CloudConnector.
//
// Enables AWS Systems Manager to manage resources in external cloud providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCloudConnectorPropsMixin := awscdkcfnpropertymixins.Aws_ssm.NewCfnCloudConnectorPropsMixin(&CfnCloudConnectorMixinProps{
//   	ConfigConnectorArn: jsii.String("configConnectorArn"),
//   	Configuration: &CloudConnectorConfigurationProperty{
//   		AzureConfiguration: &AzureConfigurationProperty{
//   			ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   			ApplicationId: jsii.String("applicationId"),
//   			Targets: &ConfigurationTargetsProperty{
//   				Subscriptions: []interface{}{
//   					&AzureSubscriptionProperty{
//   						DisplayName: jsii.String("displayName"),
//   						Id: jsii.String("id"),
//   					},
//   				},
//   			},
//   			TenantDisplayName: jsii.String("tenantDisplayName"),
//   			TenantId: jsii.String("tenantId"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html
//
type CfnCloudConnectorPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCloudConnectorMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudConnectorPropsMixin
type jsiiProxy_CfnCloudConnectorPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCloudConnectorPropsMixin) Props() *CfnCloudConnectorMixinProps {
	var returns *CfnCloudConnectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudConnectorPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::CloudConnector`.
func NewCfnCloudConnectorPropsMixin(props *CfnCloudConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCloudConnectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudConnectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudConnectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnCloudConnectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::CloudConnector`.
func NewCfnCloudConnectorPropsMixin_Override(c CfnCloudConnectorPropsMixin, props *CfnCloudConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnCloudConnectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCloudConnectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudConnectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnCloudConnectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudConnectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnCloudConnectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudConnectorPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCloudConnectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

