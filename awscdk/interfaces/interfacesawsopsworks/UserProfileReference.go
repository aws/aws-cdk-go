package interfacesawsopsworks


// A reference to a UserProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userProfileReference := &UserProfileReference{
//   	UserProfileId: jsii.String("userProfileId"),
//   }
//
type UserProfileReference struct {
	// The Id of the UserProfile resource.
	UserProfileId *string `field:"required" json:"userProfileId" yaml:"userProfileId"`
}

