package awsmemorydb


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
	// Passwords used for this user account.
	//
	// You can create up to two passwords for each user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-user-authenticationmode.html#cfn-memorydb-user-authenticationmode-passwords
	//
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// Type of authentication strategy for this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-user-authenticationmode.html#cfn-memorydb-user-authenticationmode-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

