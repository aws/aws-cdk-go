package awslogs


// This processor copies values within a log event.
//
// You can also use this processor to add metadata to log events by copying the values of the following metadata keys into the log events: `@logGroupName` , `@logGroupStream` , `@accountId` , `@regionName` .
//
// For more information about this processor including examples, see [copyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-copyValue) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyValueProperty := &CopyValueProperty{
//   	Entries: []interface{}{
//   		&CopyValueEntryProperty{
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			OverwriteIfExists: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-copyvalue.html
//
type CfnTransformer_CopyValueProperty struct {
	// An array of `CopyValueEntry` objects, where each object contains the information about one field value to copy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-copyvalue.html#cfn-logs-transformer-copyvalue-entries
	//
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
}

