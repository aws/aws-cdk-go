package awsbedrockagentcore


// Configuration for an Amazon EFS access point to mount into the AgentCore Runtime.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-efsaccesspointconfiguration.html
//
type CfnHarness_EfsAccessPointConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-efsaccesspointconfiguration.html#cfn-bedrockagentcore-harness-efsaccesspointconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-efsaccesspointconfiguration.html#cfn-bedrockagentcore-harness-efsaccesspointconfiguration-mountpath
	//
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
}

