package previewawsquicksightmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnActionConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnActionConnectorMixinProps := &CfnActionConnectorMixinProps{
//   	ActionConnectorId: jsii.String("actionConnectorId"),
//   	AuthenticationConfig: &AuthConfigProperty{
//   		AuthenticationMetadata: &AuthenticationMetadataProperty{
//   			ApiKeyConnectionMetadata: &APIKeyConnectionMetadataProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				Email: jsii.String("email"),
//   			},
//   			AuthorizationCodeGrantMetadata: &AuthorizationCodeGrantMetadataProperty{
//   				AuthorizationCodeGrantCredentialsDetails: &AuthorizationCodeGrantCredentialsDetailsProperty{
//   					AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   						AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   						ClientId: jsii.String("clientId"),
//   						ClientSecret: jsii.String("clientSecret"),
//   						TokenEndpoint: jsii.String("tokenEndpoint"),
//   					},
//   				},
//   				AuthorizationCodeGrantCredentialsSource: jsii.String("authorizationCodeGrantCredentialsSource"),
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				RedirectUrl: jsii.String("redirectUrl"),
//   			},
//   			BasicAuthConnectionMetadata: &BasicAuthConnectionMetadataProperty{
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			ClientCredentialsGrantMetadata: &ClientCredentialsGrantMetadataProperty{
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				ClientCredentialsDetails: &ClientCredentialsDetailsProperty{
//   					ClientCredentialsGrantDetails: &ClientCredentialsGrantDetailsProperty{
//   						ClientId: jsii.String("clientId"),
//   						ClientSecret: jsii.String("clientSecret"),
//   						TokenEndpoint: jsii.String("tokenEndpoint"),
//   					},
//   				},
//   				ClientCredentialsSource: jsii.String("clientCredentialsSource"),
//   			},
//   			IamConnectionMetadata: map[string]*string{
//   				"roleArn": jsii.String("roleArn"),
//   			},
//   			NoneConnectionMetadata: &NoneConnectionMetadataProperty{
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   			},
//   		},
//   		AuthenticationType: jsii.String("authenticationType"),
//   	},
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html
//
type CfnActionConnectorMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-actionconnectorid
	//
	ActionConnectorId *string `field:"optional" json:"actionConnectorId" yaml:"actionConnectorId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-authenticationconfig
	//
	AuthenticationConfig interface{} `field:"optional" json:"authenticationConfig" yaml:"authenticationConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html#cfn-quicksight-actionconnector-vpcconnectionarn
	//
	VpcConnectionArn *string `field:"optional" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

