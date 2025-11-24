package mixinsawsecs


// The configuration for a Amazon ECS Managed Instances provider.
//
// Amazon ECS uses this configuration to automatically launch, manage, and terminate Amazon EC2 instances on your behalf. Managed instances provide access to the full range of Amazon EC2 instance types and features while offloading infrastructure management to AWS .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedInstancesProviderProperty := &ManagedInstancesProviderProperty{
//   	InfrastructureOptimization: &InfrastructureOptimizationProperty{
//   		ScaleInAfter: jsii.Number(123),
//   	},
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	InstanceLaunchTemplate: &InstanceLaunchTemplateProperty{
//   		Ec2InstanceProfileArn: jsii.String("ec2InstanceProfileArn"),
//   		InstanceRequirements: &InstanceRequirementsRequestProperty{
//   			AcceleratorCount: &AcceleratorCountRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorManufacturers: []*string{
//   				jsii.String("acceleratorManufacturers"),
//   			},
//   			AcceleratorNames: []*string{
//   				jsii.String("acceleratorNames"),
//   			},
//   			AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorTypes: []*string{
//   				jsii.String("acceleratorTypes"),
//   			},
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   			BareMetal: jsii.String("bareMetal"),
//   			BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			BurstablePerformance: jsii.String("burstablePerformance"),
//   			CpuManufacturers: []*string{
//   				jsii.String("cpuManufacturers"),
//   			},
//   			ExcludedInstanceTypes: []*string{
//   				jsii.String("excludedInstanceTypes"),
//   			},
//   			InstanceGenerations: []*string{
//   				jsii.String("instanceGenerations"),
//   			},
//   			LocalStorage: jsii.String("localStorage"),
//   			LocalStorageTypes: []*string{
//   				jsii.String("localStorageTypes"),
//   			},
//   			MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   			MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			MemoryMiB: &MemoryMiBRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   			RequireHibernateSupport: jsii.Boolean(false),
//   			SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   			TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			VCpuCount: &VCpuCountRangeRequestProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   		},
//   		Monitoring: jsii.String("monitoring"),
//   		NetworkConfiguration: &ManagedInstancesNetworkConfigurationProperty{
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   		StorageConfiguration: &ManagedInstancesStorageConfigurationProperty{
//   			StorageSizeGiB: jsii.Number(123),
//   		},
//   	},
//   	PropagateTags: jsii.String("propagateTags"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html
//
type CfnCapacityProviderPropsMixin_ManagedInstancesProviderProperty struct {
	// Defines how Amazon ECS Managed Instances optimizes the infrastructure in your capacity provider.
	//
	// Configure it to turn on or off the infrastructure optimization in your capacity provider, and to control the idle EC2 instances optimization delay.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-infrastructureoptimization
	//
	InfrastructureOptimization interface{} `field:"optional" json:"infrastructureOptimization" yaml:"infrastructureOptimization"`
	// The Amazon Resource Name (ARN) of the infrastructure role that Amazon ECS assumes to manage instances.
	//
	// This role must include permissions for Amazon EC2 instance lifecycle management, networking, and any additional AWS services required for your workloads.
	//
	// For more information, see [Amazon ECS infrastructure IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/infrastructure_IAM_role.html) in the *Amazon ECS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-infrastructurerolearn
	//
	InfrastructureRoleArn *string `field:"optional" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// The launch template that defines how Amazon ECS launches Amazon ECS Managed Instances.
	//
	// This includes the instance profile for your tasks, network and storage configuration, and instance requirements that determine which Amazon EC2 instance types can be used.
	//
	// For more information, see [Store instance launch parameters in Amazon EC2 launch templates](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-instancelaunchtemplate
	//
	InstanceLaunchTemplate interface{} `field:"optional" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// Determines whether tags from the capacity provider are automatically applied to Amazon ECS Managed Instances.
	//
	// This helps with cost allocation and resource management by ensuring consistent tagging across your infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

