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
//   experimentTemplateActionProperty := &ExperimentTemplateActionProperty{
//   	ActionId: jsii.String("actionId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	StartAfter: []*string{
//   		jsii.String("startAfter"),
//   	},
//   	Targets: map[string]*string{
//   		"targetsKey": jsii.String("targets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html
//
type CfnExperimentTemplate_ExperimentTemplateActionProperty struct {
	// The ID of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-actionid
	//
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// A description for the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameters for the action, if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The name of the action that must be completed before the current action starts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-startafter
	//
	StartAfter *[]*string `field:"optional" json:"startAfter" yaml:"startAfter"`
	// One or more targets for the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

