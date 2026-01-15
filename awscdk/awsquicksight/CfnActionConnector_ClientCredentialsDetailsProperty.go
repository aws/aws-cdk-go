package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientCredentialsDetailsProperty := &ClientCredentialsDetailsProperty{
//   	ClientCredentialsGrantDetails: &ClientCredentialsGrantDetailsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsdetails.html
//
type CfnActionConnector_ClientCredentialsDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsdetails.html#cfn-quicksight-actionconnector-clientcredentialsdetails-clientcredentialsgrantdetails
	//
	ClientCredentialsGrantDetails interface{} `field:"required" json:"clientCredentialsGrantDetails" yaml:"clientCredentialsGrantDetails"`
}

