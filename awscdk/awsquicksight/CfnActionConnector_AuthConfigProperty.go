package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authConfigProperty := &AuthConfigProperty{
//   	AuthenticationMetadata: &AuthenticationMetadataProperty{
//   		ApiKeyConnectionMetadata: &APIKeyConnectionMetadataProperty{
//   			ApiKey: jsii.String("apiKey"),
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//
//   			// the properties below are optional
//   			Email: jsii.String("email"),
//   		},
//   		AuthorizationCodeGrantMetadata: &AuthorizationCodeGrantMetadataProperty{
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   			RedirectUrl: jsii.String("redirectUrl"),
//
//   			// the properties below are optional
//   			AuthorizationCodeGrantCredentialsDetails: &AuthorizationCodeGrantCredentialsDetailsProperty{
//   				AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   					AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//   				},
//   			},
//   			AuthorizationCodeGrantCredentialsSource: jsii.String("authorizationCodeGrantCredentialsSource"),
//   		},
//   		BasicAuthConnectionMetadata: &BasicAuthConnectionMetadataProperty{
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		ClientCredentialsGrantMetadata: &ClientCredentialsGrantMetadataProperty{
//   			BaseEndpoint: jsii.String("baseEndpoint"),
//
//   			// the properties below are optional
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
type CfnActionConnector_AuthConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authconfig.html#cfn-quicksight-actionconnector-authconfig-authenticationmetadata
	//
	AuthenticationMetadata interface{} `field:"required" json:"authenticationMetadata" yaml:"authenticationMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authconfig.html#cfn-quicksight-actionconnector-authconfig-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
}

