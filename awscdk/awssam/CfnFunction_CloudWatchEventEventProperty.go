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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatcheventevent.html
//
type CfnFunction_CloudWatchEventEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatcheventevent.html#cfn-serverless-function-cloudwatcheventevent-pattern
	//
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatcheventevent.html#cfn-serverless-function-cloudwatcheventevent-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatcheventevent.html#cfn-serverless-function-cloudwatcheventevent-inputpath
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

