package awslogs


// This processor deletes entries from a log event. These entries are key-value pairs.
//
// For more information about this processor including examples, see [deleteKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-deleteKeys) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deleteKeysProperty := &DeleteKeysProperty{
//   	WithKeys: []*string{
//   		jsii.String("withKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-deletekeys.html
//
type CfnTransformer_DeleteKeysProperty struct {
	// The list of keys to delete.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-deletekeys.html#cfn-logs-transformer-deletekeys-withkeys
	//
	WithKeys *[]*string `field:"required" json:"withKeys" yaml:"withKeys"`
}

