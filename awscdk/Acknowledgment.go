package awscdk


// An acknowledgment of a validation rule, used to suppress it from output.
//
// Example:
//   awscdk.Annotations_Of(this).AcknowledgeWarning(jsii.String("IAM:Group:MaxPoliciesExceeded"), jsii.String("Account has quota increased to 20"))
//
//   // Because Annotations ultimately become Validations, you can also acknowledge the Validation
//   awscdk.Validations_Of(this).Acknowledge(&Acknowledgment{
//   	Id: jsii.String("Construct-Annotations::IAM:Group:MaxPoliciesExceeded"),
//   	Reason: jsii.String("Account has quota increased to 20"),
//   })
//
type Acknowledgment struct {
	// The rule ID to acknowledge.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The reason for acknowledging this rule.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Optionally narrow the acknowledgement to only the warning occurrences whose call-site context matches this filter.
	//
	// The object is treated as an **`objectLike` partial match** against the
	// context attached to each occurrence (the same semantics as
	// `template.hasResourceProperties`): every key you specify must be present
	// with a matching value, and any key you do not mention is left unconstrained.
	// A warning emitted without any context is never matched by a non-empty
	// filter, so a specific filter cannot accidentally silence an undetailed
	// warning.
	//
	// A value is matched for equality (including non-string values such as
	// numbers, booleans, or arrays). Because the value type is `any`, richer
	// per-key matching can be introduced later without a breaking change — a
	// future release may accept a matcher as a value (e.g. a glob/regex) in
	// addition to a literal, leaving existing literal callers untouched.
	//
	// When omitted, all occurrences of the rule id are acknowledged
	// (backwards-compatible behavior).
	// Default: - acknowledges all occurrences of the rule id.
	//
	Where *map[string]interface{} `field:"optional" json:"where" yaml:"where"`
}

