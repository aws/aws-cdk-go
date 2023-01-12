package awsbatch


// Details about the compute resources managed by the compute environment.
//
// This parameter is required for managed compute environments. For more information, see [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeResourcesProperty := &computeResourcesProperty{
//   	maxvCpus: jsii.Number(123),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	bidPercentage: jsii.Number(123),
//   	desiredvCpus: jsii.Number(123),
//   	ec2Configuration: []interface{}{
//   		&ec2ConfigurationObjectProperty{
//   			imageType: jsii.String("imageType"),
//
//   			// the properties below are optional
//   			imageIdOverride: jsii.String("imageIdOverride"),
//   			imageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   		},
//   	},
//   	ec2KeyPair: jsii.String("ec2KeyPair"),
//   	imageId: jsii.String("imageId"),
//   	instanceRole: jsii.String("instanceRole"),
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   		version: jsii.String("version"),
//   	},
//   	minvCpus: jsii.Number(123),
//   	placementGroup: jsii.String("placementGroup"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	spotIamFleetRole: jsii.String("spotIamFleetRole"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	updateToLatestImageVersion: jsii.Boolean(false),
//   }
//
type CfnComputeEnvironment_ComputeResourcesProperty struct {
	// The maximum number of Amazon EC2 vCPUs that an environment can reach.
	//
	// > With both `BEST_FIT_PROGRESSIVE` and `SPOT_CAPACITY_OPTIMIZED` allocation strategies, AWS Batch might need to exceed `maxvCpus` to meet your capacity requirements. In this event, AWS Batch never exceeds `maxvCpus` by more than a single instance. That is, no more than a single instance from among those specified in your compute environment.
	MaxvCpus *float64 `field:"required" json:"maxvCpus" yaml:"maxvCpus"`
	// The VPC subnets where the compute resources are launched.
	//
	// Fargate compute resources can contain up to 16 subnets. For Fargate compute resources, providing an empty list will be handled as if this parameter wasn't specified and no change is made. For EC2 compute resources, providing an empty list removes the VPC subnets from the compute resource. For more information, see [VPCs and subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the *Amazon VPC User Guide* .
	//
	// When updating a compute environment, changing the VPC subnets requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// The type of compute environment: `EC2` , `SPOT` , `FARGATE` , or `FARGATE_SPOT` .
	//
	// For more information, see [Compute environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the *AWS Batch User Guide* .
	//
	// If you choose `SPOT` , you must also specify an Amazon EC2 Spot Fleet role with the `spotIamFleetRole` parameter. For more information, see [Amazon EC2 spot fleet role](https://docs.aws.amazon.com/batch/latest/userguide/spot_fleet_IAM_role.html) in the *AWS Batch User Guide* .
	//
	// When updating compute environment, changing the type of a compute environment requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// When updating the type of a compute environment, changing between `EC2` and `SPOT` or between `FARGATE` and `FARGATE_SPOT` will initiate an infrastructure update, but if you switch between `EC2` and `FARGATE` , AWS CloudFormation will create a new compute environment.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The allocation strategy to use for the compute resource if not enough instances of the best fitting instance type can be allocated.
	//
	// This might be because of availability of the instance type in the Region or [Amazon EC2 service limits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-resource-limits.html) . For more information, see [Allocation strategies](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) in the *AWS Batch User Guide* .
	//
	// When updating a compute environment, changing the allocation strategy requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* . `BEST_FIT` is not supported when updating a compute environment.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	//
	// - **BEST_FIT (default)** - AWS Batch selects an instance type that best fits the needs of the jobs with a preference for the lowest-cost instance type. If additional instances of the selected instance type aren't available, AWS Batch waits for the additional instances to be available. If there aren't enough instances available, or if the user is reaching [Amazon EC2 service limits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-resource-limits.html) then additional jobs aren't run until the currently running jobs have completed. This allocation strategy keeps costs lower but can limit scaling. If you are using Spot Fleets with `BEST_FIT` then the Spot Fleet IAM role must be specified.
	// - **BEST_FIT_PROGRESSIVE** - AWS Batch will select additional instance types that are large enough to meet the requirements of the jobs in the queue, with a preference for instance types with a lower cost per unit vCPU. If additional instances of the previously selected instance types aren't available, AWS Batch will select new instance types.
	// - **SPOT_CAPACITY_OPTIMIZED** - AWS Batch will select one or more instance types that are large enough to meet the requirements of the jobs in the queue, with a preference for instance types that are less likely to be interrupted. This allocation strategy is only available for Spot Instance compute resources.
	//
	// With both `BEST_FIT_PROGRESSIVE` and `SPOT_CAPACITY_OPTIMIZED` strategies, AWS Batch might need to go above `maxvCpus` to meet your capacity requirements. In this event, AWS Batch never exceeds `maxvCpus` by more than a single instance.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for that instance type before instances are launched.
	//
	// For example, if your maximum percentage is 20%, then the Spot price must be less than 20% of the current On-Demand price for that Amazon EC2 instance. You always pay the lowest (market) price and never more than your maximum percentage.
	//
	// When updating a compute environment, changing the bid percentage requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	BidPercentage *float64 `field:"optional" json:"bidPercentage" yaml:"bidPercentage"`
	// The desired number of Amazon EC2 vCPUS in the compute environment.
	//
	// AWS Batch modifies this value between the minimum and maximum values based on job queue demand.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	DesiredvCpus *float64 `field:"optional" json:"desiredvCpus" yaml:"desiredvCpus"`
	// Provides information used to select Amazon Machine Images (AMIs) for EC2 instances in the compute environment.
	//
	// If `Ec2Configuration` isn't specified, the default is `ECS_AL2` .
	//
	// When updating a compute environment, changing this setting requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* . To remove the EC2 configuration and any custom AMI ID specified in `imageIdOverride` , set this value to an empty string.
	//
	// One or two values can be provided.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	Ec2Configuration interface{} `field:"optional" json:"ec2Configuration" yaml:"ec2Configuration"`
	// The Amazon EC2 key pair that's used for instances launched in the compute environment.
	//
	// You can use this key pair to log in to your instances with SSH. To remove the Amazon EC2 key pair, set this value to an empty string.
	//
	// When updating a compute environment, changing the EC2 key pair requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	Ec2KeyPair *string `field:"optional" json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
	//
	// This parameter is overridden by the `imageIdOverride` member of the `Ec2Configuration` structure. To remove the custom AMI ID and use the default AMI ID, set this value to an empty string.
	//
	// When updating a compute environment, changing the AMI ID requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified. > The AMI that you choose for a compute environment must match the architecture of the instance types that you intend to use for that compute environment. For example, if your compute environment uses A1 instance types, the compute resource AMI that you choose must support ARM instances. Amazon ECS vends both x86 and ARM versions of the Amazon ECS-optimized Amazon Linux 2 AMI. For more information, see [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#ecs-optimized-ami-linux-variants.html) in the *Amazon Elastic Container Service Developer Guide* .
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The Amazon ECS instance profile applied to Amazon EC2 instances in a compute environment.
	//
	// You can specify the short name or full Amazon Resource Name (ARN) of an instance profile. For example, `*ecsInstanceRole*` or `arn:aws:iam:: *<aws_account_id>* :instance-profile/ *ecsInstanceRole*` . For more information, see [Amazon ECS instance role](https://docs.aws.amazon.com/batch/latest/userguide/instance_IAM_role.html) in the *AWS Batch User Guide* .
	//
	// When updating a compute environment, changing this setting requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// The instances types that can be launched.
	//
	// You can specify instance families to launch any instance type within those families (for example, `c5` or `p3` ), or you can specify specific sizes within a family (such as `c5.8xlarge` ). You can also choose `optimal` to select instance types (from the C4, M4, and R4 instance families) that match the demand of your job queues.
	//
	// When updating a compute environment, changing this setting requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified. > When you create a compute environment, the instance types that you select for the compute environment must share the same architecture. For example, you can't mix x86 and ARM instances in the same compute environment. > Currently, `optimal` uses instance types from the C4, M4, and R4 instance families. In Regions that don't have instance types from those instance families, instance types from the C5, M5. and R5 instance families are used.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The launch template to use for your compute resources.
	//
	// Any other compute resource parameters that you specify in a [CreateComputeEnvironment](https://docs.aws.amazon.com/batch/latest/APIReference/API_CreateComputeEnvironment.html) API operation override the same parameters in the launch template. You must specify either the launch template ID or launch template name in the request, but not both. For more information, see [Launch Template Support](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in the ** .
	//
	// > This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The minimum number of Amazon EC2 vCPUs that an environment should maintain (even if the compute environment is `DISABLED` ).
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	MinvCpus *float64 `field:"optional" json:"minvCpus" yaml:"minvCpus"`
	// The Amazon EC2 placement group to associate with your compute resources.
	//
	// If you intend to submit multi-node parallel jobs to your compute environment, you should consider creating a cluster placement group and associate it with your compute resources. This keeps your multi-node parallel job on a logical grouping of instances within a single Availability Zone with high network flow potential. For more information, see [Placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// When updating a compute environment, changing the placement group requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// The Amazon EC2 security groups associated with instances launched in the compute environment.
	//
	// This parameter is required for Fargate compute resources, where it can contain up to 5 security groups. For Fargate compute resources, providing an empty list is handled as if this parameter wasn't specified and no change is made. For EC2 compute resources, providing an empty list removes the security groups from the compute resource.
	//
	// When updating a compute environment, changing the EC2 security groups requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a `SPOT` compute environment.
	//
	// This role is required if the allocation strategy set to `BEST_FIT` or if the allocation strategy isn't specified. For more information, see [Amazon EC2 spot fleet role](https://docs.aws.amazon.com/batch/latest/userguide/spot_fleet_IAM_role.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified. > To tag your Spot Instances on creation, the Spot Fleet IAM role specified here must use the newer *AmazonEC2SpotFleetTaggingRole* managed policy. The previously recommended *AmazonEC2SpotFleetRole* managed policy doesn't have the required permissions to tag Spot Instances. For more information, see [Spot instances not tagged on creation](https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#spot-instance-no-tag) in the *AWS Batch User Guide* .
	SpotIamFleetRole *string `field:"optional" json:"spotIamFleetRole" yaml:"spotIamFleetRole"`
	// Key-value pair tags to be applied to EC2 resources that are launched in the compute environment.
	//
	// For AWS Batch , these take the form of "String1": "String2", where String1 is the tag key and String2 is the tag value−for example, `{ "Name": "Batch Instance - C4OnDemand" }` . This is helpful for recognizing your AWS Batch instances in the Amazon EC2 console. These tags aren't seen when using the AWS Batch `ListTagsForResource` API operation.
	//
	// When updating a compute environment, changing this setting requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether the AMI ID is updated to the latest one that's supported by AWS Batch when the compute environment has an infrastructure update.
	//
	// The default value is `false` .
	//
	// > If an AMI ID is specified in the `imageId` or `imageIdOverride` parameters or by the launch template specified in the `launchTemplate` parameter, this parameter is ignored. For more information on updating AMI IDs during an infrastructure update, see [Updating the AMI ID](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html#updating-compute-environments-ami) in the *AWS Batch User Guide* .
	//
	// When updating a compute environment, changing this setting requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	UpdateToLatestImageVersion interface{} `field:"optional" json:"updateToLatestImageVersion" yaml:"updateToLatestImageVersion"`
}

