package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   managedInstancesProviderProperty := &ManagedInstancesProviderProperty{
//   	InfrastructureOptimization: &InfrastructureOptimizationProperty{
//   		ScaleInAfter: jsii.Number(123),
//   	},
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	InstanceLaunchTemplate: &InstanceLaunchTemplateProperty{
//   		CapacityOptionType: jsii.String("capacityOptionType"),
//   		CapacityReservations: &CapacityReservationsProperty{
//   			ReservationGroupArn: jsii.String("reservationGroupArn"),
//   			ReservationPreference: jsii.String("reservationPreference"),
//   		},
//   		Ec2InstanceProfileArn: jsii.String("ec2InstanceProfileArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html
//
type CfnComputeEnvironmentPropsMixin_ManagedInstancesProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-infrastructureoptimization
	//
	InfrastructureOptimization interface{} `field:"optional" json:"infrastructureOptimization" yaml:"infrastructureOptimization"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-infrastructurerolearn
	//
	InfrastructureRoleArn *string `field:"optional" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-instancelaunchtemplate
	//
	InstanceLaunchTemplate interface{} `field:"optional" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesprovider.html#cfn-batch-computeenvironment-managedinstancesprovider-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

