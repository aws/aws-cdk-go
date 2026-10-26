package awseventsv2


// Parameters for invoking an AWS service API as a universal service integration target, used when TargetArn has the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   universalTargetParametersProperty := &UniversalTargetParametersProperty{
//   	Input: jsii.String("input"),
//   	InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-universaltargetparameters.html
//
type CfnSubscriberPropsMixin_UniversalTargetParametersProperty struct {
	// JSON string or JSONata expression that produces the API request.
	//
	// Supports {% ... %} JSONata expressions for dynamic values from the event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-universaltargetparameters.html#cfn-eventsv2-subscriber-universaltargetparameters-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Timeout in seconds for each invocation of the target (1-30, default 30).
	//
	// Must be a literal integer written as a string; JSONata expressions are not supported for this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-universaltargetparameters.html#cfn-eventsv2-subscriber-universaltargetparameters-invocationtimeoutseconds
	//
	InvocationTimeoutSeconds *string `field:"optional" json:"invocationTimeoutSeconds" yaml:"invocationTimeoutSeconds"`
}

