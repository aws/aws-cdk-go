package mixinsawsapplicationinsights


// The `AWS::ApplicationInsights::Application LogPatternSet` property type specifies the log pattern set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpatternset.html
//
type CfnApplicationPropsMixin_LogPatternSetProperty struct {
	// A list of objects that define the log patterns that belong to `LogPatternSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpatternset.html#cfn-applicationinsights-application-logpatternset-logpatterns
	//
	LogPatterns interface{} `field:"optional" json:"logPatterns" yaml:"logPatterns"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 30 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-logpatternset.html#cfn-applicationinsights-application-logpatternset-patternsetname
	//
	PatternSetName *string `field:"optional" json:"patternSetName" yaml:"patternSetName"`
}

