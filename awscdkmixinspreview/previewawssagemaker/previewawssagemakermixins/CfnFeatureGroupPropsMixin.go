package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a new `FeatureGroup` .
//
// A `FeatureGroup` is a group of `Features` defined in the `FeatureStore` to describe a `Record` .
//
// The `FeatureGroup` defines the schema and features contained in the FeatureGroup. A `FeatureGroup` definition is composed of a list of `Features` , a `RecordIdentifierFeatureName` , an `EventTimeFeatureName` and configurations for its `OnlineStore` and `OfflineStore` . Check [AWS service quotas](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) to see the `FeatureGroup` s quota for your AWS account.
//
// > You must include at least one of `OnlineStoreConfig` and `OfflineStoreConfig` to create a `FeatureGroup` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var offlineStoreConfig interface{}
//   var onlineStoreConfig interface{}
//
//   cfnFeatureGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnFeatureGroupPropsMixin(&CfnFeatureGroupMixinProps{
//   	Description: jsii.String("description"),
//   	EventTimeFeatureName: jsii.String("eventTimeFeatureName"),
//   	FeatureDefinitions: []interface{}{
//   		&FeatureDefinitionProperty{
//   			FeatureName: jsii.String("featureName"),
//   			FeatureType: jsii.String("featureType"),
//   		},
//   	},
//   	FeatureGroupName: jsii.String("featureGroupName"),
//   	OfflineStoreConfig: offlineStoreConfig,
//   	OnlineStoreConfig: onlineStoreConfig,
//   	RecordIdentifierFeatureName: jsii.String("recordIdentifierFeatureName"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThroughputConfig: &ThroughputConfigProperty{
//   		ProvisionedReadCapacityUnits: jsii.Number(123),
//   		ProvisionedWriteCapacityUnits: jsii.Number(123),
//   		ThroughputMode: jsii.String("throughputMode"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html
//
type CfnFeatureGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFeatureGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFeatureGroupPropsMixin
type jsiiProxy_CfnFeatureGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFeatureGroupPropsMixin) Props() *CfnFeatureGroupMixinProps {
	var returns *CfnFeatureGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::FeatureGroup`.
func NewCfnFeatureGroupPropsMixin(props *CfnFeatureGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFeatureGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFeatureGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFeatureGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::FeatureGroup`.
func NewCfnFeatureGroupPropsMixin_Override(c CfnFeatureGroupPropsMixin, props *CfnFeatureGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFeatureGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFeatureGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeatureGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeatureGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFeatureGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

