package mixinsawsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsopsworks/mixinsawsopsworks/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnInstancePropsMixin(&CfnInstanceMixinProps{
//   	AgentVersion: jsii.String("agentVersion"),
//   	AmiId: jsii.String("amiId"),
//   	Architecture: jsii.String("architecture"),
//   	AutoScalingType: jsii.String("autoScalingType"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BlockDeviceMappings: []interface{}{
//   		&BlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			Ebs: &EbsBlockDeviceProperty{
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				SnapshotId: jsii.String("snapshotId"),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: jsii.String("noDevice"),
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	EbsOptimized: jsii.Boolean(false),
//   	ElasticIps: []*string{
//   		jsii.String("elasticIps"),
//   	},
//   	Hostname: jsii.String("hostname"),
//   	InstallUpdatesOnBoot: jsii.Boolean(false),
//   	InstanceType: jsii.String("instanceType"),
//   	LayerIds: []*string{
//   		jsii.String("layerIds"),
//   	},
//   	Os: jsii.String("os"),
//   	RootDeviceType: jsii.String("rootDeviceType"),
//   	SshKeyName: jsii.String("sshKeyName"),
//   	StackId: jsii.String("stackId"),
//   	SubnetId: jsii.String("subnetId"),
//   	Tenancy: jsii.String("tenancy"),
//   	TimeBasedAutoScaling: &TimeBasedAutoScalingProperty{
//   		Friday: map[string]*string{
//   			"fridayKey": jsii.String("friday"),
//   		},
//   		Monday: map[string]*string{
//   			"mondayKey": jsii.String("monday"),
//   		},
//   		Saturday: map[string]*string{
//   			"saturdayKey": jsii.String("saturday"),
//   		},
//   		Sunday: map[string]*string{
//   			"sundayKey": jsii.String("sunday"),
//   		},
//   		Thursday: map[string]*string{
//   			"thursdayKey": jsii.String("thursday"),
//   		},
//   		Tuesday: map[string]*string{
//   			"tuesdayKey": jsii.String("tuesday"),
//   		},
//   		Wednesday: map[string]*string{
//   			"wednesdayKey": jsii.String("wednesday"),
//   		},
//   	},
//   	VirtualizationType: jsii.String("virtualizationType"),
//   	Volumes: []*string{
//   		jsii.String("volumes"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
//
type CfnInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstancePropsMixin
type jsiiProxy_CfnInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInstancePropsMixin) Props() *CfnInstanceMixinProps {
	var returns *CfnInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpsWorks::Instance`.
func NewCfnInstancePropsMixin(props *CfnInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpsWorks::Instance`.
func NewCfnInstancePropsMixin_Override(c CfnInstancePropsMixin, props *CfnInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstancePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

