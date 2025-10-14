package awssmsvoice


// A reference to a PhoneNumber resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   phoneNumberReference := &PhoneNumberReference{
//   	PhoneNumberArn: jsii.String("phoneNumberArn"),
//   	PhoneNumberId: jsii.String("phoneNumberId"),
//   }
//
type PhoneNumberReference struct {
	// The ARN of the PhoneNumber resource.
	PhoneNumberArn *string `field:"required" json:"phoneNumberArn" yaml:"phoneNumberArn"`
	// The PhoneNumberId of the PhoneNumber resource.
	PhoneNumberId *string `field:"required" json:"phoneNumberId" yaml:"phoneNumberId"`
}

