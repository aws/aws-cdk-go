package awslambda


// The filter criteria to control event filtering.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterCriteriaProperty := &FilterCriteriaProperty{
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-filtercriteria.html
//
type CfnEventSourceMapping_FilterCriteriaProperty struct {
	// List of filters of this FilterCriteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-filtercriteria.html#cfn-lambda-eventsourcemapping-filtercriteria-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

