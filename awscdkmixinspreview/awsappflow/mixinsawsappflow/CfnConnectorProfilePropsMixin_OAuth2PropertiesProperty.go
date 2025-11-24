package mixinsawsappflow


// The OAuth 2.0 properties required for OAuth 2.0 authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oAuth2PropertiesProperty := &OAuth2PropertiesProperty{
//   	OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   	TokenUrl: jsii.String("tokenUrl"),
//   	TokenUrlCustomProperties: map[string]*string{
//   		"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauth2properties.html
//
type CfnConnectorProfilePropsMixin_OAuth2PropertiesProperty struct {
	// The OAuth 2.0 grant type used by connector for OAuth 2.0 authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauth2properties.html#cfn-appflow-connectorprofile-oauth2properties-oauth2granttype
	//
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The token URL required for OAuth 2.0 authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauth2properties.html#cfn-appflow-connectorprofile-oauth2properties-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Associates your token URL with a map of properties that you define.
	//
	// Use this parameter to provide any additional details that the connector requires to authenticate your request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-oauth2properties.html#cfn-appflow-connectorprofile-oauth2properties-tokenurlcustomproperties
	//
	TokenUrlCustomProperties interface{} `field:"optional" json:"tokenUrlCustomProperties" yaml:"tokenUrlCustomProperties"`
}

