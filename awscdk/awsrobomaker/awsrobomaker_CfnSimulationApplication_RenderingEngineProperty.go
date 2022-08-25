package awsrobomaker


// Information about a rendering engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renderingEngineProperty := &renderingEngineProperty{
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnSimulationApplication_RenderingEngineProperty struct {
	// The name of the rendering engine.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the rendering engine.
	Version *string `field:"required" json:"version" yaml:"version"`
}

