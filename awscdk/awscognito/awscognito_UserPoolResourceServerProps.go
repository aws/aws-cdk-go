package awscognito


// Properties for the UserPoolResourceServer construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourceServerScope resourceServerScope
//   var userPool userPool
//
//   userPoolResourceServerProps := &userPoolResourceServerProps{
//   	identifier: jsii.String("identifier"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	scopes: []*resourceServerScope{
//   		resourceServerScope,
//   	},
//   	userPoolResourceServerName: jsii.String("userPoolResourceServerName"),
//   }
//
type UserPoolResourceServerProps struct {
	// A unique resource server identifier for the resource server.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Oauth scopes.
	Scopes *[]ResourceServerScope `field:"optional" json:"scopes" yaml:"scopes"`
	// A friendly name for the resource server.
	UserPoolResourceServerName *string `field:"optional" json:"userPoolResourceServerName" yaml:"userPoolResourceServerName"`
	// The user pool to add this resource server to.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
}

