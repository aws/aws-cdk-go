package awsbedrockagentcore


// Configuration for EFS access point filesystem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsAccessPointConfigurationProperty := &EfsAccessPointConfigurationProperty{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-efsaccesspointconfiguration.html
//
type CfnRuntime_EfsAccessPointConfigurationProperty struct {
	// ARN of the EFS access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-efsaccesspointconfiguration.html#cfn-bedrockagentcore-runtime-efsaccesspointconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// Mount path for filesystem configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-efsaccesspointconfiguration.html#cfn-bedrockagentcore-runtime-efsaccesspointconfiguration-mountpath
	//
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
}

