package awsobservabilityadmin

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::ObservabilityAdmin::DatasetIntegration resource represents a CloudWatch dataset integration for an account in a region.
//
// The integration enables telemetry forwarding into the caller's CloudWatch dataset via a customer-provided IAM role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDatasetIntegrationPropsMixin := awscdkcfnpropertymixins.Aws_observabilityadmin.NewCfnDatasetIntegrationPropsMixin(&CfnDatasetIntegrationMixinProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-datasetintegration.html
//
type CfnDatasetIntegrationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDatasetIntegrationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatasetIntegrationPropsMixin
type jsiiProxy_CfnDatasetIntegrationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDatasetIntegrationPropsMixin) Props() *CfnDatasetIntegrationMixinProps {
	var returns *CfnDatasetIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetIntegrationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::DatasetIntegration`.
func NewCfnDatasetIntegrationPropsMixin(props *CfnDatasetIntegrationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDatasetIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatasetIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatasetIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnDatasetIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::DatasetIntegration`.
func NewCfnDatasetIntegrationPropsMixin_Override(c CfnDatasetIntegrationPropsMixin, props *CfnDatasetIntegrationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnDatasetIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDatasetIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatasetIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnDatasetIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatasetIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnDatasetIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatasetIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDatasetIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

