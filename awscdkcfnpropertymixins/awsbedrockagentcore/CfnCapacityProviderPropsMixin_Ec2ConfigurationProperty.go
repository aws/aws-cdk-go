package awsbedrockagentcore


// Configuration for EC2-based capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ec2ConfigurationProperty := &Ec2ConfigurationProperty{
//   	LaunchTemplateSource: &LaunchTemplateSourceProperty{
//   		LaunchParameters: &LaunchParametersProperty{
//   			CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				CapacityReservationTarget: &CapacityReservationTargetProperty{
//   					CapacityReservationId: jsii.String("capacityReservationId"),
//   					CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   				},
//   			},
//   			EphemeralVolumes: []interface{}{
//   				&EphemeralBlockDeviceMappingProperty{
//   					DeviceName: jsii.String("deviceName"),
//   					Ebs: &EphemeralEBSVolumeConfigurationProperty{
//   						EbsCardIndex: jsii.Number(123),
//   						Encrypted: jsii.Boolean(false),
//   						Iops: jsii.Number(123),
//   						KmsKeyId: jsii.String("kmsKeyId"),
//   						SnapshotId: jsii.String("snapshotId"),
//   						Throughput: jsii.Number(123),
//   						VolumeInitializationRate: jsii.Number(123),
//   						VolumeSize: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//   					},
//   					VirtualName: jsii.String("virtualName"),
//   				},
//   			},
//   			InstanceProfileArn: jsii.String("instanceProfileArn"),
//   			InstanceRequirements: &InstanceRequirementsProperty{
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   			},
//   			LicenseSpecifications: []interface{}{
//   				&LicenseSpecificationProperty{
//   					LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   				},
//   			},
//   			Monitoring: jsii.String("monitoring"),
//   			OperatingSystem: jsii.String("operatingSystem"),
//   			PropagatedTags: map[string]*string{
//   				"propagatedTagsKey": jsii.String("propagatedTags"),
//   			},
//   			SshKeyName: jsii.String("sshKeyName"),
//   		},
//   	},
//   	LifecycleConfiguration: &InstanceLifecycleConfigurationProperty{
//   		IdleInstanceTimeout: jsii.Number(123),
//   		MaxLifetime: jsii.Number(123),
//   	},
//   	RootVolume: &RootVolumeConfigurationProperty{
//   		Encrypted: jsii.Boolean(false),
//   		FreeSpaceGiB: jsii.Number(123),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	Volumes: []interface{}{
//   		&VolumeConfigurationProperty{
//   			EbsConfiguration: &EbsVolumeConfigurationProperty{
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				Name: jsii.String("name"),
//   				SizeGiB: jsii.Number(123),
//   				SnapshotId: jsii.String("snapshotId"),
//   				Throughput: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   		},
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ec2configuration.html
//
type CfnCapacityProviderPropsMixin_Ec2ConfigurationProperty struct {
	// How the launch template is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ec2configuration.html#cfn-bedrockagentcore-capacityprovider-ec2configuration-launchtemplatesource
	//
	LaunchTemplateSource interface{} `field:"optional" json:"launchTemplateSource" yaml:"launchTemplateSource"`
	// Configuration for managing the lifecycle of instances in a capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ec2configuration.html#cfn-bedrockagentcore-capacityprovider-ec2configuration-lifecycleconfiguration
	//
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Customer-facing configuration for the (service-managed) root volume.
	//
	// The service provisions the root volume at its own AMI size estimate plus FreeSpaceGiB, and pins the visible free space to FreeSpaceGiB with a filler file, so the space you are guaranteed does not change as the underlying AMI grows. The device name and the delete-on-termination behavior are service-owned and are not configurable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ec2configuration.html#cfn-bedrockagentcore-capacityprovider-ec2configuration-rootvolume
	//
	RootVolume interface{} `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Named persistent EBS volumes for this capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ec2configuration.html#cfn-bedrockagentcore-capacityprovider-ec2configuration-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// VPC configuration for launching EC2 instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ec2configuration.html#cfn-bedrockagentcore-capacityprovider-ec2configuration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

