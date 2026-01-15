package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientCredentialsGrantMetadataProperty := &ClientCredentialsGrantMetadataProperty{
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//
//   	// the properties below are optional
//   	ClientCredentialsDetails: &ClientCredentialsDetailsProperty{
//   		ClientCredentialsGrantDetails: &ClientCredentialsGrantDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   	},
//   	ClientCredentialsSource: jsii.String("clientCredentialsSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantmetadata.html
//
type CfnActionConnector_ClientCredentialsGrantMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantmetadata.html#cfn-quicksight-actionconnector-clientcredentialsgrantmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"required" json:"baseEndpoint" yaml:"baseEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantmetadata.html#cfn-quicksight-actionconnector-clientcredentialsgrantmetadata-clientcredentialsdetails
	//
	ClientCredentialsDetails interface{} `field:"optional" json:"clientCredentialsDetails" yaml:"clientCredentialsDetails"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantmetadata.html#cfn-quicksight-actionconnector-clientcredentialsgrantmetadata-clientcredentialssource
	//
	ClientCredentialsSource *string `field:"optional" json:"clientCredentialsSource" yaml:"clientCredentialsSource"`
}

