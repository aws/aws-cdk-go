package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a managed node group for an Amazon EKS cluster.
//
// You can only create a node group for your cluster that is equal to the current Kubernetes version for the cluster. All node groups are created with the latest AMI release version for the respective minor Kubernetes version of the cluster, unless you deploy a custom AMI using a launch template.
//
// For later updates, you will only be able to update a node group using a launch template only if it was originally deployed with a launch template. Additionally, the launch template ID or name must match what was used when the node group was created. You can update the launch template version with necessary changes. For more information about using launch templates, see [Customizing managed nodes with launch templates](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) .
//
// An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and associated Amazon EC2 instances that are managed by AWS for an Amazon EKS cluster. For more information, see [Managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in the *Amazon EKS User Guide* .
//
// > Windows AMI types are only supported for commercial AWS Regions that support Windows on Amazon EKS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNodegroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnNodegroupPropsMixin(&CfnNodegroupMixinProps{
//   	AmiType: jsii.String("amiType"),
//   	CapacityType: jsii.String("capacityType"),
//   	ClusterName: jsii.String("clusterName"),
//   	DiskSize: jsii.Number(123),
//   	ForceUpdateEnabled: jsii.Boolean(false),
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	NodegroupName: jsii.String("nodegroupName"),
//   	NodeRepairConfig: &NodeRepairConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		MaxParallelNodesRepairedCount: jsii.Number(123),
//   		MaxParallelNodesRepairedPercentage: jsii.Number(123),
//   		MaxUnhealthyNodeThresholdCount: jsii.Number(123),
//   		MaxUnhealthyNodeThresholdPercentage: jsii.Number(123),
//   		NodeRepairConfigOverrides: []interface{}{
//   			&NodeRepairConfigOverridesProperty{
//   				MinRepairWaitTimeMins: jsii.Number(123),
//   				NodeMonitoringCondition: jsii.String("nodeMonitoringCondition"),
//   				NodeUnhealthyReason: jsii.String("nodeUnhealthyReason"),
//   				RepairAction: jsii.String("repairAction"),
//   			},
//   		},
//   	},
//   	NodeRole: jsii.String("nodeRole"),
//   	ReleaseVersion: jsii.String("releaseVersion"),
//   	RemoteAccess: &RemoteAccessProperty{
//   		Ec2SshKey: jsii.String("ec2SshKey"),
//   		SourceSecurityGroups: []*string{
//   			jsii.String("sourceSecurityGroups"),
//   		},
//   	},
//   	ScalingConfig: &ScalingConfigProperty{
//   		DesiredSize: jsii.Number(123),
//   		MaxSize: jsii.Number(123),
//   		MinSize: jsii.Number(123),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Taints: []interface{}{
//   		&TaintProperty{
//   			Effect: jsii.String("effect"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UpdateConfig: &UpdateConfigProperty{
//   		MaxUnavailable: jsii.Number(123),
//   		MaxUnavailablePercentage: jsii.Number(123),
//   		UpdateStrategy: jsii.String("updateStrategy"),
//   	},
//   	Version: jsii.String("version"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html
//
type CfnNodegroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNodegroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNodegroupPropsMixin
type jsiiProxy_CfnNodegroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNodegroupPropsMixin) Props() *CfnNodegroupMixinProps {
	var returns *CfnNodegroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EKS::Nodegroup`.
func NewCfnNodegroupPropsMixin(props *CfnNodegroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNodegroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNodegroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNodegroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EKS::Nodegroup`.
func NewCfnNodegroupPropsMixin_Override(c CfnNodegroupPropsMixin, props *CfnNodegroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNodegroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNodegroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNodegroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNodegroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNodegroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

