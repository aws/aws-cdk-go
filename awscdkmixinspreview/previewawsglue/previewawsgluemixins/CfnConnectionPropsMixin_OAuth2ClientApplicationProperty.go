package previewawsgluemixins


// The OAuth2 client app used for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oAuth2ClientApplicationProperty := &OAuth2ClientApplicationProperty{
//   	AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   	UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2clientapplication.html
//
type CfnConnectionPropsMixin_OAuth2ClientApplicationProperty struct {
	// The reference to the SaaS-side client app that is AWS managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2clientapplication.html#cfn-glue-connection-oauth2clientapplication-awsmanagedclientapplicationreference
	//
	AwsManagedClientApplicationReference *string `field:"optional" json:"awsManagedClientApplicationReference" yaml:"awsManagedClientApplicationReference"`
	// The client application clientID if the ClientAppType is `USER_MANAGED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2clientapplication.html#cfn-glue-connection-oauth2clientapplication-usermanagedclientapplicationclientid
	//
	UserManagedClientApplicationClientId *string `field:"optional" json:"userManagedClientApplicationClientId" yaml:"userManagedClientApplicationClientId"`
}

