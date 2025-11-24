package mixinsawsecs


// The execute command and managed storage configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterConfigurationProperty := &ClusterConfigurationProperty{
//   	ExecuteCommandConfiguration: &ExecuteCommandConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		LogConfiguration: &ExecuteCommandLogConfigurationProperty{
//   			CloudWatchEncryptionEnabled: jsii.Boolean(false),
//   			CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3EncryptionEnabled: jsii.Boolean(false),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   		Logging: jsii.String("logging"),
//   	},
//   	ManagedStorageConfiguration: &ManagedStorageConfigurationProperty{
//   		FargateEphemeralStorageKmsKeyId: jsii.String("fargateEphemeralStorageKmsKeyId"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clusterconfiguration.html
//
type CfnClusterPropsMixin_ClusterConfigurationProperty struct {
	// The details of the execute command configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clusterconfiguration.html#cfn-ecs-cluster-clusterconfiguration-executecommandconfiguration
	//
	ExecuteCommandConfiguration interface{} `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// The details of the managed storage configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clusterconfiguration.html#cfn-ecs-cluster-clusterconfiguration-managedstorageconfiguration
	//
	ManagedStorageConfiguration interface{} `field:"optional" json:"managedStorageConfiguration" yaml:"managedStorageConfiguration"`
}

