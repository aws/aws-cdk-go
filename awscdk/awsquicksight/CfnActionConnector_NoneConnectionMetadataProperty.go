package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noneConnectionMetadataProperty := &NoneConnectionMetadataProperty{
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-noneconnectionmetadata.html
//
type CfnActionConnector_NoneConnectionMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-noneconnectionmetadata.html#cfn-quicksight-actionconnector-noneconnectionmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"required" json:"baseEndpoint" yaml:"baseEndpoint"`
}

