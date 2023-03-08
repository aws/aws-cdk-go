package awsecs


// The execute command configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigurationProperty := &clusterConfigurationProperty{
//   	executeCommandConfiguration: &executeCommandConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		logConfiguration: &executeCommandLogConfigurationProperty{
//   			cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   			cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			s3BucketName: jsii.String("s3BucketName"),
//   			s3EncryptionEnabled: jsii.Boolean(false),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   		logging: jsii.String("logging"),
//   	},
//   }
//
type CfnCluster_ClusterConfigurationProperty struct {
	// The details of the execute command configuration.
	ExecuteCommandConfiguration interface{} `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}

