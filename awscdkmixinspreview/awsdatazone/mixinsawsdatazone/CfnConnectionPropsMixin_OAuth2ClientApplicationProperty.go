package mixinsawsdatazone


// The OAuth2Client application.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2clientapplication.html
//
type CfnConnectionPropsMixin_OAuth2ClientApplicationProperty struct {
	// The AWS managed client application reference in the OAuth2Client application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2clientapplication.html#cfn-datazone-connection-oauth2clientapplication-awsmanagedclientapplicationreference
	//
	AwsManagedClientApplicationReference *string `field:"optional" json:"awsManagedClientApplicationReference" yaml:"awsManagedClientApplicationReference"`
	// The user managed client application client ID in the OAuth2Client application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2clientapplication.html#cfn-datazone-connection-oauth2clientapplication-usermanagedclientapplicationclientid
	//
	UserManagedClientApplicationClientId *string `field:"optional" json:"userManagedClientApplicationClientId" yaml:"userManagedClientApplicationClientId"`
}

