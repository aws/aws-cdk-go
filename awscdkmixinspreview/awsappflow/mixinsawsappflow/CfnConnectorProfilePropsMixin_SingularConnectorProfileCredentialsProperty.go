package mixinsawsappflow


// The connector-specific profile credentials required when using Singular.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   singularConnectorProfileCredentialsProperty := &SingularConnectorProfileCredentialsProperty{
//   	ApiKey: jsii.String("apiKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-singularconnectorprofilecredentials.html
//
type CfnConnectorProfilePropsMixin_SingularConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-singularconnectorprofilecredentials.html#cfn-appflow-connectorprofile-singularconnectorprofilecredentials-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
}

