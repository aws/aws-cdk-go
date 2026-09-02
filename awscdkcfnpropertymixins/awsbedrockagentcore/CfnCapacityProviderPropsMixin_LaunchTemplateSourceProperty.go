package awsbedrockagentcore


// How the launch template is specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   launchTemplateSourceProperty := &LaunchTemplateSourceProperty{
//   	LaunchParameters: &LaunchParametersProperty{
//   		CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   			CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   			CapacityReservationTarget: &CapacityReservationTargetProperty{
//   				CapacityReservationId: jsii.String("capacityReservationId"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   			},
//   		},
//   		EphemeralVolumes: []interface{}{
//   			&EphemeralBlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EphemeralEBSVolumeConfigurationProperty{
//   					EbsCardIndex: jsii.Number(123),
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					SnapshotId: jsii.String("snapshotId"),
//   					Throughput: jsii.Number(123),
//   					VolumeInitializationRate: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		InstanceProfileArn: jsii.String("instanceProfileArn"),
//   		InstanceRequirements: &InstanceRequirementsProperty{
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   		},
//   		LicenseSpecifications: []interface{}{
//   			&LicenseSpecificationProperty{
//   				LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   			},
//   		},
//   		Monitoring: jsii.String("monitoring"),
//   		OperatingSystem: jsii.String("operatingSystem"),
//   		PropagatedTags: map[string]*string{
//   			"propagatedTagsKey": jsii.String("propagatedTags"),
//   		},
//   		SshKeyName: jsii.String("sshKeyName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchtemplatesource.html
//
type CfnCapacityProviderPropsMixin_LaunchTemplateSourceProperty struct {
	// Parameters for launching EC2 instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-launchtemplatesource.html#cfn-bedrockagentcore-capacityprovider-launchtemplatesource-launchparameters
	//
	LaunchParameters interface{} `field:"optional" json:"launchParameters" yaml:"launchParameters"`
}

