package awslogs


// This object defines one value type that will be converted using the [typeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-typeConverter) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typeConverterEntryProperty := &TypeConverterEntryProperty{
//   	Key: jsii.String("key"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-typeconverterentry.html
//
type CfnTransformer_TypeConverterEntryProperty struct {
	// The key with the value that is to be converted to a different type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-typeconverterentry.html#cfn-logs-transformer-typeconverterentry-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The type to convert the field value to.
	//
	// Valid values are `integer` , `double` , `string` and `boolean` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-typeconverterentry.html#cfn-logs-transformer-typeconverterentry-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

