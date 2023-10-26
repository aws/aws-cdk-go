package awsmemorydb


// Denotes the user's authentication properties, such as whether it requires a password to authenticate.
//
// Used in output responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationModeProperty := &AuthenticationModeProperty{
//   	Passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-user-authenticationmode.html
//
type CfnUser_AuthenticationModeProperty struct {
	// The password(s) used for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-user-authenticationmode.html#cfn-memorydb-user-authenticationmode-passwords
	//
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// Indicates whether the user requires a password to authenticate.
	//
	// All newly-created users require a password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-user-authenticationmode.html#cfn-memorydb-user-authenticationmode-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

