package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDetectorModel`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html
//
type CfnDetectorModelProps struct {
	// Information that defines how a detector operates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-detectormodeldefinition
	//
	DetectorModelDefinition interface{} `field:"required" json:"detectorModelDefinition" yaml:"detectorModelDefinition"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-rolearn
	//
	RoleArn interface{} `field:"required" json:"roleArn" yaml:"roleArn"`
	// A brief description of the detector model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-detectormodeldescription
	//
	DetectorModelDescription *string `field:"optional" json:"detectorModelDescription" yaml:"detectorModelDescription"`
	// The name of the detector model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-detectormodelname
	//
	DetectorModelName *string `field:"optional" json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-evaluationmethod
	//
	EvaluationMethod *string `field:"optional" json:"evaluationMethod" yaml:"evaluationMethod"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

