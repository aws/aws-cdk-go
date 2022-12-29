package cloudformationinclude


// Construction properties of {@link CfnInclude}.
//
// Example:
//   parentTemplate := cfn_inc.NewCfnInclude(this, jsii.String("ParentStack"), &cfnIncludeProps{
//   	templateFile: jsii.String("path/to/my-parent-template.json"),
//   	loadNestedStacks: map[string]*cfnIncludeProps{
//   		"ChildStack": &cfnIncludeProps{
//   			"templateFile": jsii.String("path/to/my-nested-template.json"),
//   		},
//   	},
//   })
//
type CfnIncludeProps struct {
	// Path to the template file.
	//
	// Both JSON and YAML template formats are supported.
	TemplateFile *string `field:"required" json:"templateFile" yaml:"templateFile"`
	// Specifies whether to allow cyclical references, effectively disregarding safeguards meant to avoid undeployable templates.
	//
	// This should only be set to true in the case of templates utilizing cloud transforms (e.g. SAM) that
	// after processing the transform will no longer contain any circular references.
	AllowCyclicalReferences *bool `field:"optional" json:"allowCyclicalReferences" yaml:"allowCyclicalReferences"`
	// Specifies the template files that define nested stacks that should be included.
	//
	// If your template specifies a stack that isn't included here, it won't be created as a NestedStack
	// resource, and it won't be accessible from the {@link CfnInclude.getNestedStack} method
	// (but will still be accessible from the {@link CfnInclude.getResource} method).
	//
	// If you include a stack here with an ID that isn't in the template,
	// or is in the template but is not a nested stack,
	// template creation will fail and an error will be thrown.
	LoadNestedStacks *map[string]*CfnIncludeProps `field:"optional" json:"loadNestedStacks" yaml:"loadNestedStacks"`
	// Specifies parameters to be replaced by the values in this mapping.
	//
	// Any parameters in the template that aren't specified here will be left unmodified.
	// If you include a parameter here with an ID that isn't in the template,
	// template creation will fail and an error will be thrown.
	//
	// If you are importing a parameter from a live stack, we cannot know the value of that
	// parameter. You will need to supply a value for your parameters, else the default
	// value will be used.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Whether the resources should have the same logical IDs in the resulting CDK template as they did in the original CloudFormation template file.
	//
	// If you're vending a Construct using an existing CloudFormation template,
	// make sure to pass this as `false`.
	//
	// **Note**: regardless of whether this option is true or false,
	// the {@link CfnInclude.getResource} and related methods always uses the original logical ID of the resource/element,
	// as specified in the template file.
	PreserveLogicalIds *bool `field:"optional" json:"preserveLogicalIds" yaml:"preserveLogicalIds"`
}

