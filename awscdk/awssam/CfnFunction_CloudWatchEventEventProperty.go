package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   cloudWatchEventEventProperty := &CloudWatchEventEventProperty{
//   	Pattern: pattern,
//
//   	// the properties below are optional
//   	Input: jsii.String("input"),
//   	InputPath: jsii.String("inputPath"),
//   }
//
type CfnFunction_CloudWatchEventEventProperty struct {
	// `CfnFunction.CloudWatchEventEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnFunction.CloudWatchEventEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnFunction.CloudWatchEventEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

