package awscdkgluealpha


// Represents a trigger predicate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var job job
//
//   predicate := &Predicate{
//   	Conditions: []condition{
//   		&condition{
//   			CrawlerName: jsii.String("crawlerName"),
//   			CrawlState: glue_alpha.CrawlerState_RUNNING,
//   			Job: job,
//   			LogicalOperator: glue_alpha.ConditionLogicalOperator_EQUALS,
//   			State: glue_alpha.JobState_SUCCEEDED,
//   		},
//   	},
//   	Logical: glue_alpha.PredicateLogical_AND,
//   }
//
// Experimental.
type Predicate struct {
	// A list of the conditions that determine when the trigger will fire.
	// Default: - no conditions are provided.
	//
	// Experimental.
	Conditions *[]*Condition `field:"optional" json:"conditions" yaml:"conditions"`
	// The logical operator to be applied to the conditions.
	// Default: - ConditionLogical.AND if multiple conditions are provided, no logical operator if only one condition
	//
	// Experimental.
	Logical PredicateLogical `field:"optional" json:"logical" yaml:"logical"`
}

