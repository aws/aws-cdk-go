package awseventsv2


// A single filter entry: an event pattern and the scope of the event it is evaluated against.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Pattern: jsii.String("pattern"),
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-filter.html
//
type CfnSubscriber_FilterProperty struct {
	// The event pattern, as a JSON string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-filter.html#cfn-eventsv2-subscriber-filter-pattern
	//
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Which part of the event the pattern is evaluated against: DATA (the event payload), METADATA (event metadata), or SYSTEM_METADATA (service-generated metadata).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-filter.html#cfn-eventsv2-subscriber-filter-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

