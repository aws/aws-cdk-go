package awsmanagedblockchain


// A reference to a Member resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberReference := &MemberReference{
//   	MemberId: jsii.String("memberId"),
//   }
//
type MemberReference struct {
	// The MemberId of the Member resource.
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
}

