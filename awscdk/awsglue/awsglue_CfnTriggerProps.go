package awsglue


// Properties for defining a `CfnTrigger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var arguments_ interface{}
//   var tags interface{}
//
//   cfnTriggerProps := &cfnTriggerProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			arguments: arguments_,
//   			crawlerName: jsii.String("crawlerName"),
//   			jobName: jsii.String("jobName"),
//   			notificationProperty: &notificationPropertyProperty{
//   				notifyDelayAfter: jsii.Number(123),
//   			},
//   			securityConfiguration: jsii.String("securityConfiguration"),
//   			timeout: jsii.Number(123),
//   		},
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	eventBatchingCondition: &eventBatchingConditionProperty{
//   		batchSize: jsii.Number(123),
//
//   		// the properties below are optional
//   		batchWindow: jsii.Number(123),
//   	},
//   	name: jsii.String("name"),
//   	predicate: &predicateProperty{
//   		conditions: []interface{}{
//   			&conditionProperty{
//   				crawlerName: jsii.String("crawlerName"),
//   				crawlState: jsii.String("crawlState"),
//   				jobName: jsii.String("jobName"),
//   				logicalOperator: jsii.String("logicalOperator"),
//   				state: jsii.String("state"),
//   			},
//   		},
//   		logical: jsii.String("logical"),
//   	},
//   	schedule: jsii.String("schedule"),
//   	startOnCreation: jsii.Boolean(false),
//   	tags: tags,
//   	workflowName: jsii.String("workflowName"),
//   }
//
type CfnTriggerProps struct {
	// The actions initiated by this trigger.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The type of trigger that this is.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of this trigger.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Glue::Trigger.EventBatchingCondition`.
	EventBatchingCondition interface{} `field:"optional" json:"eventBatchingCondition" yaml:"eventBatchingCondition"`
	// The name of the trigger.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The predicate of this trigger, which defines when it will fire.
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
	// A `cron` expression used to specify the schedule.
	//
	// For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html) in the *AWS Glue Developer Guide* . For example, to run something every day at 12:15 UTC, specify `cron(15 12 * * ? *)` .
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created.
	//
	// True is not supported for `ON_DEMAND` triggers.
	StartOnCreation interface{} `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The tags to use with this trigger.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workflow associated with the trigger.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

