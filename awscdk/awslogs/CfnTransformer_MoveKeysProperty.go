package awslogs


// This processor moves a key from one field to another. The original key is deleted.
//
// For more information about this processor including examples, see [moveKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-moveKeys) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   moveKeysProperty := &MoveKeysProperty{
//   	Entries: []interface{}{
//   		&MoveKeyEntryProperty{
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			OverwriteIfExists: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeys.html
//
type CfnTransformer_MoveKeysProperty struct {
	// An array of objects, where each object contains the information about one key to move.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeys.html#cfn-logs-transformer-movekeys-entries
	//
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
}

