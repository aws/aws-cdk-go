package cloudformationinclude


// Construction properties of `CfnInclude`.
//
// Example:
//   parentTemplate := cfn_inc.NewCfnInclude(this, jsii.String("ParentStack"), &CfnIncludeProps{
//   	TemplateFile: jsii.String("path/to/my-parent-template.json"),
//   	LoadNestedStacks: map[string]cfnIncludeProps{
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
	// Default: - will throw an error on detecting any cyclical references.
	//
	AllowCyclicalReferences *bool `field:"optional" json:"allowCyclicalReferences" yaml:"allowCyclicalReferences"`
	// Specifies a list of LogicalIDs for resources that will be included in the CDK Stack, but will not be parsed and converted to CDK types.
	//
	// This allows you to use CFN templates
	// that rely on Intrinsic placement that `cfn-include`
	// would otherwise reject, such as non-primitive values in resource update policies.
	// Default: - All resources are hydrated.
	//
	DehydratedResources *[]*string `field:"optional" json:"dehydratedResources" yaml:"dehydratedResources"`
	// Specifies the template files that define nested stacks that should be included.
	//
	// If your template specifies a stack that isn't included here, it won't be created as a NestedStack
	// resource, and it won't be accessible from the `CfnInclude.getNestedStack` method
	// (but will still be accessible from the `CfnInclude.getResource` method).
	//
	// If you include a stack here with an ID that isn't in the template,
	// or is in the template but is not a nested stack,
	// template creation will fail and an error will be thrown.
	// Default: - no nested stacks will be included.
	//
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
	// Default: - parameters will retain their original definitions.
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Whether the resources should have the same logical IDs in the resulting CDK template as they did in the original CloudFormation template file.
	//
	// If you're vending a Construct using an existing CloudFormation template,
	// make sure to pass this as `false`.
	//
	// **Note**: regardless of whether this option is true or false,
	// the `CfnInclude.getResource` and related methods always uses the original logical ID of the resource/element,
	// as specified in the template file.
	// Default: true.
	//
	PreserveLogicalIds *bool `field:"optional" json:"preserveLogicalIds" yaml:"preserveLogicalIds"`
}

