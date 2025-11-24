package mixinsawsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsbatch/mixinsawsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Batch::ComputeEnvironment` resource defines your AWS Batch compute environment.
//
// You can define `MANAGED` or `UNMANAGED` compute environments. `MANAGED` compute environments can use Amazon EC2 or AWS Fargate resources. `UNMANAGED` compute environments can only use EC2 resources. For more information, see [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the ** .
//
// In a managed compute environment, AWS Batch manages the capacity and instance types of the compute resources within the environment. This is based on the compute resource specification that you define or the [launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) that you specify when you create the compute environment. You can choose either to use EC2 On-Demand Instances and EC2 Spot Instances, or to use Fargate and Fargate Spot capacity in your managed compute environment. You can optionally set a maximum price so that Spot Instances only launch when the Spot Instance price is below a specified percentage of the On-Demand price.
//
// > Multi-node parallel jobs are not supported on Spot Instances.
//
// In an unmanaged compute environment, you can manage your own EC2 compute resources and have a lot of flexibility with how you configure your compute resources. For example, you can use custom AMI. However, you need to verify that your AMI meets the Amazon ECS container instance AMI specification. For more information, see [container instance AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html) in the *Amazon Elastic Container Service Developer Guide* . After you have created your unmanaged compute environment, you can use the [DescribeComputeEnvironments](https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeComputeEnvironments.html) operation to find the Amazon ECS cluster that is associated with it. Then, manually launch your container instances into that Amazon ECS cluster. For more information, see [Launching an Amazon ECS container instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > To create a compute environment that uses EKS resources, the caller must have permissions to call `eks:DescribeCluster` . > AWS Batch doesn't upgrade the AMIs in a compute environment after it's created except under specific conditions. For example, it doesn't automatically update the AMIs when a newer version of the Amazon ECS optimized AMI is available. Therefore, you're responsible for the management of the guest operating system (including updates and security patches) and any additional application software or utilities that you install on the compute resources. There are two ways to use a new AMI for your AWS Batch jobs. The original method is to complete these steps:
// >
// > - Create a new compute environment with the new AMI.
// > - Add the compute environment to an existing job queue.
// > - Remove the earlier compute environment from your job queue.
// > - Delete the earlier compute environment.
// >
// > In April 2022, AWS Batch added enhanced support for updating compute environments. For example, the `UpdateComputeEnvironent` API lets you use the `ReplaceComputeEnvironment` property to dynamically update compute environment parameters such as the launch template or instance type without replacement. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
// >
// > To use the enhanced updating of compute environments to update AMIs, follow these rules:
// >
// > - Either do not set the [ServiceRole](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-servicerole) property or set it to the *AWSServiceRoleForBatch* service-linked role.
// > - Set the [AllocationStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-allocationstrategy) property to `BEST_FIT_PROGRESSIVE` or `SPOT_CAPACITY_OPTIMIZED` .
// > - Set the [ReplaceComputeEnvironment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-replacecomputeenvironment) property to `false` .
// >
// > > Set the `ReplaceComputeEnvironment` property to `true` if the compute environment uses the `BEST_FIT` allocation strategy. > If the `ReplaceComputeEnvironment` property is set to `false` , you might receive an error message when you update the CFN template for a compute environment. This issue occurs if the updated `desiredvcpus` value is less than the current `desiredvcpus` value. As a workaround, delete the `desiredvcpus` value from the updated template or use the `minvcpus` property to manage the number of vCPUs. For information, see [Error message when you update the `DesiredvCpus` setting](https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#error-desired-vcpus-update) .
// > - Set the [UpdateToLatestImageVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) property to `true` . This property is used when you update a compute environment. The [UpdateToLatestImageVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) property is ignored when you create a compute environment.
// > - Either do not specify an image ID in [ImageId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-imageid) or [ImageIdOverride](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html#cfn-batch-computeenvironment-ec2configurationobject-imageidoverride) properties, or in the launch template identified by the [Launch Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-launchtemplate) property. In that case AWS Batch will select the latest Amazon ECS optimized AMI supported by AWS Batch at the time the infrastructure update is initiated. Alternatively you can specify the AMI ID in the `ImageId` or `ImageIdOverride` properties, or the launch template identified by the `LaunchTemplate` properties. Changing any of these properties will trigger an infrastructure update.
// >
// > If these rules are followed, any update that triggers an infrastructure update will cause the AMI ID to be re-selected. If the [Version](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-version) property of the [LaunchTemplateSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html) is set to `$Latest` or `$Default` , the latest or default version of the launch template will be evaluated up at the time of the infrastructure update, even if the `LaunchTemplateSpecification` was not updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnComputeEnvironmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnComputeEnvironmentPropsMixin(&CfnComputeEnvironmentMixinProps{
//   	ComputeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	ComputeResources: &ComputeResourcesProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		BidPercentage: jsii.Number(123),
//   		DesiredvCpus: jsii.Number(123),
//   		Ec2Configuration: []interface{}{
//   			&Ec2ConfigurationObjectProperty{
//   				ImageIdOverride: jsii.String("imageIdOverride"),
//   				ImageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   				ImageType: jsii.String("imageType"),
//   			},
//   		},
//   		Ec2KeyPair: jsii.String("ec2KeyPair"),
//   		ImageId: jsii.String("imageId"),
//   		InstanceRole: jsii.String("instanceRole"),
//   		InstanceTypes: []*string{
//   			jsii.String("instanceTypes"),
//   		},
//   		LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   			LaunchTemplateId: jsii.String("launchTemplateId"),
//   			LaunchTemplateName: jsii.String("launchTemplateName"),
//   			Overrides: []interface{}{
//   				&LaunchTemplateSpecificationOverrideProperty{
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					LaunchTemplateName: jsii.String("launchTemplateName"),
//   					TargetInstanceTypes: []*string{
//   						jsii.String("targetInstanceTypes"),
//   					},
//   					UserdataType: jsii.String("userdataType"),
//   					Version: jsii.String("version"),
//   				},
//   			},
//   			UserdataType: jsii.String("userdataType"),
//   			Version: jsii.String("version"),
//   		},
//   		MaxvCpus: jsii.Number(123),
//   		MinvCpus: jsii.Number(123),
//   		PlacementGroup: jsii.String("placementGroup"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SpotIamFleetRole: jsii.String("spotIamFleetRole"),
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		Type: jsii.String("type"),
//   		UpdateToLatestImageVersion: jsii.Boolean(false),
//   	},
//   	Context: jsii.String("context"),
//   	EksConfiguration: &EksConfigurationProperty{
//   		EksClusterArn: jsii.String("eksClusterArn"),
//   		KubernetesNamespace: jsii.String("kubernetesNamespace"),
//   	},
//   	ReplaceComputeEnvironment: jsii.Boolean(false),
//   	ServiceRole: jsii.String("serviceRole"),
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Type: jsii.String("type"),
//   	UnmanagedvCpus: jsii.Number(123),
//   	UpdatePolicy: &UpdatePolicyProperty{
//   		JobExecutionTimeoutMinutes: jsii.Number(123),
//   		TerminateJobsOnUpdate: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html
//
type CfnComputeEnvironmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnComputeEnvironmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnComputeEnvironmentPropsMixin
type jsiiProxy_CfnComputeEnvironmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnComputeEnvironmentPropsMixin) Props() *CfnComputeEnvironmentMixinProps {
	var returns *CfnComputeEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Batch::ComputeEnvironment`.
func NewCfnComputeEnvironmentPropsMixin(props *CfnComputeEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnComputeEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnComputeEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComputeEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Batch::ComputeEnvironment`.
func NewCfnComputeEnvironmentPropsMixin_Override(c CfnComputeEnvironmentPropsMixin, props *CfnComputeEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnComputeEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComputeEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComputeEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComputeEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnComputeEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

