package awsconnect


// The outbound caller ID name, number, and outbound whisper flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outboundCallerConfigProperty := &OutboundCallerConfigProperty{
//   	OutboundCallerIdName: jsii.String("outboundCallerIdName"),
//   	OutboundCallerIdNumberArn: jsii.String("outboundCallerIdNumberArn"),
//   	OutboundFlowArn: jsii.String("outboundFlowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html
//
type CfnQueue_OutboundCallerConfigProperty struct {
	// The caller ID name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html#cfn-connect-queue-outboundcallerconfig-outboundcalleridname
	//
	OutboundCallerIdName *string `field:"optional" json:"outboundCallerIdName" yaml:"outboundCallerIdName"`
	// The caller ID number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html#cfn-connect-queue-outboundcallerconfig-outboundcalleridnumberarn
	//
	OutboundCallerIdNumberArn *string `field:"optional" json:"outboundCallerIdNumberArn" yaml:"outboundCallerIdNumberArn"`
	// The outbound whisper flow to be used during an outbound call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html#cfn-connect-queue-outboundcallerconfig-outboundflowarn
	//
	OutboundFlowArn *string `field:"optional" json:"outboundFlowArn" yaml:"outboundFlowArn"`
}

