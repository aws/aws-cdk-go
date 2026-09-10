package awscdkgluealpha


// Base options for triggers that can be started when they are created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var action Action
//
//   startableTriggerOptions := &StartableTriggerOptions{
//   	Actions: []Action{
//   		action,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StartOnCreation: jsii.Boolean(false),
//   }
//
// Experimental.
type StartableTriggerOptions struct {
	// The actions initiated by this trigger.
	// Experimental.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
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
}

