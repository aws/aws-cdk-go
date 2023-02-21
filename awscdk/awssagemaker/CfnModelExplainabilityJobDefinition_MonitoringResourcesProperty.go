package awssagemaker


// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringResourcesProperty := &MonitoringResourcesProperty{
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
type CfnModelExplainabilityJobDefinition_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

