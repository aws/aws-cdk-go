package awsmedialive


// Settings that apply only if the input is a MediaConnect input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaConnectFlowRequestProperty := &MediaConnectFlowRequestProperty{
//   	FlowArn: jsii.String("flowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-mediaconnectflowrequest.html
//
type CfnInput_MediaConnectFlowRequestProperty struct {
	// The ARN of one or two MediaConnect flows that are the sources for this MediaConnect input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-mediaconnectflowrequest.html#cfn-medialive-input-mediaconnectflowrequest-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
}

