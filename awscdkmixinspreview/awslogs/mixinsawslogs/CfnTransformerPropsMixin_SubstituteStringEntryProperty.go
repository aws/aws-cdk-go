package mixinsawslogs


// This object defines one log field key that will be replaced using the [substituteString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-substituteString) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   substituteStringEntryProperty := &SubstituteStringEntryProperty{
//   	From: jsii.String("from"),
//   	Source: jsii.String("source"),
//   	To: jsii.String("to"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html
//
type CfnTransformerPropsMixin_SubstituteStringEntryProperty struct {
	// The regular expression string to be replaced.
	//
	// Special regex characters such as [ and ] must be escaped using \\ when using double quotes and with \ when using single quotes. For more information, see [Class Pattern](https://docs.aws.amazon.com/https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/regex/Pattern.html) on the Oracle web site.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html#cfn-logs-transformer-substitutestringentry-from
	//
	From *string `field:"optional" json:"from" yaml:"from"`
	// The key to modify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html#cfn-logs-transformer-substitutestringentry-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The string to be substituted for each match of `from`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestringentry.html#cfn-logs-transformer-substitutestringentry-to
	//
	To *string `field:"optional" json:"to" yaml:"to"`
}

