package previewawsdatazonemixins


// The authorization code properties of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authorizationCodePropertiesProperty := &AuthorizationCodePropertiesProperty{
//   	AuthorizationCode: jsii.String("authorizationCode"),
//   	RedirectUri: jsii.String("redirectUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authorizationcodeproperties.html
//
type CfnConnectionPropsMixin_AuthorizationCodePropertiesProperty struct {
	// The authorization code of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authorizationcodeproperties.html#cfn-datazone-connection-authorizationcodeproperties-authorizationcode
	//
	AuthorizationCode *string `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// The redirect URI of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authorizationcodeproperties.html#cfn-datazone-connection-authorizationcodeproperties-redirecturi
	//
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

