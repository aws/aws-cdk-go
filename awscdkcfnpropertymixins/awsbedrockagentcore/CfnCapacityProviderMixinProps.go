package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCapacityProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCapacityProviderMixinProps := &CfnCapacityProviderMixinProps{
//   	ComputeConfiguration: &ComputeConfigurationProperty{
//   		Ec2Configuration: &Ec2ConfigurationProperty{
//   			LaunchTemplateSource: &LaunchTemplateSourceProperty{
//   				LaunchParameters: &LaunchParametersProperty{
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
//   					InstanceRequirements: &InstanceRequirementsProperty{
//   						AllowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   					},
//   					LicenseSpecifications: []interface{}{
//   						&LicenseSpecificationProperty{
//   							LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   						},
//   					},
//   					Monitoring: jsii.String("monitoring"),
//   					OperatingSystem: jsii.String("operatingSystem"),
//   					PropagatedTags: map[string]*string{
//   						"propagatedTagsKey": jsii.String("propagatedTags"),
//   					},
//   					SshKeyName: jsii.String("sshKeyName"),
//   				},
//   			},
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
//   						Encrypted: jsii.Boolean(false),
//   						Iops: jsii.Number(123),
//   						KmsKeyId: jsii.String("kmsKeyId"),
//   						Name: jsii.String("name"),
//   						SizeGiB: jsii.Number(123),
//   						SnapshotId: jsii.String("snapshotId"),
//   						Throughput: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//   					},
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
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PermissionsConfiguration: &PermissionsConfigurationProperty{
//   		CapacityProviderOperatorRoleArn: jsii.String("capacityProviderOperatorRoleArn"),
//   	},
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
type CfnCapacityProviderMixinProps struct {
	// The capacity configuration for the capacity provider.
	//
	// Defines the compute resources for this capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-computeconfiguration
	//
	ComputeConfiguration interface{} `field:"optional" json:"computeConfiguration" yaml:"computeConfiguration"`
	// An optional description of the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configuration for permissions associated with a capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-permissionsconfiguration
	//
	PermissionsConfiguration interface{} `field:"optional" json:"permissionsConfiguration" yaml:"permissionsConfiguration"`
	// An array of key-value pairs to apply to the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-capacityprovider.html#cfn-bedrockagentcore-capacityprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

