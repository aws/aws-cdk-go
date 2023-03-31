package awsconnect


// Contains information about a phone number for a quick connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   phoneNumberQuickConnectConfigProperty := &phoneNumberQuickConnectConfigProperty{
//   	phoneNumber: jsii.String("phoneNumber"),
//   }
//
type CfnQuickConnect_PhoneNumberQuickConnectConfigProperty struct {
	// The phone number in E.164 format.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}

