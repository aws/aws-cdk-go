package awsapplicationinsights


// The `AWS::ApplicationInsights::Application LogPatternSet` property type specifies the log pattern set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logPatternSetProperty := &LogPatternSetProperty{
//   	LogPatterns: []interface{}{
//   		&LogPatternProperty{
//   			Pattern: jsii.String("pattern"),
//   			PatternName: jsii.String("patternName"),
//   			Rank: jsii.Number(123),
//   		},
//   	},
//   	PatternSetName: jsii.String("patternSetName"),
//   }
//
type CfnApplication_LogPatternSetProperty struct {
	// A list of objects that define the log patterns that belong to `LogPatternSet` .
	LogPatterns interface{} `field:"required" json:"logPatterns" yaml:"logPatterns"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 30 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	PatternSetName *string `field:"required" json:"patternSetName" yaml:"patternSetName"`
}

