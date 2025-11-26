package previewawselasticachemixins


// Specifies the authentication mode to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticationModeProperty := &AuthenticationModeProperty{
//   	Passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authenticationmode.html
//
type CfnUserPropsMixin_AuthenticationModeProperty struct {
	// Specifies the passwords to use for authentication if `Type` is set to `password` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authenticationmode.html#cfn-elasticache-user-authenticationmode-passwords
	//
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// Specifies the authentication type.
	//
	// Possible options are IAM authentication, password and no password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authenticationmode.html#cfn-elasticache-user-authenticationmode-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

