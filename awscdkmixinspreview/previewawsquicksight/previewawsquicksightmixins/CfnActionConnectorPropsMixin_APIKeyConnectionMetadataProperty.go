package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aPIKeyConnectionMetadataProperty := &APIKeyConnectionMetadataProperty{
//   	ApiKey: jsii.String("apiKey"),
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//   	Email: jsii.String("email"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html
//
type CfnActionConnectorPropsMixin_APIKeyConnectionMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html#cfn-quicksight-actionconnector-apikeyconnectionmetadata-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html#cfn-quicksight-actionconnector-apikeyconnectionmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"optional" json:"baseEndpoint" yaml:"baseEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html#cfn-quicksight-actionconnector-apikeyconnectionmetadata-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
}

