package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonOpenSearchServerlessBufferingHintsProperty := &amazonOpenSearchServerlessBufferingHintsProperty{
//   	intervalInSeconds: jsii.Number(123),
//   	sizeInMBs: jsii.Number(123),
//   }
//
type CfnDeliveryStream_AmazonOpenSearchServerlessBufferingHintsProperty struct {
	// `CfnDeliveryStream.AmazonOpenSearchServerlessBufferingHintsProperty.IntervalInSeconds`.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessBufferingHintsProperty.SizeInMBs`.
	SizeInMBs *float64 `field:"optional" json:"sizeInMBs" yaml:"sizeInMBs"`
}

