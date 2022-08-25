package awsbackup


// A tag condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagCondition := &tagCondition{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	operation: awscdk.Aws_backup.tagOperation_STRING_EQUALS,
//   }
//
type TagCondition struct {
	// The key in a key-value pair.
	//
	// For example, in `"ec2:ResourceTag/Department": "accounting"`,
	// `ec2:ResourceTag/Department` is the key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value in a key-value pair.
	//
	// For example, in `"ec2:ResourceTag/Department": "accounting"`,
	// `accounting` is the value.
	Value *string `field:"required" json:"value" yaml:"value"`
	// An operation that is applied to a key-value pair used to filter resources in a selection.
	Operation TagOperation `field:"optional" json:"operation" yaml:"operation"`
}

