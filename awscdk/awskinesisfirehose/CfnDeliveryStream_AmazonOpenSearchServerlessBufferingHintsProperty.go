package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonOpenSearchServerlessBufferingHintsProperty := &AmazonOpenSearchServerlessBufferingHintsProperty{
//   	IntervalInSeconds: jsii.Number(123),
//   	SizeInMBs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessbufferinghints.html
//
type CfnDeliveryStream_AmazonOpenSearchServerlessBufferingHintsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessbufferinghints.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessbufferinghints-intervalinseconds
	//
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessbufferinghints.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessbufferinghints-sizeinmbs
	//
	SizeInMBs *float64 `field:"optional" json:"sizeInMBs" yaml:"sizeInMBs"`
}

