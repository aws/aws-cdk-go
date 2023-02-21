package awspipes


// The collection of event patterns used to filter events.
//
// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
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
type CfnPipe_FilterCriteriaProperty struct {
	// The event patterns.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

