package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagPairProperty := &lFTagPairProperty{
//   	catalogId: jsii.String("catalogId"),
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnTagAssociation_LFTagPairProperty struct {
	// `CfnTagAssociation.LFTagPairProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.LFTagPairProperty.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `CfnTagAssociation.LFTagPairProperty.TagValues`.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

