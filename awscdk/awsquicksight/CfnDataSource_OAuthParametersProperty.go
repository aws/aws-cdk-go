package awsquicksight


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-tokenproviderurl
	//
	TokenProviderUrl *string `field:"required" json:"tokenProviderUrl" yaml:"tokenProviderUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-identityproviderresourceuri
	//
	IdentityProviderResourceUri *string `field:"optional" json:"identityProviderResourceUri" yaml:"identityProviderResourceUri"`
	// <p>VPC connection properties.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-identityprovidervpcconnectionproperties
	//
	IdentityProviderVpcConnectionProperties interface{} `field:"optional" json:"identityProviderVpcConnectionProperties" yaml:"identityProviderVpcConnectionProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-oauthparameters.html#cfn-quicksight-datasource-oauthparameters-oauthscope
	//
	OAuthScope *string `field:"optional" json:"oAuthScope" yaml:"oAuthScope"`
}

