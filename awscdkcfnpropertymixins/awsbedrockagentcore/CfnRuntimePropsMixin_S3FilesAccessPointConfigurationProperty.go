package awsbedrockagentcore


// Configuration for S3 Files access point filesystem.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3filesaccesspointconfiguration.html
//
type CfnRuntimePropsMixin_S3FilesAccessPointConfigurationProperty struct {
	// ARN of the S3 Files access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3filesaccesspointconfiguration.html#cfn-bedrockagentcore-runtime-s3filesaccesspointconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// Mount path for filesystem configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3filesaccesspointconfiguration.html#cfn-bedrockagentcore-runtime-s3filesaccesspointconfiguration-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

