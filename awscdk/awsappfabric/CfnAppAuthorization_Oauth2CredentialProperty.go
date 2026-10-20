package awsappfabric


// Contains OAuth2 client credential information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oauth2CredentialProperty := &Oauth2CredentialProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-oauth2credential.html
//
type CfnAppAuthorization_Oauth2CredentialProperty struct {
	// The client ID of the client application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-oauth2credential.html#cfn-appfabric-appauthorization-oauth2credential-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret of the client application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-oauth2credential.html#cfn-appfabric-appauthorization-oauth2credential-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

