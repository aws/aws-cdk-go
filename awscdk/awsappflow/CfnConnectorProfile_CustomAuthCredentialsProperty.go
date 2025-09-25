package awsappflow


// The custom credentials required for custom authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAuthCredentialsProperty := &CustomAuthCredentialsProperty{
//   	CustomAuthenticationType: jsii.String("customAuthenticationType"),
//
//   	// the properties below are optional
//   	CredentialsMap: map[string]*string{
//   		"credentialsMapKey": jsii.String("credentialsMap"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html
//
type CfnConnectorProfile_CustomAuthCredentialsProperty struct {
	// The custom authentication type that the connector uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-customauthenticationtype
	//
	CustomAuthenticationType *string `field:"required" json:"customAuthenticationType" yaml:"customAuthenticationType"`
	// A map that holds custom authentication credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-credentialsmap
	//
	CredentialsMap interface{} `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
}

