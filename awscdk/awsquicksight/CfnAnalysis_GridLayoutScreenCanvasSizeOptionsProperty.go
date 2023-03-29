package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutScreenCanvasSizeOptionsProperty := &GridLayoutScreenCanvasSizeOptionsProperty{
//   	ResizeOption: jsii.String("resizeOption"),
//
//   	// the properties below are optional
//   	OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   }
//
type CfnAnalysis_GridLayoutScreenCanvasSizeOptionsProperty struct {
	// `CfnAnalysis.GridLayoutScreenCanvasSizeOptionsProperty.ResizeOption`.
	ResizeOption *string `field:"required" json:"resizeOption" yaml:"resizeOption"`
	// `CfnAnalysis.GridLayoutScreenCanvasSizeOptionsProperty.OptimizedViewPortWidth`.
	OptimizedViewPortWidth *string `field:"optional" json:"optimizedViewPortWidth" yaml:"optimizedViewPortWidth"`
}

