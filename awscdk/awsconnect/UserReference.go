package awsconnect


// A reference to a User resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userReference := &UserReference{
//   	UserArn: jsii.String("userArn"),
//   }
//
type UserReference struct {
	// The UserArn of the User resource.
	UserArn *string `field:"required" json:"userArn" yaml:"userArn"`
}

