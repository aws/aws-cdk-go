package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspcs/previewawspcsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS PCS compute node group resource.
//
// For more information, see [Creating a compute node group in AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/working-with_cng_create.html) in the *AWS PCS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnComputeNodeGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnComputeNodeGroupPropsMixin(&CfnComputeNodeGroupMixinProps{
//   	AmiId: jsii.String("amiId"),
//   	ClusterId: jsii.String("clusterId"),
//   	CustomLaunchTemplate: &CustomLaunchTemplateProperty{
//   		TemplateId: jsii.String("templateId"),
//   		Version: jsii.String("version"),
//   	},
//   	IamInstanceProfileArn: jsii.String("iamInstanceProfileArn"),
//   	InstanceConfigs: []interface{}{
//   		&InstanceConfigProperty{
//   			InstanceType: jsii.String("instanceType"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PurchaseOption: jsii.String("purchaseOption"),
//   	ScalingConfiguration: &ScalingConfigurationProperty{
//   		MaxInstanceCount: jsii.Number(123),
//   		MinInstanceCount: jsii.Number(123),
//   	},
//   	SlurmConfiguration: &SlurmConfigurationProperty{
//   		SlurmCustomSettings: []interface{}{
//   			&SlurmCustomSettingProperty{
//   				ParameterName: jsii.String("parameterName"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   	},
//   	SpotOptions: &SpotOptionsProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-computenodegroup.html
//
type CfnComputeNodeGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnComputeNodeGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnComputeNodeGroupPropsMixin
type jsiiProxy_CfnComputeNodeGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnComputeNodeGroupPropsMixin) Props() *CfnComputeNodeGroupMixinProps {
	var returns *CfnComputeNodeGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeNodeGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PCS::ComputeNodeGroup`.
func NewCfnComputeNodeGroupPropsMixin(props *CfnComputeNodeGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnComputeNodeGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnComputeNodeGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComputeNodeGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PCS::ComputeNodeGroup`.
func NewCfnComputeNodeGroupPropsMixin_Override(c CfnComputeNodeGroupPropsMixin, props *CfnComputeNodeGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnComputeNodeGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComputeNodeGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComputeNodeGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComputeNodeGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnComputeNodeGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

