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
//   mediaConnectFlowRequestProperty := &mediaConnectFlowRequestProperty{
//   	flowArn: jsii.String("flowArn"),
//   }
//
type CfnInput_MediaConnectFlowRequestProperty struct {
	// The ARN of one or two MediaConnect flows that are the sources for this MediaConnect input.
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
}

