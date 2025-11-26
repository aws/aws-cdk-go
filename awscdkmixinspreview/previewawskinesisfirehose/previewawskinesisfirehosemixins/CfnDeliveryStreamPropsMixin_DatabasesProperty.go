package previewawskinesisfirehosemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   databasesProperty := &DatabasesProperty{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Include: []*string{
//   		jsii.String("include"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databases.html
//
type CfnDeliveryStreamPropsMixin_DatabasesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databases.html#cfn-kinesisfirehose-deliverystream-databases-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databases.html#cfn-kinesisfirehose-deliverystream-databases-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

