package awsbackup


// Pair of two related strings.
//
// Allowed characters are letters, white space, and numbers that can be represented in UTF-8 and the following characters: `+ - = . _ : /`
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValueProperty := &KeyValueProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingselection-keyvalue.html
//
type CfnRestoreTestingSelection_KeyValueProperty struct {
	// The tag key (String). The key can't start with `aws:` .
	//
	// Length Constraints: Minimum length of 1. Maximum length of 128.
	//
	// Pattern: `^(?![aA]{1}[wW]{1}[sS]{1}:)([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingselection-keyvalue.html#cfn-backup-restoretestingselection-keyvalue-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the key.
	//
	// Length Constraints: Maximum length of 256.
	//
	// Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingselection-keyvalue.html#cfn-backup-restoretestingselection-keyvalue-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

