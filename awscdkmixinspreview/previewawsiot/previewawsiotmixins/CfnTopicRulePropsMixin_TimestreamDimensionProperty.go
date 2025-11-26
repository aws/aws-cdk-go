package previewawsiotmixins


// Metadata attributes of the time series that are written in each measure record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timestreamDimensionProperty := &TimestreamDimensionProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamdimension.html
//
type CfnTopicRulePropsMixin_TimestreamDimensionProperty struct {
	// The metadata dimension name.
	//
	// This is the name of the column in the Amazon Timestream database table record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamdimension.html#cfn-iot-topicrule-timestreamdimension-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value to write in this column of the database record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamdimension.html#cfn-iot-topicrule-timestreamdimension-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

