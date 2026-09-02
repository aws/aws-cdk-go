package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCapacityProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityProviderProps := &CfnCapacityProviderProps{
//   	ComputeConfiguration: &ComputeConfigurationProperty{
//   		Ec2Configuration: &Ec2ConfigurationProperty{
//   			LaunchTemplateSource: &LaunchTemplateSourceProperty{
//   				LaunchParameters: &LaunchParametersProperty{
//   					InstanceRequirements: &InstanceRequirementsProperty{
//   						AllowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   					},
//   					OperatingSystem: jsii.String("operatingSystem"),
//
//   					// the properties below are optional
//   					CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   						CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   						CapacityReservationTarget: &CapacityReservationTargetProperty{
//   							CapacityReservationId: jsii.String("capacityReservationId"),
//   							CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   						},
//   					},
//   					EphemeralVolumes: []interface{}{
//   						&EphemeralBlockDeviceMappingProperty{
//   							DeviceName: jsii.String("deviceName"),
//   							Ebs: &EphemeralEBSVolumeConfigurationProperty{
//   								EbsCardIndex: jsii.Number(123),
//   								Encrypted: jsii.Boolean(false),
//   								Iops: jsii.Number(123),
//   								KmsKeyId: jsii.String("kmsKeyId"),
//   								SnapshotId: jsii.String("snapshotId"),
//   								Throughput: jsii.Number(123),
//   								VolumeInitializationRate: jsii.Number(123),
//   								VolumeSize: jsii.Number(123),
//   								VolumeType: jsii.String("volumeType"),
//   							},
//   							VirtualName: jsii.String("virtualName"),
//   						},
//   					},
//   					InstanceProfileArn: jsii.String("instanceProfileArn"),
//   					LicenseSpecifications: []interface{}{
//   						&LicenseSpecificationProperty{
//   							LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   						},
//   					},
//   					Monitoring: jsii.String("monitoring"),
//   					PropagatedTags: map[string]*string{
//   						"propagatedTagsKey": jsii.String("propagatedTags"),
//   					},
//   					SshKeyName: jsii.String("sshKeyName"),
//   				},
//   			},
//   			VpcConfiguration: &VpcConfigurationProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			LifecycleConfiguration: &InstanceLifecycleConfigurationProperty{
//   				IdleInstanceTimeout: jsii.Number(123),
//   				MaxLifetime: jsii.Number(123),
//   			},
//   			RootVolume: &RootVolumeConfigurationProperty{
//   				Encrypted: jsii.Boolean(false),
//   				FreeSpaceGiB: jsii.Number(123),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				Throughput: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			Volumes: []interface{}{
//   				&VolumeConfigurationProperty{
//   					EbsConfiguration: &EbsVolumeConfigurationProperty{
//   						Name: jsii.String("name"),
//   						SizeGiB: jsii.Number(123),
//
//   						// the properties below are optional
//   						Encrypted: jsii.Boolean(false),
//   						Iops: jsii.Number(123),
//   						KmsKeyId: jsii.String("kmsKeyId"),
//   						SnapshotId: jsii.String("snapshotId"),
//   						Throughput: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PermissionsConfiguration: &PermissionsConfigurationProperty{
//   		CapacityProviderOperatorRoleArn: jsii.String("capacityProviderOperatorRoleArn"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html
//
type CfnCapacityProviderProps struct {
	// The capacity configuration for the capacity provider.
	//
	// Defines the compute resources for this capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-computeconfiguration
	//
	ComputeConfiguration interface{} `field:"required" json:"computeConfiguration" yaml:"computeConfiguration"`
	// The name of the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for permissions associated with a capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-permissionsconfiguration
	//
	PermissionsConfiguration interface{} `field:"required" json:"permissionsConfiguration" yaml:"permissionsConfiguration"`
	// An optional description of the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

