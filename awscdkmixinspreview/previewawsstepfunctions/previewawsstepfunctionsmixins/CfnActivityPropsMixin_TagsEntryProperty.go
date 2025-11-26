package previewawsstepfunctionsmixins


// The `TagsEntry` property specifies *tags* to identify an activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagsEntryProperty := &TagsEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-tagsentry.html
//
type CfnActivityPropsMixin_TagsEntryProperty struct {
	// The `key` for a key-value pair in a tag entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-tagsentry.html#cfn-stepfunctions-activity-tagsentry-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The `value` for a key-value pair in a tag entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-tagsentry.html#cfn-stepfunctions-activity-tagsentry-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

