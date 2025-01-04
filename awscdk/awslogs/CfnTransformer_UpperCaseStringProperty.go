package awslogs


// This processor converts a string field to uppercase.
//
// For more information about this processor including examples, see [upperCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-upperCaseString) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   upperCaseStringProperty := &UpperCaseStringProperty{
//   	WithKeys: []*string{
//   		jsii.String("withKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-uppercasestring.html
//
type CfnTransformer_UpperCaseStringProperty struct {
	// The array of containing the keys of the field to convert to uppercase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-uppercasestring.html#cfn-logs-transformer-uppercasestring-withkeys
	//
	WithKeys *[]*string `field:"required" json:"withKeys" yaml:"withKeys"`
}

