package previewawsapplicationinsightsmixins


// The `AWS::ApplicationInsights::Application LogPattern` property type specifies an object that defines the log patterns that belong to a `LogPatternSet` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logPatternProperty := &LogPatternProperty{
//   	Pattern: jsii.String("pattern"),
//   	PatternName: jsii.String("patternName"),
//   	Rank: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpattern.html
//
type CfnApplicationPropsMixin_LogPatternProperty struct {
	// A regular expression that defines the log pattern.
	//
	// A log pattern can contain up to 50 characters, and it cannot be empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpattern.html#cfn-applicationinsights-application-logpattern-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 50 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpattern.html#cfn-applicationinsights-application-logpattern-patternname
	//
	PatternName *string `field:"optional" json:"patternName" yaml:"patternName"`
	// The rank of the log pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpattern.html#cfn-applicationinsights-application-logpattern-rank
	//
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

