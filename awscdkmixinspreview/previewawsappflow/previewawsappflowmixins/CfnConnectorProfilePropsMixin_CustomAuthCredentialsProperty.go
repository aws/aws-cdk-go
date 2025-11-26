package previewawsappflowmixins


// The custom credentials required for custom authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customAuthCredentialsProperty := &CustomAuthCredentialsProperty{
//   	CredentialsMap: map[string]*string{
//   		"credentialsMapKey": jsii.String("credentialsMap"),
//   	},
//   	CustomAuthenticationType: jsii.String("customAuthenticationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html
//
type CfnConnectorProfilePropsMixin_CustomAuthCredentialsProperty struct {
	// A map that holds custom authentication credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-credentialsmap
	//
	CredentialsMap interface{} `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
	// The custom authentication type that the connector uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-customauthenticationtype
	//
	CustomAuthenticationType *string `field:"optional" json:"customAuthenticationType" yaml:"customAuthenticationType"`
}

