package interfacesawsmacie


// A reference to a Member resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberReference := &MemberReference{
//   	MemberArn: jsii.String("memberArn"),
//   }
//
type MemberReference struct {
	// The Arn of the Member resource.
	MemberArn *string `field:"required" json:"memberArn" yaml:"memberArn"`
}

