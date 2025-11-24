package mixinsawslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awslightsail/mixinsawslightsail/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lightsail::Instance` resource specifies an Amazon Lightsail instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnInstancePropsMixin(&CfnInstanceMixinProps{
//   	AddOns: []interface{}{
//   		&AddOnProperty{
//   			AddOnType: jsii.String("addOnType"),
//   			AutoSnapshotAddOnRequest: &AutoSnapshotAddOnProperty{
//   				SnapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BlueprintId: jsii.String("blueprintId"),
//   	BundleId: jsii.String("bundleId"),
//   	Hardware: &HardwareProperty{
//   		CpuCount: jsii.Number(123),
//   		Disks: []interface{}{
//   			&DiskProperty{
//   				AttachedTo: jsii.String("attachedTo"),
//   				AttachmentState: jsii.String("attachmentState"),
//   				DiskName: jsii.String("diskName"),
//   				Iops: jsii.Number(123),
//   				IsSystemDisk: jsii.Boolean(false),
//   				Path: jsii.String("path"),
//   				SizeInGb: jsii.String("sizeInGb"),
//   			},
//   		},
//   		RamSizeInGb: jsii.Number(123),
//   	},
//   	InstanceName: jsii.String("instanceName"),
//   	KeyPairName: jsii.String("keyPairName"),
//   	Location: &LocationProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		RegionName: jsii.String("regionName"),
//   	},
//   	Networking: &NetworkingProperty{
//   		MonthlyTransfer: &MonthlyTransferProperty{
//   			GbPerMonthAllocated: jsii.String("gbPerMonthAllocated"),
//   		},
//   		Ports: []interface{}{
//   			&PortProperty{
//   				AccessDirection: jsii.String("accessDirection"),
//   				AccessFrom: jsii.String("accessFrom"),
//   				AccessType: jsii.String("accessType"),
//   				CidrListAliases: []*string{
//   					jsii.String("cidrListAliases"),
//   				},
//   				Cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   				CommonName: jsii.String("commonName"),
//   				FromPort: jsii.Number(123),
//   				Ipv6Cidrs: []*string{
//   					jsii.String("ipv6Cidrs"),
//   				},
//   				Protocol: jsii.String("protocol"),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   	},
//   	State: &StateProperty{
//   		Code: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserData: jsii.String("userData"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html
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


// Create a mixin to apply properties to `AWS::Lightsail::Instance`.
func NewCfnInstancePropsMixin(props *CfnInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::Instance`.
func NewCfnInstancePropsMixin_Override(c CfnInstancePropsMixin, props *CfnInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin",
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

