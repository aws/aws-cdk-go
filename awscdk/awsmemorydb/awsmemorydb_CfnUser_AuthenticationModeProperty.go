package awsmemorydb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationModeProperty := &authenticationModeProperty{
//   	passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnUser_AuthenticationModeProperty struct {
	// `CfnUser.AuthenticationModeProperty.Passwords`.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// `CfnUser.AuthenticationModeProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

