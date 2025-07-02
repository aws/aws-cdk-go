package awslambda


// A structure within a `FilterCriteria` object that defines an event filtering pattern.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-filter.html
//
type CfnEventSourceMapping_FilterProperty struct {
	// A filter pattern.
	//
	// For more information on the syntax of a filter pattern, see [Filter rule syntax](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-syntax) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-filter.html#cfn-lambda-eventsourcemapping-filter-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

