package awslogs


// Use this processor to remove leading and trailing whitespace.
//
// For more information about this processor including examples, see [trimString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-trimString) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trimStringProperty := &TrimStringProperty{
//   	WithKeys: []*string{
//   		jsii.String("withKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-trimstring.html
//
type CfnTransformer_TrimStringProperty struct {
	// The array containing the keys of the fields to trim.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-trimstring.html#cfn-logs-transformer-trimstring-withkeys
	//
	WithKeys *[]*string `field:"required" json:"withKeys" yaml:"withKeys"`
}

