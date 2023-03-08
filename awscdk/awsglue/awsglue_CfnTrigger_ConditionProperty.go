package awsglue


// Defines a condition under which a trigger fires.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &conditionProperty{
//   	crawlerName: jsii.String("crawlerName"),
//   	crawlState: jsii.String("crawlState"),
//   	jobName: jsii.String("jobName"),
//   	logicalOperator: jsii.String("logicalOperator"),
//   	state: jsii.String("state"),
//   }
//
type CfnTrigger_ConditionProperty struct {
	// The name of the crawler to which this condition applies.
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// The state of the crawler to which this condition applies.
	CrawlState *string `field:"optional" json:"crawlState" yaml:"crawlState"`
	// The name of the job whose `JobRuns` this condition applies to, and on which this trigger waits.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// A logical operator.
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// The condition state.
	//
	// Currently, the values supported are `SUCCEEDED` , `STOPPED` , `TIMEOUT` , and `FAILED` .
	State *string `field:"optional" json:"state" yaml:"state"`
}

