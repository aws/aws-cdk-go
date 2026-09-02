package awstranslate


// Properties for CfnParallelDataPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnParallelDataMixinProps := &CfnParallelDataMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKey: &EncryptionKeyProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	ParallelDataConfig: &ParallelDataConfigProperty{
//   		Format: jsii.String("format"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html
//
type CfnParallelDataMixinProps struct {
	// A custom description for the parallel data resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html#cfn-translate-paralleldata-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The encryption key used to encrypt this object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html#cfn-translate-paralleldata-encryptionkey
	//
	EncryptionKey interface{} `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A custom name for the parallel data resource.
	//
	// Must be unique in the account and region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html#cfn-translate-paralleldata-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the format and S3 location of the parallel data input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html#cfn-translate-paralleldata-paralleldataconfig
	//
	ParallelDataConfig interface{} `field:"optional" json:"parallelDataConfig" yaml:"parallelDataConfig"`
	// Tags associated with the parallel data resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html#cfn-translate-paralleldata-tags
	//
	Tags *[]*CfnParallelDataPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

