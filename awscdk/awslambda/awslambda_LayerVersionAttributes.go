package awslambda


// Properties necessary to import a LayerVersion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var runtime runtime
//
//   layerVersionAttributes := &layerVersionAttributes{
//   	layerVersionArn: jsii.String("layerVersionArn"),
//
//   	// the properties below are optional
//   	compatibleRuntimes: []*runtime{
//   		runtime,
//   	},
//   }
//
type LayerVersionAttributes struct {
	// The ARN of the LayerVersion.
	LayerVersionArn *string `field:"required" json:"layerVersionArn" yaml:"layerVersionArn"`
	// The list of compatible runtimes with this Layer.
	CompatibleRuntimes *[]Runtime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
}

