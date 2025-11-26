package previewawsioteventsmixins


// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iotEventsProperty := &IotEventsProperty{
//   	InputName: jsii.String("inputName"),
//   	Payload: &PayloadProperty{
//   		ContentExpression: jsii.String("contentExpression"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html
//
type CfnDetectorModelPropsMixin_IotEventsProperty struct {
	// The name of the AWS IoT Events input where the data is sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html#cfn-iotevents-detectormodel-iotevents-inputname
	//
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// You can configure the action payload when you send a message to an AWS IoT Events input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html#cfn-iotevents-detectormodel-iotevents-payload
	//
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

