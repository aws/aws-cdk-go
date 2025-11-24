package mixinsawsglue


// Defines a condition under which a trigger fires.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionProperty := &ConditionProperty{
//   	CrawlerName: jsii.String("crawlerName"),
//   	CrawlState: jsii.String("crawlState"),
//   	JobName: jsii.String("jobName"),
//   	LogicalOperator: jsii.String("logicalOperator"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html
//
type CfnTriggerPropsMixin_ConditionProperty struct {
	// The name of the crawler to which this condition applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-crawlername
	//
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// The state of the crawler to which this condition applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-crawlstate
	//
	CrawlState *string `field:"optional" json:"crawlState" yaml:"crawlState"`
	// The name of the job whose `JobRuns` this condition applies to, and on which this trigger waits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-jobname
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// A logical operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-logicaloperator
	//
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// The condition state.
	//
	// Currently, the values supported are `SUCCEEDED` , `STOPPED` , `TIMEOUT` , and `FAILED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

