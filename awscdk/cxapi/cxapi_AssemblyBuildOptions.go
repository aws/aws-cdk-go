package cxapi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assemblyBuildOptions := &AssemblyBuildOptions{
//   	RuntimeInfo: &RuntimeInfo{
//   		Libraries: map[string]*string{
//   			"librariesKey": jsii.String("libraries"),
//   		},
//   	},
//   }
//
// Experimental.
type AssemblyBuildOptions struct {
	// Include the specified runtime information (module versions) in manifest.
	// Deprecated: All template modifications that should result from this should
	// have already been inserted into the template.
	RuntimeInfo *RuntimeInfo `field:"optional" json:"runtimeInfo" yaml:"runtimeInfo"`
}

