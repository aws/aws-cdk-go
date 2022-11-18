package cloudassemblyschema


// Information about the application's runtime components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeInfo := &runtimeInfo{
//   	libraries: map[string]*string{
//   		"librariesKey": jsii.String("libraries"),
//   	},
//   }
//
// Experimental.
type RuntimeInfo struct {
	// The list of libraries loaded in the application, associated with their versions.
	// Experimental.
	Libraries *map[string]*string `field:"required" json:"libraries" yaml:"libraries"`
}

