package awselasticache


// Specifies the authentication mode to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationModeProperty := &AuthenticationModeProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authenticationmode.html
//
type CfnUser_AuthenticationModeProperty struct {
	// Specifies the authentication type.
	//
	// Possible options are IAM authentication, password and no password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authenticationmode.html#cfn-elasticache-user-authenticationmode-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the passwords to use for authentication if `Type` is set to `password` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authenticationmode.html#cfn-elasticache-user-authenticationmode-passwords
	//
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

