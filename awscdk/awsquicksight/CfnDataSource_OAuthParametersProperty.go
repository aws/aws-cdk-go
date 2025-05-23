package awsquicksight


// An object that contains information needed to create a data source connection that uses OAuth client credentials.
//
// This option is available for data source connections that are made with Snowflake and Starburst.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthParametersProperty := &OAuthParametersProperty{
//   	TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   	// the properties below are optional
//   	IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   	IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   		VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   	OAuthScope: jsii.String("oAuthScope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html
//
type CfnDataSource_OAuthParametersProperty struct {
	// The token endpoint URL of the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-tokenproviderurl
	//
	TokenProviderUrl *string `field:"required" json:"tokenProviderUrl" yaml:"tokenProviderUrl"`
	// The resource uri of the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-identityproviderresourceuri
	//
	IdentityProviderResourceUri *string `field:"optional" json:"identityProviderResourceUri" yaml:"identityProviderResourceUri"`
	// <p>VPC connection properties.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-identityprovidervpcconnectionproperties
	//
	IdentityProviderVpcConnectionProperties interface{} `field:"optional" json:"identityProviderVpcConnectionProperties" yaml:"identityProviderVpcConnectionProperties"`
	// The OAuth scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-oauthscope
	//
	OAuthScope *string `field:"optional" json:"oAuthScope" yaml:"oAuthScope"`
}

