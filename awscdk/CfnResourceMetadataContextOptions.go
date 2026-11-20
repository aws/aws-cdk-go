package awscdk


// Options for `CfnResourceMetadataContext.of()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var constructSelector IConstructSelector
//
//   cfnResourceMetadataContextOptions := &CfnResourceMetadataContextOptions{
//   	Selector: constructSelector,
//   }
//
type CfnResourceMetadataContextOptions struct {
	// Selects the constructs targeted by `set()` or `clear()`.
	//
	// Selection runs during synthesis and only selected `CfnResource` constructs
	// are updated. Synthesis fails when no `CfnResource` is selected. The default
	// selects a `CfnResource` scope or its immediate `CfnResource` default child.
	// Default: ConstructSelector.cfnResource()
	//
	Selector IConstructSelector `field:"optional" json:"selector" yaml:"selector"`
}

