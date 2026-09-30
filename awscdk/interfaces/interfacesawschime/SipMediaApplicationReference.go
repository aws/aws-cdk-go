package interfacesawschime


// A reference to a SipMediaApplication resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sipMediaApplicationReference := &SipMediaApplicationReference{
//   	SipMediaApplicationArn: jsii.String("sipMediaApplicationArn"),
//   }
//
type SipMediaApplicationReference struct {
	// The SipMediaApplicationArn of the SipMediaApplication resource.
	SipMediaApplicationArn *string `field:"required" json:"sipMediaApplicationArn" yaml:"sipMediaApplicationArn"`
}

