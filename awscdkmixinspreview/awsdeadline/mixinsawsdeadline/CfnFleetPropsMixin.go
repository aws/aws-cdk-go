package mixinsawsdeadline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsdeadline/mixinsawsdeadline/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a fleet.
//
// Fleets gather information relating to compute, or capacity, for renders within your farms. You can choose to manage your own capacity or opt to have fleets fully managed by Deadline Cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnFleetPropsMixin(&CfnFleetMixinProps{
//   	Configuration: &FleetConfigurationProperty{
//   		CustomerManaged: &CustomerManagedFleetConfigurationProperty{
//   			Mode: jsii.String("mode"),
//   			StorageProfileId: jsii.String("storageProfileId"),
//   			TagPropagationMode: jsii.String("tagPropagationMode"),
//   			WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
//   				AcceleratorCount: &AcceleratorCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   				CustomAmounts: []interface{}{
//   					&FleetAmountCapabilityProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				CustomAttributes: []interface{}{
//   					&FleetAttributeCapabilityProperty{
//   						Name: jsii.String("name"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				MemoryMiB: &MemoryMiBRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OsFamily: jsii.String("osFamily"),
//   				VCpuCount: &VCpuCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ServiceManagedEc2: &ServiceManagedEc2FleetConfigurationProperty{
//   			InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   				AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   					Count: &AcceleratorCountRangeProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					Selections: []interface{}{
//   						&AcceleratorSelectionProperty{
//   							Name: jsii.String("name"),
//   							Runtime: jsii.String("runtime"),
//   						},
//   					},
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   				CustomAmounts: []interface{}{
//   					&FleetAmountCapabilityProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				CustomAttributes: []interface{}{
//   					&FleetAttributeCapabilityProperty{
//   						Name: jsii.String("name"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				ExcludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				MemoryMiB: &MemoryMiBRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OsFamily: jsii.String("osFamily"),
//   				RootEbsVolume: &Ec2EbsVolumeProperty{
//   					Iops: jsii.Number(123),
//   					SizeGiB: jsii.Number(123),
//   					ThroughputMiB: jsii.Number(123),
//   				},
//   				VCpuCount: &VCpuCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   			InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   				Type: jsii.String("type"),
//   			},
//   			StorageProfileId: jsii.String("storageProfileId"),
//   			VpcConfiguration: &VpcConfigurationProperty{
//   				ResourceConfigurationArns: []*string{
//   					jsii.String("resourceConfigurationArns"),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	HostConfiguration: &HostConfigurationProperty{
//   		ScriptBody: jsii.String("scriptBody"),
//   		ScriptTimeoutSeconds: jsii.Number(123),
//   	},
//   	MaxWorkerCount: jsii.Number(123),
//   	MinWorkerCount: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html
//
type CfnFleetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFleetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFleetPropsMixin
type jsiiProxy_CfnFleetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFleetPropsMixin) Props() *CfnFleetMixinProps {
	var returns *CfnFleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Deadline::Fleet`.
func NewCfnFleetPropsMixin(props *CfnFleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Deadline::Fleet`.
func NewCfnFleetPropsMixin_Override(c CfnFleetPropsMixin, props *CfnFleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFleetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

