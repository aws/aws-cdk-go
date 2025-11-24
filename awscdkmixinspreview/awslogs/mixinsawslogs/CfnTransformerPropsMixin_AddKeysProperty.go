package mixinsawslogs


// This processor adds new key-value pairs to the log event.
//
// For more information about this processor including examples, see [addKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-addKeys) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   addKeysProperty := &AddKeysProperty{
//   	Entries: []interface{}{
//   		&AddKeyEntryProperty{
//   			Key: jsii.String("key"),
//   			OverwriteIfExists: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-addkeys.html
//
type CfnTransformerPropsMixin_AddKeysProperty struct {
	// An array of objects, where each object contains the information about one key to add to the log event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-addkeys.html#cfn-logs-transformer-addkeys-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
}

