package previewawshealthimagingmixins


// Properties for CfnDatastorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDatastoreMixinProps := &CfnDatastoreMixinProps{
//   	DatastoreName: jsii.String("datastoreName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthimaging-datastore.html
//
type CfnDatastoreMixinProps struct {
	// The data store name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthimaging-datastore.html#cfn-healthimaging-datastore-datastorename
	//
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The Amazon Resource Name (ARN) assigned to the Key Management Service (KMS) key for accessing encrypted data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthimaging-datastore.html#cfn-healthimaging-datastore-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The tags provided when creating a data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthimaging-datastore.html#cfn-healthimaging-datastore-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

