package cloudformationinclude

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The type returned from {@link CfnInclude.getNestedStack}. Contains both the NestedStack object and CfnInclude representations of the child stack.
//
// Example:
//   var parentTemplate cfnInclude
//
//
//   includedChildStack := parentTemplate.getNestedStack(jsii.String("ChildStack"))
//   childStack := includedChildStack.stack
//   childTemplate := includedChildStack.includedTemplate
//
// Experimental.
type IncludedNestedStack struct {
	// The CfnInclude that represents the template, which can be used to access Resources and other template elements.
	// Experimental.
	IncludedTemplate CfnInclude `field:"required" json:"includedTemplate" yaml:"includedTemplate"`
	// The NestedStack object which represents the scope of the template.
	// Experimental.
	Stack awscdk.NestedStack `field:"required" json:"stack" yaml:"stack"`
}

