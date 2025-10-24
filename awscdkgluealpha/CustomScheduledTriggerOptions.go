package awscdkgluealpha


// Properties for configuring a custom-scheduled Glue Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnCrawler CfnCrawler
//   var job Job
//   var securityConfiguration SecurityConfiguration
//   var triggerSchedule TriggerSchedule
//
//   customScheduledTriggerOptions := &CustomScheduledTriggerOptions{
//   	Actions: []Action{
//   		&Action{
//   			Arguments: map[string]*string{
//   				"argumentsKey": jsii.String("arguments"),
//   			},
//   			Crawler: cfnCrawler,
//   			Job: job,
//   			SecurityConfiguration: securityConfiguration,
//   			Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   		},
//   	},
//   	Schedule: triggerSchedule,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StartOnCreation: jsii.Boolean(false),
//   }
//
// Experimental.
type CustomScheduledTriggerOptions struct {
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
	// The custom schedule for the trigger.
	// Experimental.
	Schedule TriggerSchedule `field:"required" json:"schedule" yaml:"schedule"`
}

