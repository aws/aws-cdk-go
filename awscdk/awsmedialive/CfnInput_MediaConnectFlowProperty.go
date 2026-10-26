package awsmedialive


// The settings for a MediaConnect Flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaConnectFlowProperty := &MediaConnectFlowProperty{
//   	FlowArn: jsii.String("flowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-mediaconnectflow.html
//
type CfnInput_MediaConnectFlowProperty struct {
	// The unique ARN of the MediaConnect Flow being used as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-mediaconnectflow.html#cfn-medialive-input-mediaconnectflow-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
}

