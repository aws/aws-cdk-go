package previewawskinesisfirehosemixins


// The configuration to enable automatic table creation.
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableCreationConfigurationProperty := &TableCreationConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-tablecreationconfiguration.html
//
type CfnDeliveryStreamPropsMixin_TableCreationConfigurationProperty struct {
	// Specify whether you want to enable automatic table creation.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-tablecreationconfiguration.html#cfn-kinesisfirehose-deliverystream-tablecreationconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

