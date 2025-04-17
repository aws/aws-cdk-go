package awsiot


// Structure that contains payloadVersion and targetArn.
//
// Provisioning hooks can be used when fleet provisioning to validate device parameters before allowing the device to be provisioned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisioningHookProperty := &ProvisioningHookProperty{
//   	PayloadVersion: jsii.String("payloadVersion"),
//   	TargetArn: jsii.String("targetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-provisioningtemplate-provisioninghook.html
//
type CfnProvisioningTemplate_ProvisioningHookProperty struct {
	// The payload that was sent to the target function.
	//
	// The valid payload is `"2020-04-01"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-provisioningtemplate-provisioninghook.html#cfn-iot-provisioningtemplate-provisioninghook-payloadversion
	//
	PayloadVersion *string `field:"optional" json:"payloadVersion" yaml:"payloadVersion"`
	// The ARN of the target function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-provisioningtemplate-provisioninghook.html#cfn-iot-provisioningtemplate-provisioninghook-targetarn
	//
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

