package awsglue


// Defines the predicate of the trigger, which determines when it fires.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predicateProperty := &PredicateProperty{
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			CrawlerName: jsii.String("crawlerName"),
//   			CrawlState: jsii.String("crawlState"),
//   			JobName: jsii.String("jobName"),
//   			LogicalOperator: jsii.String("logicalOperator"),
//   			State: jsii.String("state"),
//   		},
//   	},
//   	Logical: jsii.String("logical"),
//   }
//
type CfnTrigger_PredicateProperty struct {
	// A list of the conditions that determine when the trigger will fire.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// An optional field if only one condition is listed.
	//
	// If multiple conditions are listed, then this field is required.
	Logical *string `field:"optional" json:"logical" yaml:"logical"`
}

