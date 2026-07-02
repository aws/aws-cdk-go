package awsbedrockagentcore


// Configuration for an Amazon S3 Files access point to mount into the AgentCore Runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3FilesAccessPointConfigurationProperty := &S3FilesAccessPointConfigurationProperty{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-s3filesaccesspointconfiguration.html
//
type CfnHarnessPropsMixin_S3FilesAccessPointConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-s3filesaccesspointconfiguration.html#cfn-bedrockagentcore-harness-s3filesaccesspointconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-s3filesaccesspointconfiguration.html#cfn-bedrockagentcore-harness-s3filesaccesspointconfiguration-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

