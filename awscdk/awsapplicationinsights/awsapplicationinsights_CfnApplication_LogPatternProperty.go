package awsapplicationinsights


// The `AWS::ApplicationInsights::Application LogPattern` property type specifies an object that defines the log patterns that belong to a `LogPatternSet` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logPatternProperty := &logPatternProperty{
//   	pattern: jsii.String("pattern"),
//   	patternName: jsii.String("patternName"),
//   	rank: jsii.Number(123),
//   }
//
type CfnApplication_LogPatternProperty struct {
	// A regular expression that defines the log pattern.
	//
	// A log pattern can contain up to 50 characters, and it cannot be empty.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 50 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	PatternName *string `field:"required" json:"patternName" yaml:"patternName"`
	// The rank of the log pattern.
	Rank *float64 `field:"required" json:"rank" yaml:"rank"`
}

