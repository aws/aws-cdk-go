package awsquicksight


// Configuration options for the canvas of a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutCanvasSizeOptionsProperty := &GridLayoutCanvasSizeOptionsProperty{
//   	ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   		ResizeOption: jsii.String("resizeOption"),
//
//   		// the properties below are optional
//   		OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   	},
//   }
//
type CfnAnalysis_GridLayoutCanvasSizeOptionsProperty struct {
	// The options that determine the sizing of the canvas used in a grid layout.
	ScreenCanvasSizeOptions interface{} `field:"optional" json:"screenCanvasSizeOptions" yaml:"screenCanvasSizeOptions"`
}

