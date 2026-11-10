package awscdk


// Options for adding an acknowledgeable warning via `Annotations.addWarningV2` or `Validations.addWarning`.
//
// Example:
//   // The construct attaches context describing this specific occurrence:
//   awscdk.Validations_Of(this).AddWarning(jsii.String("my-lib:unsupportedActionType"), jsii.String("This action type is not supported and will be ignored"), &WarningOptions{
//   	Context: map[string]interface{}{
//   		"service": jsii.String("aiops"),
//   		"resource": jsii.String("investigation-group"),
//   	},
//   })
//
//   // Acknowledge only the occurrences that match. `where` is an objectLike partial
//   // match against the context: any key you do not mention is left unconstrained,
//   // so the filter below also suppresses an occurrence whose context is
//   // `{ service: 'aiops', resource: 'other' }`:
//   awscdk.Validations_Of(this).Acknowledge(&Acknowledgment{
//   	Id: jsii.String("my-lib:unsupportedActionType"),
//   	Reason: jsii.String("aiops actions are attached intentionally"),
//   	Where: map[string]interface{}{
//   		"service": jsii.String("aiops"),
//   	},
//   })
//
type WarningOptions struct {
	// Structured context describing this specific warning occurrence.
	//
	// The construct chooses the keys and values (e.g. `{ service: 'aiops' }`, or
	// richer values such as numbers or arrays). A later `acknowledge({ where })`
	// is matched against this context as an `objectLike` partial match, to
	// selectively suppress only the matching occurrences of the warning id while
	// the rest continue to warn. When omitted, the warning carries no context and
	// can only be suppressed by an unfiltered acknowledgement.
	// Default: - no context is attached.
	//
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
}

