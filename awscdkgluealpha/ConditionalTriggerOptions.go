package awscdkgluealpha


// Properties for configuring a Condition (Predicate) based Glue Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnCrawler cfnCrawler
//   var job job
//   var securityConfiguration securityConfiguration
//
//   conditionalTriggerOptions := &ConditionalTriggerOptions{
//   	Actions: []action{
//   		&action{
//   			Arguments: map[string]*string{
//   				"argumentsKey": jsii.String("arguments"),
//   			},
//   			Crawler: cfnCrawler,
//   			Job: job,
//   			SecurityConfiguration: securityConfiguration,
//   			Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   		},
//   	},
//   	Predicate: &Predicate{
//   		Conditions: []condition{
//   			&condition{
//   				CrawlerName: jsii.String("crawlerName"),
//   				CrawlState: glue_alpha.CrawlerState_RUNNING,
//   				Job: job,
//   				LogicalOperator: glue_alpha.ConditionLogicalOperator_EQUALS,
//   				State: glue_alpha.JobState_SUCCEEDED,
//   			},
//   		},
//   		Logical: glue_alpha.PredicateLogical_AND,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StartOnCreation: jsii.Boolean(false),
//   }
//
// Experimental.
type ConditionalTriggerOptions struct {
	// The actions initiated by this trigger.
	// Experimental.
	Actions *[]*Action `field:"required" json:"actions" yaml:"actions"`
	// A description for the trigger.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the trigger.
	// Default: - no name is provided.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether to start the trigger on creation or not.
	// Default: - false.
	//
	// Experimental.
	StartOnCreation *bool `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The predicate for the trigger.
	// Experimental.
	Predicate *Predicate `field:"required" json:"predicate" yaml:"predicate"`
}

