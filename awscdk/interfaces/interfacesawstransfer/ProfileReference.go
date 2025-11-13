package interfacesawstransfer


// A reference to a Profile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileReference := &ProfileReference{
//   	ProfileArn: jsii.String("profileArn"),
//   	ProfileId: jsii.String("profileId"),
//   }
//
type ProfileReference struct {
	// The ARN of the Profile resource.
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
	// The ProfileId of the Profile resource.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
}

