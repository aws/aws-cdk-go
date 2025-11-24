package mixinsawsappflow


// The OAuth properties required for OAuth type authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oAuthPropertiesProperty := &OAuthPropertiesProperty{
//   	AuthCodeUrl: jsii.String("authCodeUrl"),
//   	OAuthScopes: []*string{
//   		jsii.String("oAuthScopes"),
//   	},
//   	TokenUrl: jsii.String("tokenUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauthproperties.html
//
type CfnConnectorProfilePropsMixin_OAuthPropertiesProperty struct {
	// The authorization code url required to redirect to SAP Login Page to fetch authorization code for OAuth type authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauthproperties.html#cfn-appflow-connectorprofile-oauthproperties-authcodeurl
	//
	AuthCodeUrl *string `field:"optional" json:"authCodeUrl" yaml:"authCodeUrl"`
	// The OAuth scopes required for OAuth type authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauthproperties.html#cfn-appflow-connectorprofile-oauthproperties-oauthscopes
	//
	OAuthScopes *[]*string `field:"optional" json:"oAuthScopes" yaml:"oAuthScopes"`
	// The token url required to fetch access/refresh tokens using authorization code and also to refresh expired access token using refresh token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauthproperties.html#cfn-appflow-connectorprofile-oauthproperties-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}

