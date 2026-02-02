package previewawsemrmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemr/previewawsemrmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use `InstanceFleetConfig` to define instance fleets for an EMR cluster.
//
// A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. > You can currently only add a task instance fleet to a cluster with this resource. If you use this resource, CloudFormation waits for the cluster launch to complete before adding the task instance fleet to the cluster. In order to add a task instance fleet to the cluster as part of the cluster launch and minimize delays in provisioning task nodes, use the `TaskInstanceFleets` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configurationProperty_ ConfigurationProperty
//
//   cfnInstanceFleetConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnInstanceFleetConfigPropsMixin(&CfnInstanceFleetConfigMixinProps{
//   	ClusterId: jsii.String("clusterId"),
//   	InstanceFleetType: jsii.String("instanceFleetType"),
//   	InstanceTypeConfigs: []interface{}{
//   		&InstanceTypeConfigProperty{
//   			BidPrice: jsii.String("bidPrice"),
//   			BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			Configurations: []interface{}{
//   				&ConfigurationProperty{
//   					Classification: jsii.String("classification"),
//   					ConfigurationProperties: map[string]*string{
//   						"configurationPropertiesKey": jsii.String("configurationProperties"),
//   					},
//   					Configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   				},
//   			},
//   			CustomAmiId: jsii.String("customAmiId"),
//   			EbsConfiguration: &EbsConfigurationProperty{
//   				EbsBlockDeviceConfigs: []interface{}{
//   					&EbsBlockDeviceConfigProperty{
//   						VolumeSpecification: &VolumeSpecificationProperty{
//   							Iops: jsii.Number(123),
//   							SizeInGb: jsii.Number(123),
//   							Throughput: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			Priority: jsii.Number(123),
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   		OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   				UsageStrategy: jsii.String("usageStrategy"),
//   			},
//   		},
//   		SpotSpecification: &SpotProvisioningSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			BlockDurationMinutes: jsii.Number(123),
//   			TimeoutAction: jsii.String("timeoutAction"),
//   			TimeoutDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	ResizeSpecifications: &InstanceFleetResizingSpecificationsProperty{
//   		OnDemandResizeSpecification: &OnDemandResizingSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   				UsageStrategy: jsii.String("usageStrategy"),
//   			},
//   			TimeoutDurationMinutes: jsii.Number(123),
//   		},
//   		SpotResizeSpecification: &SpotResizingSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			TimeoutDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	TargetOnDemandCapacity: jsii.Number(123),
//   	TargetSpotCapacity: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html
//
type CfnInstanceFleetConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInstanceFleetConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstanceFleetConfigPropsMixin
type jsiiProxy_CfnInstanceFleetConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInstanceFleetConfigPropsMixin) Props() *CfnInstanceFleetConfigMixinProps {
	var returns *CfnInstanceFleetConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMR::InstanceFleetConfig`.
func NewCfnInstanceFleetConfigPropsMixin(props *CfnInstanceFleetConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInstanceFleetConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstanceFleetConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceFleetConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMR::InstanceFleetConfig`.
func NewCfnInstanceFleetConfigPropsMixin_Override(c CfnInstanceFleetConfigPropsMixin, props *CfnInstanceFleetConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInstanceFleetConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceFleetConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceFleetConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

