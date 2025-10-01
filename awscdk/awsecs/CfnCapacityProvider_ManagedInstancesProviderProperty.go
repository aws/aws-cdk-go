package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstancesProviderProperty := &ManagedInstancesProviderProperty{
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	InstanceLaunchTemplate: &InstanceLaunchTemplateProperty{
//   		Ec2InstanceProfileArn: jsii.String("ec2InstanceProfileArn"),
//   		NetworkConfiguration: &ManagedInstancesNetworkConfigurationProperty{
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		InstanceRequirements: &InstanceRequirementsRequestProperty{
//   			MemoryMiB: &MemoryMiBRequestProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   			VCpuCount: &VCpuCountRangeRequestProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
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
//   		},
//   		Monitoring: jsii.String("monitoring"),
//   		StorageConfiguration: &ManagedInstancesStorageConfigurationProperty{
//   			StorageSizeGiB: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	PropagateTags: jsii.String("propagateTags"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html
//
type CfnCapacityProvider_ManagedInstancesProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-infrastructurerolearn
	//
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-instancelaunchtemplate
	//
	InstanceLaunchTemplate interface{} `field:"required" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

