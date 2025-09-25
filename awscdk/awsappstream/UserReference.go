package awsappstream


// A reference to a User resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userReference := &UserReference{
//   	UserId: jsii.String("userId"),
//   }
//
type UserReference struct {
	// The Id of the User resource.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

