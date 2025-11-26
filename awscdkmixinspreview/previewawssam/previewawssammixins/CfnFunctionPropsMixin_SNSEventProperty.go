package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sNSEventProperty := &SNSEventProperty{
//   	Topic: jsii.String("topic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-snsevent.html
//
type CfnFunctionPropsMixin_SNSEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-snsevent.html#cfn-serverless-function-snsevent-topic
	//
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

