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
//   provisioningHookProperty := &provisioningHookProperty{
//   	payloadVersion: jsii.String("payloadVersion"),
//   	targetArn: jsii.String("targetArn"),
//   }
//
type CfnProvisioningTemplate_ProvisioningHookProperty struct {
	// The payload that was sent to the target function.
	//
	// The valid payload is `"2020-04-01"` .
	PayloadVersion *string `field:"optional" json:"payloadVersion" yaml:"payloadVersion"`
	// The ARN of the target function.
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

