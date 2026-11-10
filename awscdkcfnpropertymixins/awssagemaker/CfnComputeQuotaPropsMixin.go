package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SageMaker::ComputeQuota.
//
// Defines a compute allocation rule for an EKS-orchestrated SageMaker HyperPod cluster's task governance. Available in 15 regions (us-east-1, us-east-2, us-west-1, us-west-2, ap-south-1, ap-southeast-1, ap-southeast-2, ap-southeast-3, ap-northeast-1, eu-central-1, eu-west-1, eu-west-2, eu-north-1, eu-south-2, sa-east-1).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnComputeQuotaPropsMixin := awscdkcfnpropertymixins.Aws_sagemaker.NewCfnComputeQuotaPropsMixin(&CfnComputeQuotaMixinProps{
//   	ActivationState: jsii.String("activationState"),
//   	ClusterArn: jsii.String("clusterArn"),
//   	ComputeQuotaConfig: &ComputeQuotaConfigProperty{
//   		ComputeQuotaResources: []interface{}{
//   			&ComputeQuotaResourceConfigProperty{
//   				AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   					Count: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   				Accelerators: jsii.Number(123),
//   				Count: jsii.Number(123),
//   				InstanceType: jsii.String("instanceType"),
//   				MemoryInGiB: jsii.Number(123),
//   				VCpu: jsii.Number(123),
//   			},
//   		},
//   		PreemptTeamTasks: jsii.String("preemptTeamTasks"),
//   		ResourceSharingConfig: &ResourceSharingConfigProperty{
//   			AbsoluteBorrowLimits: []interface{}{
//   				&ComputeQuotaResourceConfigProperty{
//   					AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   						Count: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   					Accelerators: jsii.Number(123),
//   					Count: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//   					MemoryInGiB: jsii.Number(123),
//   					VCpu: jsii.Number(123),
//   				},
//   			},
//   			BorrowLimit: jsii.Number(123),
//   			Strategy: jsii.String("strategy"),
//   		},
//   	},
//   	ComputeQuotaTarget: &ComputeQuotaTargetProperty{
//   		FairShareWeight: jsii.Number(123),
//   		TeamName: jsii.String("teamName"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html
//
type CfnComputeQuotaPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnComputeQuotaMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnComputeQuotaPropsMixin
type jsiiProxy_CfnComputeQuotaPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnComputeQuotaPropsMixin) Props() *CfnComputeQuotaMixinProps {
	var returns *CfnComputeQuotaMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeQuotaPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::ComputeQuota`.
func NewCfnComputeQuotaPropsMixin(props *CfnComputeQuotaMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnComputeQuotaPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnComputeQuotaPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComputeQuotaPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnComputeQuotaPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::ComputeQuota`.
func NewCfnComputeQuotaPropsMixin_Override(c CfnComputeQuotaPropsMixin, props *CfnComputeQuotaMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnComputeQuotaPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnComputeQuotaPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComputeQuotaPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnComputeQuotaPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComputeQuotaPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnComputeQuotaPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComputeQuotaPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnComputeQuotaPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

