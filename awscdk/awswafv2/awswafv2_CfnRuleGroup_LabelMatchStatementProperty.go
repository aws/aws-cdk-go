package awswafv2


// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL.
//
// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelMatchStatementProperty := &labelMatchStatementProperty{
//   	key: jsii.String("key"),
//   	scope: jsii.String("scope"),
//   }
//
type CfnRuleGroup_LabelMatchStatementProperty struct {
	// The string to match against. The setting you provide for this depends on the match statement's `Scope` setting:.
	//
	// - If the `Scope` indicates `LABEL` , then this specification must include the name and can include any number of preceding namespace specifications and prefix up to providing the fully qualified label name.
	// - If the `Scope` indicates `NAMESPACE` , then this specification can include any number of contiguous namespace strings, and can include the entire label namespace prefix from the rule group or web ACL where the label originates.
	//
	// Labels are case sensitive and components of a label must be separated by colon, for example `NS1:NS2:name` .
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specify whether you want to match using the label name or just the namespace.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

