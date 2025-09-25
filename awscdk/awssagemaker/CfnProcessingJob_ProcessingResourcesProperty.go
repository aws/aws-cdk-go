package awssagemaker


// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job.
//
// In distributed training, you specify more than one instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processingResourcesProperty := &ProcessingResourcesProperty{
//   	ClusterConfig: &ClusterConfigProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		VolumeSizeInGb: jsii.Number(123),
//
//   		// the properties below are optional
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingresources.html
//
type CfnProcessingJob_ProcessingResourcesProperty struct {
	// The configuration for the resources in a cluster used to run the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingresources.html#cfn-sagemaker-processingjob-processingresources-clusterconfig
	//
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

