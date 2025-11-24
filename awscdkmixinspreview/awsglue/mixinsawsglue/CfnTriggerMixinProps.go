package mixinsawsglue


// Properties for CfnTriggerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var arguments_ interface{}
//   var tags interface{}
//
//   cfnTriggerMixinProps := &CfnTriggerMixinProps{
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
//   	Description: jsii.String("description"),
//   	EventBatchingCondition: &EventBatchingConditionProperty{
//   		BatchSize: jsii.Number(123),
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
//   	Type: jsii.String("type"),
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html
//
type CfnTriggerMixinProps struct {
	// The actions initiated by this trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// A description of this trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-eventbatchingcondition
	//
	EventBatchingCondition interface{} `field:"optional" json:"eventBatchingCondition" yaml:"eventBatchingCondition"`
	// The name of the trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The predicate of this trigger, which defines when it will fire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-predicate
	//
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
	// A `cron` expression used to specify the schedule.
	//
	// For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html) in the *AWS Glue Developer Guide* . For example, to run something every day at 12:15 UTC, specify `cron(15 12 * * ? *)` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created.
	//
	// True is not supported for `ON_DEMAND` triggers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-startoncreation
	//
	StartOnCreation interface{} `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The tags to use with this trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of trigger that this is.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The name of the workflow associated with the trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-workflowname
	//
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

