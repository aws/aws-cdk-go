package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tagsItemsProperty := &TagsItemsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-tagsitems.html
//
type CfnLegalHoldPropsMixin_TagsItemsProperty struct {
	// The key name of the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-tagsitems.html#cfn-backup-legalhold-tagsitems-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-tagsitems.html#cfn-backup-legalhold-tagsitems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

