package awstranslate


// Properties for CfnTerminologyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTerminologyMixinProps := &CfnTerminologyMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKey: &EncryptionKeyProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	MergeStrategy: jsii.String("mergeStrategy"),
//   	Name: jsii.String("name"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminologyData: &TerminologyDataProperty{
//   		Directionality: jsii.String("directionality"),
//   		File: jsii.String("file"),
//   		Format: jsii.String("format"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html
//
type CfnTerminologyMixinProps struct {
	// The description of the custom terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The encryption key for the custom terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-encryptionkey
	//
	EncryptionKey interface{} `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The merge strategy for the custom terminology.
	//
	// Currently only OVERWRITE is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-mergestrategy
	//
	MergeStrategy *string `field:"optional" json:"mergeStrategy" yaml:"mergeStrategy"`
	// The name of the custom terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with the terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-tags
	//
	Tags *[]*CfnTerminologyPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The terminology data for the custom terminology being imported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-terminologydata
	//
	TerminologyData interface{} `field:"optional" json:"terminologyData" yaml:"terminologyData"`
}

