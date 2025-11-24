package mixinsawslogs


// This processor moves a key from one field to another. The original key is deleted.
//
// For more information about this processor including examples, see [moveKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-moveKeys) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   moveKeysProperty := &MoveKeysProperty{
//   	Entries: []interface{}{
//   		&MoveKeyEntryProperty{
//   			OverwriteIfExists: jsii.Boolean(false),
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeys.html
//
type CfnTransformerPropsMixin_MoveKeysProperty struct {
	// An array of objects, where each object contains the information about one key to move.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeys.html#cfn-logs-transformer-movekeys-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
}

