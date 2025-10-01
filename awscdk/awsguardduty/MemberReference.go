package awsguardduty


// A reference to a Member resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberReference := &MemberReference{
//   	DetectorId: jsii.String("detectorId"),
//   	MemberId: jsii.String("memberId"),
//   }
//
type MemberReference struct {
	// The DetectorId of the Member resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The MemberId of the Member resource.
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
}

