package previewawsemrmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemr/previewawsemrmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use `InstanceGroupConfig` to define instance groups for an EMR cluster.
//
// A cluster can not use both instance groups and instance fleets. For more information, see [Create a Cluster with Instance Fleets or Uniform Instance Groups](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// > You can currently only add task instance groups to a cluster with this resource. If you use this resource, CloudFormation waits for the cluster launch to complete before adding the task instance group to the cluster. In order to add task instance groups to the cluster as part of the cluster launch and minimize delays in provisioning task nodes, use the `TaskInstanceGroups` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configurationProperty_ ConfigurationProperty
//
//   cfnInstanceGroupConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnInstanceGroupConfigPropsMixin(&CfnInstanceGroupConfigMixinProps{
//   	AutoScalingPolicy: &AutoScalingPolicyProperty{
//   		Constraints: &ScalingConstraintsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   		},
//   		Rules: []interface{}{
//   			&ScalingRuleProperty{
//   				Action: &ScalingActionProperty{
//   					Market: jsii.String("market"),
//   					SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   						AdjustmentType: jsii.String("adjustmentType"),
//   						CoolDown: jsii.Number(123),
//   						ScalingAdjustment: jsii.Number(123),
//   					},
//   				},
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Trigger: &ScalingTriggerProperty{
//   					CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   						ComparisonOperator: jsii.String("comparisonOperator"),
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						EvaluationPeriods: jsii.Number(123),
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   						Period: jsii.Number(123),
//   						Statistic: jsii.String("statistic"),
//   						Threshold: jsii.Number(123),
//   						Unit: jsii.String("unit"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	BidPrice: jsii.String("bidPrice"),
//   	Configurations: []interface{}{
//   		&ConfigurationProperty{
//   			Classification: jsii.String("classification"),
//   			ConfigurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			Configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   	CustomAmiId: jsii.String("customAmiId"),
//   	EbsConfiguration: &EbsConfigurationProperty{
//   		EbsBlockDeviceConfigs: []interface{}{
//   			&EbsBlockDeviceConfigProperty{
//   				VolumeSpecification: &VolumeSpecificationProperty{
//   					Iops: jsii.Number(123),
//   					SizeInGb: jsii.Number(123),
//   					Throughput: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	InstanceCount: jsii.Number(123),
//   	InstanceRole: jsii.String("instanceRole"),
//   	InstanceType: jsii.String("instanceType"),
//   	JobFlowId: jsii.String("jobFlowId"),
//   	Market: jsii.String("market"),
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html
//
type CfnInstanceGroupConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInstanceGroupConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstanceGroupConfigPropsMixin
type jsiiProxy_CfnInstanceGroupConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInstanceGroupConfigPropsMixin) Props() *CfnInstanceGroupConfigMixinProps {
	var returns *CfnInstanceGroupConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfigPropsMixin(props *CfnInstanceGroupConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInstanceGroupConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstanceGroupConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceGroupConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfigPropsMixin_Override(c CfnInstanceGroupConfigPropsMixin, props *CfnInstanceGroupConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInstanceGroupConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceGroupConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceGroupConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfigPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInstanceGroupConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

