package awscloudwatch


// The query the alert runs, and the language it is written in.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryProperty := &QueryProperty{
//   	Expression: jsii.String("expression"),
//   	Language: jsii.String("language"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-query.html
//
type CfnAlert_QueryProperty struct {
	// The query expression.
	//
	// For PromQL, the series value is returned in a field named value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-query.html#cfn-cloudwatch-alert-query-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The query language.
	//
	// Valid values: SQL, PROMQL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-query.html#cfn-cloudwatch-alert-query-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
}

