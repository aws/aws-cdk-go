package awslogs


// Use this processor to split a field into an array of strings using a delimiting character.
//
// For more information about this processor including examples, see [splitString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-splitString) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splitStringProperty := &SplitStringProperty{
//   	Entries: []interface{}{
//   		&SplitStringEntryProperty{
//   			Delimiter: jsii.String("delimiter"),
//   			Source: jsii.String("source"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-splitstring.html
//
type CfnTransformer_SplitStringProperty struct {
	// An array of `SplitStringEntry` objects, where each object contains the information about one field to split.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-splitstring.html#cfn-logs-transformer-splitstring-entries
	//
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
}

