package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationCodeGrantMetadataProperty := &AuthorizationCodeGrantMetadataProperty{
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//   	RedirectUrl: jsii.String("redirectUrl"),
//
//   	// the properties below are optional
//   	AuthorizationCodeGrantCredentialsDetails: &AuthorizationCodeGrantCredentialsDetailsProperty{
//   		AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   	},
//   	AuthorizationCodeGrantCredentialsSource: jsii.String("authorizationCodeGrantCredentialsSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantmetadata.html
//
type CfnActionConnector_AuthorizationCodeGrantMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantmetadata.html#cfn-quicksight-actionconnector-authorizationcodegrantmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"required" json:"baseEndpoint" yaml:"baseEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantmetadata.html#cfn-quicksight-actionconnector-authorizationcodegrantmetadata-redirecturl
	//
	RedirectUrl *string `field:"required" json:"redirectUrl" yaml:"redirectUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantmetadata.html#cfn-quicksight-actionconnector-authorizationcodegrantmetadata-authorizationcodegrantcredentialsdetails
	//
	AuthorizationCodeGrantCredentialsDetails interface{} `field:"optional" json:"authorizationCodeGrantCredentialsDetails" yaml:"authorizationCodeGrantCredentialsDetails"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-authorizationcodegrantmetadata.html#cfn-quicksight-actionconnector-authorizationcodegrantmetadata-authorizationcodegrantcredentialssource
	//
	AuthorizationCodeGrantCredentialsSource *string `field:"optional" json:"authorizationCodeGrantCredentialsSource" yaml:"authorizationCodeGrantCredentialsSource"`
}

