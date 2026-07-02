package awsbedrockagentcore


// Configuration for an Amazon S3 Files access point to mount into the AgentCore Runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FilesAccessPointConfigurationProperty := &S3FilesAccessPointConfigurationProperty{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-s3filesaccesspointconfiguration.html
//
type CfnHarness_S3FilesAccessPointConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-s3filesaccesspointconfiguration.html#cfn-bedrockagentcore-harness-s3filesaccesspointconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-s3filesaccesspointconfiguration.html#cfn-bedrockagentcore-harness-s3filesaccesspointconfiguration-mountpath
	//
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
}

