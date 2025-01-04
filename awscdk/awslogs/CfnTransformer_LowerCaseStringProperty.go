package awslogs


// This processor converts a string to lowercase.
//
// For more information about this processor including examples, see [lowerCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-lowerCaseString) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lowerCaseStringProperty := &LowerCaseStringProperty{
//   	WithKeys: []*string{
//   		jsii.String("withKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-lowercasestring.html
//
type CfnTransformer_LowerCaseStringProperty struct {
	// The array caontaining the keys of the fields to convert to lowercase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-lowercasestring.html#cfn-logs-transformer-lowercasestring-withkeys
	//
	WithKeys *[]*string `field:"required" json:"withKeys" yaml:"withKeys"`
}

