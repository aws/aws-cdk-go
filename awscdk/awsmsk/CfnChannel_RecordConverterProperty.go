package awsmsk


// Record converter configuration for a topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordConverterProperty := &RecordConverterProperty{
//   	ValueConverter: jsii.String("valueConverter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-recordconverter.html
//
type CfnChannel_RecordConverterProperty struct {
	// Value converter for topic data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-recordconverter.html#cfn-msk-channel-recordconverter-valueconverter
	//
	ValueConverter *string `field:"required" json:"valueConverter" yaml:"valueConverter"`
}

