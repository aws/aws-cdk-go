package awsec2


// A reference to a VerifiedAccessGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessGroupReference := &VerifiedAccessGroupReference{
//   	VerifiedAccessGroupArn: jsii.String("verifiedAccessGroupArn"),
//   	VerifiedAccessGroupId: jsii.String("verifiedAccessGroupId"),
//   }
//
type VerifiedAccessGroupReference struct {
	// The ARN of the VerifiedAccessGroup resource.
	VerifiedAccessGroupArn *string `field:"required" json:"verifiedAccessGroupArn" yaml:"verifiedAccessGroupArn"`
	// The VerifiedAccessGroupId of the VerifiedAccessGroup resource.
	VerifiedAccessGroupId *string `field:"required" json:"verifiedAccessGroupId" yaml:"verifiedAccessGroupId"`
}

