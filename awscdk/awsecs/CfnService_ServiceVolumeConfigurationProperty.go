package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceVolumeConfigurationProperty := &ServiceVolumeConfigurationProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ManagedEbsVolume: &ServiceManagedEBSVolumeConfigurationProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		Encrypted: jsii.Boolean(false),
//   		FilesystemType: jsii.String("filesystemType"),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		SizeInGiB: jsii.Number(123),
//   		SnapshotId: jsii.String("snapshotId"),
//   		TagSpecifications: []interface{}{
//   			&EBSTagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//
//   				// the properties below are optional
//   				PropagateTags: jsii.String("propagateTags"),
//   				Tags: []cfnTag{
//   					&cfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-servicevolumeconfiguration.html
//
type CfnService_ServiceVolumeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-servicevolumeconfiguration.html#cfn-ecs-service-servicevolumeconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-servicevolumeconfiguration.html#cfn-ecs-service-servicevolumeconfiguration-managedebsvolume
	//
	ManagedEbsVolume interface{} `field:"optional" json:"managedEbsVolume" yaml:"managedEbsVolume"`
}

