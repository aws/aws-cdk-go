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
//   cfnTriggerProps := &CfnTriggerProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			Arguments: arguments_,
//   			CrawlerName: jsii.String("crawlerName"),
//   			JobName: jsii.String("jobName"),
//   			NotificationProperty: &NotificationPropertyProperty{
//   				NotifyDelayAfter: jsii.Number(123),
//   			},
//   			SecurityConfiguration: jsii.String("securityConfiguration"),
//   			Timeout: jsii.Number(123),
//   		},
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EventBatchingCondition: &EventBatchingConditionProperty{
//   		BatchSize: jsii.Number(123),
//
//   		// the properties below are optional
//   		BatchWindow: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	Predicate: &PredicateProperty{
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				CrawlerName: jsii.String("crawlerName"),
//   				CrawlState: jsii.String("crawlState"),
//   				JobName: jsii.String("jobName"),
//   				LogicalOperator: jsii.String("logicalOperator"),
//   				State: jsii.String("state"),
//   			},
//   		},
//   		Logical: jsii.String("logical"),
//   	},
//   	Schedule: jsii.String("schedule"),
//   	StartOnCreation: jsii.Boolean(false),
//   	Tags: tags,
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
type CfnTriggerProps struct {
	// The actions initiated by this trigger.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The type of trigger that this is.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of this trigger.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
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

