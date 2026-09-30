package awseventsv2


// JSONata transform settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonataConfigurationProperty := &JsonataConfigurationProperty{
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-jsonataconfiguration.html
//
type CfnSubscriber_JsonataConfigurationProperty struct {
	// The JSONata expression that transforms the event, enclosed in {% %} delimiters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-jsonataconfiguration.html#cfn-eventsv2-subscriber-jsonataconfiguration-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

