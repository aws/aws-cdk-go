package awsconnectcampaigns


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outboundCallConfigProperty := &outboundCallConfigProperty{
//   	connectContactFlowArn: jsii.String("connectContactFlowArn"),
//   	connectQueueArn: jsii.String("connectQueueArn"),
//
//   	// the properties below are optional
//   	connectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   }
//
type CfnCampaign_OutboundCallConfigProperty struct {
	// `CfnCampaign.OutboundCallConfigProperty.ConnectContactFlowArn`.
	ConnectContactFlowArn *string `field:"required" json:"connectContactFlowArn" yaml:"connectContactFlowArn"`
	// `CfnCampaign.OutboundCallConfigProperty.ConnectQueueArn`.
	ConnectQueueArn *string `field:"required" json:"connectQueueArn" yaml:"connectQueueArn"`
	// `CfnCampaign.OutboundCallConfigProperty.ConnectSourcePhoneNumber`.
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
}

