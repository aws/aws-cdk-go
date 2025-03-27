package awsdatazone


// OAuth2 Client Application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuth2ClientApplicationProperty := &OAuth2ClientApplicationProperty{
//   	AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   	UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2clientapplication.html
//
type CfnConnection_OAuth2ClientApplicationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2clientapplication.html#cfn-datazone-connection-oauth2clientapplication-awsmanagedclientapplicationreference
	//
	AwsManagedClientApplicationReference *string `field:"optional" json:"awsManagedClientApplicationReference" yaml:"awsManagedClientApplicationReference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2clientapplication.html#cfn-datazone-connection-oauth2clientapplication-usermanagedclientapplicationclientid
	//
	UserManagedClientApplicationClientId *string `field:"optional" json:"userManagedClientApplicationClientId" yaml:"userManagedClientApplicationClientId"`
}

