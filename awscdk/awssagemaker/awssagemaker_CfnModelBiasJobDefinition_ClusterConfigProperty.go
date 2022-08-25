package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigProperty := &clusterConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	volumeSizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnModelBiasJobDefinition_ClusterConfigProperty struct {
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.InstanceCount`.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.InstanceType`.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.VolumeSizeInGB`.
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.VolumeKmsKeyId`.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

