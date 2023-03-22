package awscognito


// Properties for defining a `CfnUserPoolResourceServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolResourceServerProps := &cfnUserPoolResourceServerProps{
//   	identifier: jsii.String("identifier"),
//   	name: jsii.String("name"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	scopes: []interface{}{
//   		&resourceServerScopeTypeProperty{
//   			scopeDescription: jsii.String("scopeDescription"),
//   			scopeName: jsii.String("scopeName"),
//   		},
//   	},
//   }
//
type CfnUserPoolResourceServerProps struct {
	// A unique resource server identifier for the resource server.
	//
	// This could be an HTTPS endpoint where the resource server is located. For example: `https://my-weather-api.example.com` .
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// A friendly name for the resource server.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The user pool ID for the user pool.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A list of scopes.
	//
	// Each scope is a map with keys `ScopeName` and `ScopeDescription` .
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
}

