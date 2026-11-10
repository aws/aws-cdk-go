package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsItemsProperty := &TagsItemsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-snapshot-tagsitems.html
//
type CfnSnapshot_TagsItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-snapshot-tagsitems.html#cfn-ec2-snapshot-tagsitems-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-snapshot-tagsitems.html#cfn-ec2-snapshot-tagsitems-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

