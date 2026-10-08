package awstranslate


// Properties for defining a `CfnTerminology`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTerminologyProps := &CfnTerminologyProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EncryptionKey: &EncryptionKeyProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	MergeStrategy: jsii.String("mergeStrategy"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminologyData: &TerminologyDataProperty{
//   		File: jsii.String("file"),
//   		Format: jsii.String("format"),
//
//   		// the properties below are optional
//   		Directionality: jsii.String("directionality"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html
//
type CfnTerminologyProps struct {
	// The name of the custom terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
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
	// Tags associated with the terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-tags
	//
	Tags *[]*CfnTerminology_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The terminology data for the custom terminology being imported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html#cfn-translate-terminology-terminologydata
	//
	TerminologyData interface{} `field:"optional" json:"terminologyData" yaml:"terminologyData"`
}

