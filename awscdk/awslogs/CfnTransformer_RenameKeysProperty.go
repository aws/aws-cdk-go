package awslogs


// Use this processor to rename keys in a log event.
//
// For more information about this processor including examples, see [renameKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-renameKeys) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameKeysProperty := &RenameKeysProperty{
//   	Entries: []interface{}{
//   		&RenameKeyEntryProperty{
//   			Key: jsii.String("key"),
//   			RenameTo: jsii.String("renameTo"),
//
//   			// the properties below are optional
//   			OverwriteIfExists: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-renamekeys.html
//
type CfnTransformer_RenameKeysProperty struct {
	// An array of `RenameKeyEntry` objects, where each object contains the information about a single key to rename.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-renamekeys.html#cfn-logs-transformer-renamekeys-entries
	//
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
}

