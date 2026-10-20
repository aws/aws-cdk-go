package awsappfabric


// Contains credential information for the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialProperty := &CredentialProperty{
//   	ApiKeyCredential: &ApiKeyCredentialProperty{
//   		ApiKey: jsii.String("apiKey"),
//   	},
//   	Oauth2Credential: &Oauth2CredentialProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-credential.html
//
type CfnAppAuthorization_CredentialProperty struct {
	// Contains API key credential information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-credential.html#cfn-appfabric-appauthorization-credential-apikeycredential
	//
	ApiKeyCredential interface{} `field:"optional" json:"apiKeyCredential" yaml:"apiKeyCredential"`
	// Contains OAuth2 client credential information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-credential.html#cfn-appfabric-appauthorization-credential-oauth2credential
	//
	Oauth2Credential interface{} `field:"optional" json:"oauth2Credential" yaml:"oauth2Credential"`
}

