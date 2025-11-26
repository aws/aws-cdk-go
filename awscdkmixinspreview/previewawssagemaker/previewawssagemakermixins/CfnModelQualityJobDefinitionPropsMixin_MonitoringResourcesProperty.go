package previewawssagemakermixins


// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringResourcesProperty := &MonitoringResourcesProperty{
//   	ClusterConfig: &ClusterConfigProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		VolumeSizeInGb: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-monitoringresources.html
//
type CfnModelQualityJobDefinitionPropsMixin_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-monitoringresources.html#cfn-sagemaker-modelqualityjobdefinition-monitoringresources-clusterconfig
	//
	ClusterConfig interface{} `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
}

