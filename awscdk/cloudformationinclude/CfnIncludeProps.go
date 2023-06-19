package cloudformationinclude


// Construction properties of {@link CfnInclude}.
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
// Experimental.
type CfnIncludeProps struct {
	// Path to the template file.
	//
	// Both JSON and YAML template formats are supported.
	// Experimental.
	TemplateFile *string `field:"required" json:"templateFile" yaml:"templateFile"`
	// Specifies the template files that define nested stacks that should be included.
	//
	// If your template specifies a stack that isn't included here, it won't be created as a NestedStack
	// resource, and it won't be accessible from the {@link CfnInclude.getNestedStack} method
	// (but will still be accessible from the {@link CfnInclude.getResource} method).
	//
	// If you include a stack here with an ID that isn't in the template,
	// or is in the template but is not a nested stack,
	// template creation will fail and an error will be thrown.
	// Experimental.
	LoadNestedStacks *map[string]*CfnIncludeProps `field:"optional" json:"loadNestedStacks" yaml:"loadNestedStacks"`
	// Specifies parameters to be replaced by the values in this mapping.
	//
	// Any parameters in the template that aren't specified here will be left unmodified.
	// If you include a parameter here with an ID that isn't in the template,
	// template creation will fail and an error will be thrown.
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Whether the resources should have the same logical IDs in the resulting CDK template as they did in the original CloudFormation template file.
	//
	// If you're vending a Construct using an existing CloudFormation template,
	// make sure to pass this as `false`.
	//
	// **Note**: regardless of whether this option is true or false,
	// the {@link CfnInclude.getResource} and related methods always uses the original logical ID of the resource/element,
	// as specified in the template file.
	// Experimental.
	PreserveLogicalIds *bool `field:"optional" json:"preserveLogicalIds" yaml:"preserveLogicalIds"`
}

