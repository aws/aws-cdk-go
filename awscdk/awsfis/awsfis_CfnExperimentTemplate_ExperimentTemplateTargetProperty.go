package awsfis


// Specifies a target for an experiment.
//
// You must specify at least one Amazon Resource Name (ARN) or at least one resource tag. You cannot specify both ARNs and tags.
//
// For more information, see [Targets](https://docs.aws.amazon.com/fis/latest/userguide/targets.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateTargetProperty := &experimentTemplateTargetProperty{
//   	resourceType: jsii.String("resourceType"),
//   	selectionMode: jsii.String("selectionMode"),
//
//   	// the properties below are optional
//   	filters: []interface{}{
//   		&experimentTemplateTargetFilterProperty{
//   			path: jsii.String("path"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	resourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateTargetProperty struct {
	// The resource type.
	//
	// The resource type must be supported for the specified action.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Scopes the identified resources to a specific count of the resources at random, or a percentage of the resources.
	//
	// All identified resources are included in the target.
	//
	// - ALL - Run the action on all identified targets. This is the default.
	// - COUNT(n) - Run the action on the specified number of targets, chosen from the identified targets at random. For example, COUNT(1) selects one of the targets.
	// - PERCENT(n) - Run the action on the specified percentage of targets, chosen from the identified targets at random. For example, PERCENT(25) selects 25% of the targets.
	SelectionMode *string `field:"required" json:"selectionMode" yaml:"selectionMode"`
	// The filters to apply to identify target resources using specific attributes.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The parameters for the resource type.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Names (ARNs) of the resources.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// The tags for the target resources.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

