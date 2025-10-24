package awscdkgluealpha


// Represents a trigger condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var job Job
//
//   condition := &Condition{
//   	CrawlerName: jsii.String("crawlerName"),
//   	CrawlState: glue_alpha.CrawlerState_RUNNING,
//   	Job: job,
//   	LogicalOperator: glue_alpha.ConditionLogicalOperator_EQUALS,
//   	State: glue_alpha.JobState_SUCCEEDED,
//   }
//
// Experimental.
type Condition struct {
	// The name of the crawler to which this condition applies.
	// Default: - no crawler is specified.
	//
	// Experimental.
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// The condition crawler state.
	// Default: - no crawler state is specified.
	//
	// Experimental.
	CrawlState CrawlerState `field:"optional" json:"crawlState" yaml:"crawlState"`
	// The job to which this condition applies.
	// Default: - no job is specified.
	//
	// Experimental.
	Job IJob `field:"optional" json:"job" yaml:"job"`
	// The logical operator for the condition.
	// Default: ConditionLogicalOperator.EQUALS
	//
	// Experimental.
	LogicalOperator ConditionLogicalOperator `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// The condition job state.
	// Default: - no job state is specified.
	//
	// Experimental.
	State JobState `field:"optional" json:"state" yaml:"state"`
}

