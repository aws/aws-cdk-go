package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authConfigProperty := &AuthConfigProperty{
//   	AuthenticationMetadata: &AuthenticationMetadataProperty{
//   		ApiKeyConnectionMetadata: &APIKeyConnectionMetadataProperty{
//   			ApiKey: jsii.String("apiKey"),
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   			Email: jsii.String("email"),
//   		},
//   		AuthorizationCodeGrantMetadata: &AuthorizationCodeGrantMetadataProperty{
//   			AuthorizationCodeGrantCredentialsDetails: &AuthorizationCodeGrantCredentialsDetailsProperty{
//   				AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   					AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//   				},
//   			},
//   			AuthorizationCodeGrantCredentialsSource: jsii.String("authorizationCodeGrantCredentialsSource"),
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   			RedirectUrl: jsii.String("redirectUrl"),
//   		},
//   		BasicAuthConnectionMetadata: &BasicAuthConnectionMetadataProperty{
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		ClientCredentialsGrantMetadata: &ClientCredentialsGrantMetadataProperty{
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   			ClientCredentialsDetails: &ClientCredentialsDetailsProperty{
//   				ClientCredentialsGrantDetails: &ClientCredentialsGrantDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//   				},
//   			},
//   			ClientCredentialsSource: jsii.String("clientCredentialsSource"),
//   		},
//   		IamConnectionMetadata: map[string]*string{
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   		NoneConnectionMetadata: &NoneConnectionMetadataProperty{
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   		},
//   	},
//   	AuthenticationType: jsii.String("authenticationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authconfig.html
//
type CfnActionConnectorPropsMixin_AuthConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authconfig.html#cfn-quicksight-actionconnector-authconfig-authenticationmetadata
	//
	AuthenticationMetadata interface{} `field:"optional" json:"authenticationMetadata" yaml:"authenticationMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authconfig.html#cfn-quicksight-actionconnector-authconfig-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
}

