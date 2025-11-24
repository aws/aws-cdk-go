package mixinsawssagemaker


// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job.
//
// In distributed training, you specify more than one instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   processingResourcesProperty := &ProcessingResourcesProperty{
//   	ClusterConfig: &ClusterConfigProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		VolumeSizeInGb: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingresources.html
//
type CfnProcessingJobPropsMixin_ProcessingResourcesProperty struct {
	// The configuration for the resources in a cluster used to run the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingresources.html#cfn-sagemaker-processingjob-processingresources-clusterconfig
	//
	ClusterConfig interface{} `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
}

