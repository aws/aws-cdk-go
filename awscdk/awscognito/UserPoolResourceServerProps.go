package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
)

// Properties for the UserPoolResourceServer construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourceServerScope ResourceServerScope
//   var userPoolRef IUserPoolRef
//
//   userPoolResourceServerProps := &UserPoolResourceServerProps{
//   	Identifier: jsii.String("identifier"),
//   	UserPool: userPoolRef,
//
//   	// the properties below are optional
//   	Scopes: []ResourceServerScope{
//   		resourceServerScope,
//   	},
//   	UserPoolResourceServerName: jsii.String("userPoolResourceServerName"),
//   }
//
type UserPoolResourceServerProps struct {
	// A unique resource server identifier for the resource server.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Oauth scopes.
	// Default: - No scopes will be added.
	//
	Scopes *[]ResourceServerScope `field:"optional" json:"scopes" yaml:"scopes"`
	// A friendly name for the resource server.
	// Default: - same as `identifier`.
	//
	UserPoolResourceServerName *string `field:"optional" json:"userPoolResourceServerName" yaml:"userPoolResourceServerName"`
	// The user pool to add this resource server to.
	UserPool interfacesawscognito.IUserPoolRef `field:"required" json:"userPool" yaml:"userPool"`
}

