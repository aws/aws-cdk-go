package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientCredentialsGrantDetailsProperty := &ClientCredentialsGrantDetailsProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantdetails.html
//
type CfnActionConnectorPropsMixin_ClientCredentialsGrantDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantdetails.html#cfn-quicksight-actionconnector-clientcredentialsgrantdetails-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantdetails.html#cfn-quicksight-actionconnector-clientcredentialsgrantdetails-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-clientcredentialsgrantdetails.html#cfn-quicksight-actionconnector-clientcredentialsgrantdetails-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

