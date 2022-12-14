package awselasticache


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationModeProperty := &authenticationModeProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   }
//
type CfnUser_AuthenticationModeProperty struct {
	// `CfnUser.AuthenticationModeProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnUser.AuthenticationModeProperty.Passwords`.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

