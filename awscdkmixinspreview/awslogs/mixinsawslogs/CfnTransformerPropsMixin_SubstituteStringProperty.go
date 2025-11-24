package mixinsawslogs


// This processor matches a key’s value against a regular expression and replaces all matches with a replacement string.
//
// For more information about this processor including examples, see [substituteString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-substituteString) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   substituteStringProperty := &SubstituteStringProperty{
//   	Entries: []interface{}{
//   		&SubstituteStringEntryProperty{
//   			From: jsii.String("from"),
//   			Source: jsii.String("source"),
//   			To: jsii.String("to"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestring.html
//
type CfnTransformerPropsMixin_SubstituteStringProperty struct {
	// An array of objects, where each object contains the information about one key to match and replace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-substitutestring.html#cfn-logs-transformer-substitutestring-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
}

