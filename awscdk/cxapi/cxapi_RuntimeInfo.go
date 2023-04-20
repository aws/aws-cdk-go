package cxapi


// Backwards compatibility for when `RuntimeInfo` was defined here.
//
// This is necessary because its used as an input in the stable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeInfo := &RuntimeInfo{
//   	Libraries: map[string]*string{
//   		"librariesKey": jsii.String("libraries"),
//   	},
//   }
//
// See: core.ConstructNode.synth
//
// Deprecated: moved to package 'cloud-assembly-schema'.
type RuntimeInfo struct {
	// The list of libraries loaded in the application, associated with their versions.
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Libraries *map[string]*string `field:"required" json:"libraries" yaml:"libraries"`
}

