package awslambda


// An object that contains the filters for an event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterCriteriaProperty := &filterCriteriaProperty{
//   	filters: []interface{}{
//   		&filterProperty{
//   			pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
type CfnEventSourceMapping_FilterCriteriaProperty struct {
	// A list of filters.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

