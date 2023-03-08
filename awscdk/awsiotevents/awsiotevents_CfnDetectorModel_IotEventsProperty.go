package awsiotevents


// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotEventsProperty := &iotEventsProperty{
//   	inputName: jsii.String("inputName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_IotEventsProperty struct {
	// The name of the AWS IoT Events input where the data is sent.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// You can configure the action payload when you send a message to an AWS IoT Events input.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

