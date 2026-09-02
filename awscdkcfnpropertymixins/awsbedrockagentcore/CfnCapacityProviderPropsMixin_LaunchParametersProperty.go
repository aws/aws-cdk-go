package awsbedrockagentcore


// Parameters for launching EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   launchParametersProperty := &LaunchParametersProperty{
//   	CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   		CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   		CapacityReservationTarget: &CapacityReservationTargetProperty{
//   			CapacityReservationId: jsii.String("capacityReservationId"),
//   			CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   		},
//   	},
//   	EphemeralVolumes: []interface{}{
//   		&EphemeralBlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			Ebs: &EphemeralEBSVolumeConfigurationProperty{
//   				EbsCardIndex: jsii.Number(123),
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				SnapshotId: jsii.String("snapshotId"),
//   				Throughput: jsii.Number(123),
//   				VolumeInitializationRate: jsii.Number(123),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   	InstanceRequirements: &InstanceRequirementsProperty{
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   	},
//   	LicenseSpecifications: []interface{}{
//   		&LicenseSpecificationProperty{
//   			LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   		},
//   	},
//   	Monitoring: jsii.String("monitoring"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	PropagatedTags: map[string]*string{
//   		"propagatedTagsKey": jsii.String("propagatedTags"),
//   	},
//   	SshKeyName: jsii.String("sshKeyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html
//
type CfnCapacityProviderPropsMixin_LaunchParametersProperty struct {
	// The Capacity Reservation targeting option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-capacityreservationspecification
	//
	CapacityReservationSpecification interface{} `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// The block device mapping for ephemeral (instance store) volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-ephemeralvolumes
	//
	EphemeralVolumes interface{} `field:"optional" json:"ephemeralVolumes" yaml:"ephemeralVolumes"`
	// The ARN of the IAM instance profile to associate with launched instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-instanceprofilearn
	//
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Requirements for EC2 instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The license configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-licensespecifications
	//
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// The monitoring level for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-monitoring
	//
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The operating system and CPU architecture for the instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-operatingsystem
	//
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// Tags to apply to all EC2 resources (instances, volumes, and network interfaces) created by this capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-propagatedtags
	//
	PropagatedTags interface{} `field:"optional" json:"propagatedTags" yaml:"propagatedTags"`
	// The name of the SSH key pair to configure on instances for SSH connectivity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchparameters.html#cfn-bedrockagentcore-capacityprovider-launchparameters-sshkeyname
	//
	SshKeyName *string `field:"optional" json:"sshKeyName" yaml:"sshKeyName"`
}

