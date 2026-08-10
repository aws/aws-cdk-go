package interfacesawswellarchitected


// A reference to a Profile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileReference := &ProfileReference{
//   	ProfileArn: jsii.String("profileArn"),
//   }
//
type ProfileReference struct {
	// The ProfileArn of the Profile resource.
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
}

