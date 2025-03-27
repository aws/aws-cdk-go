package awsdatazone


// Authorization Code Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationCodePropertiesProperty := &AuthorizationCodePropertiesProperty{
//   	AuthorizationCode: jsii.String("authorizationCode"),
//   	RedirectUri: jsii.String("redirectUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authorizationcodeproperties.html
//
type CfnConnection_AuthorizationCodePropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authorizationcodeproperties.html#cfn-datazone-connection-authorizationcodeproperties-authorizationcode
	//
	AuthorizationCode *string `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authorizationcodeproperties.html#cfn-datazone-connection-authorizationcodeproperties-redirecturi
	//
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

