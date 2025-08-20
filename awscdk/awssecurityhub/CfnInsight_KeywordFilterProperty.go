package awssecurityhub


// A keyword filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keywordFilterProperty := &KeywordFilterProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-keywordfilter.html
//
type CfnInsight_KeywordFilterProperty struct {
	// A value for the keyword.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-keywordfilter.html#cfn-securityhub-insight-keywordfilter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

