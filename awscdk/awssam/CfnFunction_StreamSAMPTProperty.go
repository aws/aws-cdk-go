package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSAMPTProperty := &StreamSAMPTProperty{
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-streamsampt.html
//
type CfnFunction_StreamSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-streamsampt.html#cfn-serverless-function-streamsampt-streamname
	//
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

