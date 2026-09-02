package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   instanceLaunchTemplateProperty := &InstanceLaunchTemplateProperty{
//   	CapacityOptionType: jsii.String("capacityOptionType"),
//   	CapacityReservations: &CapacityReservationsProperty{
//   		ReservationGroupArn: jsii.String("reservationGroupArn"),
//   		ReservationPreference: jsii.String("reservationPreference"),
//   	},
//   	Ec2InstanceProfileArn: jsii.String("ec2InstanceProfileArn"),
//   	FipsEnabled: jsii.Boolean(false),
//   	InstanceMetadataTagsPropagation: jsii.Boolean(false),
//   	InstanceRequirements: &InstanceRequirementsProperty{
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   	},
//   	LocalStorageConfiguration: &ManagedInstancesLocalStorageConfigurationProperty{
//   		UseLocalStorage: jsii.Boolean(false),
//   	},
//   	Monitoring: jsii.String("monitoring"),
//   	NetworkConfiguration: &ManagedInstancesNetworkConfigurationProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   	StorageConfiguration: &ManagedInstancesStorageConfigurationProperty{
//   		StorageSizeGiB: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html
//
type CfnComputeEnvironmentPropsMixin_InstanceLaunchTemplateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-capacityoptiontype
	//
	CapacityOptionType *string `field:"optional" json:"capacityOptionType" yaml:"capacityOptionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-capacityreservations
	//
	CapacityReservations interface{} `field:"optional" json:"capacityReservations" yaml:"capacityReservations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-ec2instanceprofilearn
	//
	Ec2InstanceProfileArn *string `field:"optional" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-fipsenabled
	//
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-instancemetadatatagspropagation
	//
	InstanceMetadataTagsPropagation interface{} `field:"optional" json:"instanceMetadataTagsPropagation" yaml:"instanceMetadataTagsPropagation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-localstorageconfiguration
	//
	LocalStorageConfiguration interface{} `field:"optional" json:"localStorageConfiguration" yaml:"localStorageConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-monitoring
	//
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-instancelaunchtemplate.html#cfn-batch-computeenvironment-instancelaunchtemplate-storageconfiguration
	//
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

