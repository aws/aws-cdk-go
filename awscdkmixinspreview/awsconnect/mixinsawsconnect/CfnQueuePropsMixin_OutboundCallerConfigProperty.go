package mixinsawsconnect


// The outbound caller ID name, number, and outbound whisper flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outboundCallerConfigProperty := &OutboundCallerConfigProperty{
//   	OutboundCallerIdName: jsii.String("outboundCallerIdName"),
//   	OutboundCallerIdNumberArn: jsii.String("outboundCallerIdNumberArn"),
//   	OutboundFlowArn: jsii.String("outboundFlowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html
//
type CfnQueuePropsMixin_OutboundCallerConfigProperty struct {
	// The caller ID name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html#cfn-connect-queue-outboundcallerconfig-outboundcalleridname
	//
	OutboundCallerIdName *string `field:"optional" json:"outboundCallerIdName" yaml:"outboundCallerIdName"`
	// The Amazon Resource Name (ARN) of the outbound caller ID number.
	//
	// > Only use the phone number ARN format that doesn't contain `instance` in the path, for example, `arn:aws:connect:us-east-1:1234567890:phone-number/uuid` . This is the same ARN format that is returned when you create a phone number using CloudFormation , or when you call the [ListPhoneNumbersV2](https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html) API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html#cfn-connect-queue-outboundcallerconfig-outboundcalleridnumberarn
	//
	OutboundCallerIdNumberArn *string `field:"optional" json:"outboundCallerIdNumberArn" yaml:"outboundCallerIdNumberArn"`
	// The Amazon Resource Name (ARN) of the outbound flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundcallerconfig.html#cfn-connect-queue-outboundcallerconfig-outboundflowarn
	//
	OutboundFlowArn *string `field:"optional" json:"outboundFlowArn" yaml:"outboundFlowArn"`
}

