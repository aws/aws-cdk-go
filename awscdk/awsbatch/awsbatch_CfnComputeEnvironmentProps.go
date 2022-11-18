package awsbatch


// Properties for defining a `CfnComputeEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComputeEnvironmentProps := &cfnComputeEnvironmentProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	computeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	computeResources: &computeResourcesProperty{
//   		maxvCpus: jsii.Number(123),
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		bidPercentage: jsii.Number(123),
//   		desiredvCpus: jsii.Number(123),
//   		ec2Configuration: []interface{}{
//   			&ec2ConfigurationObjectProperty{
//   				imageType: jsii.String("imageType"),
//
//   				// the properties below are optional
//   				imageIdOverride: jsii.String("imageIdOverride"),
//   				imageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   			},
//   		},
//   		ec2KeyPair: jsii.String("ec2KeyPair"),
//   		imageId: jsii.String("imageId"),
//   		instanceRole: jsii.String("instanceRole"),
//   		instanceTypes: []*string{
//   			jsii.String("instanceTypes"),
//   		},
//   		launchTemplate: &launchTemplateSpecificationProperty{
//   			launchTemplateId: jsii.String("launchTemplateId"),
//   			launchTemplateName: jsii.String("launchTemplateName"),
//   			version: jsii.String("version"),
//   		},
//   		minvCpus: jsii.Number(123),
//   		placementGroup: jsii.String("placementGroup"),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		spotIamFleetRole: jsii.String("spotIamFleetRole"),
//   		tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		updateToLatestImageVersion: jsii.Boolean(false),
//   	},
//   	eksConfiguration: &eksConfigurationProperty{
//   		eksClusterArn: jsii.String("eksClusterArn"),
//   		kubernetesNamespace: jsii.String("kubernetesNamespace"),
//   	},
//   	replaceComputeEnvironment: jsii.Boolean(false),
//   	serviceRole: jsii.String("serviceRole"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	unmanagedvCpus: jsii.Number(123),
//   	updatePolicy: &updatePolicyProperty{
//   		jobExecutionTimeoutMinutes: jsii.Number(123),
//   		terminateJobsOnUpdate: jsii.Boolean(false),
//   	},
//   }
//
type CfnComputeEnvironmentProps struct {
	// The type of the compute environment: `MANAGED` or `UNMANAGED` .
	//
	// For more information, see [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the *AWS Batch User Guide* .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name for your compute environment.
	//
	// It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	ComputeEnvironmentName *string `field:"optional" json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// The ComputeResources property type specifies details of the compute resources managed by the compute environment.
	//
	// This parameter is required for managed compute environments. For more information, see [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the ** .
	ComputeResources interface{} `field:"optional" json:"computeResources" yaml:"computeResources"`
	// `AWS::Batch::ComputeEnvironment.EksConfiguration`.
	EksConfiguration interface{} `field:"optional" json:"eksConfiguration" yaml:"eksConfiguration"`
	// Specifies whether the compute environment should be replaced if an update is made that requires replacing the instances in the compute environment.
	//
	// The default value is `true` . To enable more properties to be updated, set this property to `false` . When changing the value of this property to `false` , no other properties should be changed at the same time. If other properties are changed at the same time, and the change needs to be rolled back but it can't, it's possible for the stack to go into the `UPDATE_ROLLBACK_FAILED` state. You can't update a stack that is in the `UPDATE_ROLLBACK_FAILED` state. However, if you can continue to roll it back, you can return the stack to its original settings and then try to update it again. For more information, see [Continue rolling back an update](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html) in the *AWS CloudFormation User Guide* .
	//
	// The properties that can't be changed without replacing the compute environment are in the [`ComputeResources`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html) property type: [`AllocationStrategy`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-allocationstrategy) , [`BidPercentage`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-bidpercentage) , [`Ec2Configuration`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-ec2configuration) , [`Ec2KeyPair`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-ec2keypair) , [`Ec2KeyPair`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-ec2keypair) , [`ImageId`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-imageid) , [`InstanceRole`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-instancerole) , [`InstanceTypes`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-instancetypes) , [`LaunchTemplate`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-launchtemplate) , [`MaxvCpus`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-maxvcpus) , [`MinvCpus`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-minvcpus) , [`PlacementGroup`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-placementgroup) , [`SecurityGroupIds`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-securitygroupids) , [`Subnets`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-subnets) , [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-tags) , [`Type`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-type) , and [`UpdateToLatestImageVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) .
	ReplaceComputeEnvironment interface{} `field:"optional" json:"replaceComputeEnvironment" yaml:"replaceComputeEnvironment"`
	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
	//
	// For more information, see [AWS Batch service IAM role](https://docs.aws.amazon.com/batch/latest/userguide/service_IAM_role.html) in the *AWS Batch User Guide* .
	//
	// > If your account already created the AWS Batch service-linked role, that role is used by default for your compute environment unless you specify a different role here. If the AWS Batch service-linked role doesn't exist in your account, and no role is specified here, the service attempts to create the AWS Batch service-linked role in your account.
	//
	// If your specified role has a path other than `/` , then you must specify either the full role ARN (recommended) or prefix the role name with the path. For example, if a role with the name `bar` has a path of `/foo/` then you would specify `/foo/bar` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	//
	// > Depending on how you created your AWS Batch service role, its ARN might contain the `service-role` path prefix. When you only specify the name of the service role, AWS Batch assumes that your ARN doesn't use the `service-role` path prefix. Because of this, we recommend that you specify the full ARN of your service role when you create compute environments.
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The state of the compute environment.
	//
	// If the state is `ENABLED` , then the compute environment accepts jobs from a queue and can scale out automatically based on queues.
	//
	// If the state is `ENABLED` , then the AWS Batch scheduler can attempt to place jobs from an associated job queue on the compute resources within the environment. If the compute environment is managed, then it can scale its instances out or in automatically, based on the job queue demand.
	//
	// If the state is `DISABLED` , then the AWS Batch scheduler doesn't attempt to place jobs within the environment. Jobs in a `STARTING` or `RUNNING` state continue to progress normally. Managed compute environments in the `DISABLED` state don't scale out. However, they scale in to `minvCpus` value after instances become idle.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags applied to the compute environment.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The maximum number of vCPUs for an unmanaged compute environment.
	//
	// This parameter is only used for fair share scheduling to reserve vCPU capacity for new share identifiers. If this parameter isn't provided for a fair share job queue, no vCPU capacity is reserved.
	//
	// > This parameter is only supported when the `type` parameter is set to `UNMANAGED` .
	UnmanagedvCpus *float64 `field:"optional" json:"unmanagedvCpus" yaml:"unmanagedvCpus"`
	// Specifies the infrastructure update policy for the compute environment.
	//
	// For more information about infrastructure updates, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	UpdatePolicy interface{} `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
}

