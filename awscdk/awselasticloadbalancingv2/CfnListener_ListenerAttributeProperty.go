package awselasticloadbalancingv2


// Information about a listener attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerAttributeProperty := &ListenerAttributeProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-listenerattribute.html
//
type CfnListener_ListenerAttributeProperty struct {
	// The name of the attribute.
	//
	// The following attribute is supported by Network Load Balancers, and Gateway Load Balancers.
	//
	// - `tcp.idle_timeout.seconds` - The tcp idle timeout value, in seconds. The valid range is 60-6000 seconds. The default is 350 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-listenerattribute.html#cfn-elasticloadbalancingv2-listener-listenerattribute-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-listenerattribute.html#cfn-elasticloadbalancingv2-listener-listenerattribute-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

