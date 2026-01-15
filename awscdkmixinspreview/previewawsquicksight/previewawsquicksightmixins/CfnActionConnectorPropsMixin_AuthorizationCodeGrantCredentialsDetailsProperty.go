package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authorizationCodeGrantCredentialsDetailsProperty := &AuthorizationCodeGrantCredentialsDetailsProperty{
//   	AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantcredentialsdetails.html
//
type CfnActionConnectorPropsMixin_AuthorizationCodeGrantCredentialsDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantcredentialsdetails.html#cfn-quicksight-actionconnector-authorizationcodegrantcredentialsdetails-authorizationcodegrantdetails
	//
	AuthorizationCodeGrantDetails interface{} `field:"optional" json:"authorizationCodeGrantDetails" yaml:"authorizationCodeGrantDetails"`
}

