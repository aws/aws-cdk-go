package awsfis


// Specifies an action for an experiment template.
//
// For more information, see [Actions](https://docs.aws.amazon.com/fis/latest/userguide/actions.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateActionProperty := &experimentTemplateActionProperty{
//   	actionId: jsii.String("actionId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	startAfter: []*string{
//   		jsii.String("startAfter"),
//   	},
//   	targets: map[string]*string{
//   		"targetsKey": jsii.String("targets"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateActionProperty struct {
	// The ID of the action.
	//
	// The format of the action ID is: aws: *service-name* : *action-type* .
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// A description for the action.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameters for the action, if applicable.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The name of the action that must be completed before the current action starts.
	//
	// Omit this parameter to run the action at the start of the experiment.
	StartAfter *[]*string `field:"optional" json:"startAfter" yaml:"startAfter"`
	// The targets for the action.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

