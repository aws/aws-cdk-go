package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &EventSourceProperty{
//   	Properties: &AlexaSkillEventProperty{
//   		SkillId: jsii.String("skillId"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventsource.html
//
type CfnFunction_EventSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventsource.html#cfn-serverless-function-eventsource-properties
	//
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventsource.html#cfn-serverless-function-eventsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

