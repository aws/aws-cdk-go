package awsec2


// A reference to a VerifiedAccessInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessInstanceReference := &VerifiedAccessInstanceReference{
//   	VerifiedAccessInstanceId: jsii.String("verifiedAccessInstanceId"),
//   }
//
type VerifiedAccessInstanceReference struct {
	// The VerifiedAccessInstanceId of the VerifiedAccessInstance resource.
	VerifiedAccessInstanceId *string `field:"required" json:"verifiedAccessInstanceId" yaml:"verifiedAccessInstanceId"`
}

