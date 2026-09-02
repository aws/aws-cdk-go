package awsbatch


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
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		CapacityOptionType: jsii.String("capacityOptionType"),
//   		CapacityReservations: &CapacityReservationsProperty{
//   			ReservationGroupArn: jsii.String("reservationGroupArn"),
//   			ReservationPreference: jsii.String("reservationPreference"),
//   		},
//   		FipsEnabled: jsii.Boolean(false),
//   		InstanceMetadataTagsPropagation: jsii.Boolean(false),
//   		InstanceRequirements: &InstanceRequirementsProperty{
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   		},
//   		LocalStorageConfiguration: &ManagedInstancesLocalStorageConfigurationProperty{
//   			UseLocalStorage: jsii.Boolean(false),
//   		},
//   		Monitoring: jsii.String("monitoring"),
//   		StorageConfiguration: &ManagedInstancesStorageConfigurationProperty{
//   			StorageSizeGiB: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	InfrastructureOptimization: &InfrastructureOptimizationProperty{
//   		ScaleInAfter: jsii.Number(123),
//   	},
//   	PropagateTags: jsii.String("propagateTags"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html
//
type CfnComputeEnvironment_ManagedInstancesProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-infrastructurerolearn
	//
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-instancelaunchtemplate
	//
	InstanceLaunchTemplate interface{} `field:"required" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-infrastructureoptimization
	//
	InfrastructureOptimization interface{} `field:"optional" json:"infrastructureOptimization" yaml:"infrastructureOptimization"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

