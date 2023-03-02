package awsappflow


// The OAuth properties required for OAuth type authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthPropertiesProperty := &oAuthPropertiesProperty{
//   	authCodeUrl: jsii.String("authCodeUrl"),
//   	oAuthScopes: []*string{
//   		jsii.String("oAuthScopes"),
//   	},
//   	tokenUrl: jsii.String("tokenUrl"),
//   }
//
type CfnConnectorProfile_OAuthPropertiesProperty struct {
	// The authorization code url required to redirect to SAP Login Page to fetch authorization code for OAuth type authentication.
	AuthCodeUrl *string `field:"optional" json:"authCodeUrl" yaml:"authCodeUrl"`
	// The OAuth scopes required for OAuth type authentication.
	OAuthScopes *[]*string `field:"optional" json:"oAuthScopes" yaml:"oAuthScopes"`
	// The token url required to fetch access/refresh tokens using authorization code and also to refresh expired access token using refresh token.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}

