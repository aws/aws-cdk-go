package mixinsawsecs


// The configuration for a volume specified in the task definition as a volume that is configured at launch time.
//
// Currently, the only supported volume type is an Amazon EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceVolumeConfigurationProperty := &ServiceVolumeConfigurationProperty{
//   	ManagedEbsVolume: &ServiceManagedEBSVolumeConfigurationProperty{
//   		Encrypted: jsii.Boolean(false),
//   		FilesystemType: jsii.String("filesystemType"),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		RoleArn: jsii.String("roleArn"),
//   		SizeInGiB: jsii.Number(123),
//   		SnapshotId: jsii.String("snapshotId"),
//   		TagSpecifications: []interface{}{
//   			&EBSTagSpecificationProperty{
//   				PropagateTags: jsii.String("propagateTags"),
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Throughput: jsii.Number(123),
//   		VolumeInitializationRate: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-servicevolumeconfiguration.html
//
type CfnServicePropsMixin_ServiceVolumeConfigurationProperty struct {
	// The configuration for the Amazon EBS volume that Amazon ECS creates and manages on your behalf.
	//
	// These settings are used to create each Amazon EBS volume, with one volume created for each task in the service. The Amazon EBS volumes are visible in your account in the Amazon EC2 console once they are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-servicevolumeconfiguration.html#cfn-ecs-service-servicevolumeconfiguration-managedebsvolume
	//
	ManagedEbsVolume interface{} `field:"optional" json:"managedEbsVolume" yaml:"managedEbsVolume"`
	// The name of the volume.
	//
	// This value must match the volume name from the `Volume` object in the task definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-servicevolumeconfiguration.html#cfn-ecs-service-servicevolumeconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

