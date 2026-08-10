package awsomics


// Properties for CfnReferencePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnReferenceMixinProps := &CfnReferenceMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ReferenceStoreId: jsii.String("referenceStoreId"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-reference.html
//
type CfnReferenceMixinProps struct {
	// The reference's description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-reference.html#cfn-omics-reference-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The reference's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-reference.html#cfn-omics-reference-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The reference's reference store ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-reference.html#cfn-omics-reference-referencestoreid
	//
	ReferenceStoreId *string `field:"optional" json:"referenceStoreId" yaml:"referenceStoreId"`
	// Tags for the reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-reference.html#cfn-omics-reference-tags
	//
	Tags *[]*CfnReferencePropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

