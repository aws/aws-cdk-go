package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPIKeyConnectionMetadataProperty := &APIKeyConnectionMetadataProperty{
//   	ApiKey: jsii.String("apiKey"),
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//
//   	// the properties below are optional
//   	Email: jsii.String("email"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html
//
type CfnActionConnector_APIKeyConnectionMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html#cfn-quicksight-actionconnector-apikeyconnectionmetadata-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html#cfn-quicksight-actionconnector-apikeyconnectionmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"required" json:"baseEndpoint" yaml:"baseEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.html#cfn-quicksight-actionconnector-apikeyconnectionmetadata-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
}

