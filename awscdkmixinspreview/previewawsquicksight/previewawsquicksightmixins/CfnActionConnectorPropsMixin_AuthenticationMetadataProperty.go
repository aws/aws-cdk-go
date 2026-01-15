package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticationMetadataProperty := &AuthenticationMetadataProperty{
//   	ApiKeyConnectionMetadata: &APIKeyConnectionMetadataProperty{
//   		ApiKey: jsii.String("apiKey"),
//   		BaseEndpoint: jsii.String("baseEndpoint"),
//   		Email: jsii.String("email"),
//   	},
//   	AuthorizationCodeGrantMetadata: &AuthorizationCodeGrantMetadataProperty{
//   		AuthorizationCodeGrantCredentialsDetails: &AuthorizationCodeGrantCredentialsDetailsProperty{
//   			AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   				AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//   			},
//   		},
//   		AuthorizationCodeGrantCredentialsSource: jsii.String("authorizationCodeGrantCredentialsSource"),
//   		BaseEndpoint: jsii.String("baseEndpoint"),
//   		RedirectUrl: jsii.String("redirectUrl"),
//   	},
//   	BasicAuthConnectionMetadata: &BasicAuthConnectionMetadataProperty{
//   		BaseEndpoint: jsii.String("baseEndpoint"),
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	ClientCredentialsGrantMetadata: &ClientCredentialsGrantMetadataProperty{
//   		BaseEndpoint: jsii.String("baseEndpoint"),
//   		ClientCredentialsDetails: &ClientCredentialsDetailsProperty{
//   			ClientCredentialsGrantDetails: &ClientCredentialsGrantDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//   			},
//   		},
//   		ClientCredentialsSource: jsii.String("clientCredentialsSource"),
//   	},
//   	IamConnectionMetadata: map[string]*string{
//   		"roleArn": jsii.String("roleArn"),
//   	},
//   	NoneConnectionMetadata: &NoneConnectionMetadataProperty{
//   		BaseEndpoint: jsii.String("baseEndpoint"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html
//
type CfnActionConnectorPropsMixin_AuthenticationMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html#cfn-quicksight-actionconnector-authenticationmetadata-apikeyconnectionmetadata
	//
	ApiKeyConnectionMetadata interface{} `field:"optional" json:"apiKeyConnectionMetadata" yaml:"apiKeyConnectionMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html#cfn-quicksight-actionconnector-authenticationmetadata-authorizationcodegrantmetadata
	//
	AuthorizationCodeGrantMetadata interface{} `field:"optional" json:"authorizationCodeGrantMetadata" yaml:"authorizationCodeGrantMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html#cfn-quicksight-actionconnector-authenticationmetadata-basicauthconnectionmetadata
	//
	BasicAuthConnectionMetadata interface{} `field:"optional" json:"basicAuthConnectionMetadata" yaml:"basicAuthConnectionMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html#cfn-quicksight-actionconnector-authenticationmetadata-clientcredentialsgrantmetadata
	//
	ClientCredentialsGrantMetadata interface{} `field:"optional" json:"clientCredentialsGrantMetadata" yaml:"clientCredentialsGrantMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html#cfn-quicksight-actionconnector-authenticationmetadata-iamconnectionmetadata
	//
	IamConnectionMetadata interface{} `field:"optional" json:"iamConnectionMetadata" yaml:"iamConnectionMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authenticationmetadata.html#cfn-quicksight-actionconnector-authenticationmetadata-noneconnectionmetadata
	//
	NoneConnectionMetadata interface{} `field:"optional" json:"noneConnectionMetadata" yaml:"noneConnectionMetadata"`
}

