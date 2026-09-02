package interfacesawssupportauthz


// A reference to a SupportPermit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   supportPermitReference := &SupportPermitReference{
//   	SupportPermitArn: jsii.String("supportPermitArn"),
//   }
//
type SupportPermitReference struct {
	// The Arn of the SupportPermit resource.
	SupportPermitArn *string `field:"required" json:"supportPermitArn" yaml:"supportPermitArn"`
}

