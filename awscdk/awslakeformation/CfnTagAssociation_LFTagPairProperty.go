package awslakeformation


// A structure containing the catalog ID, tag key, and tag values of an LF-tag key-value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagPairProperty := &LFTagPairProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	TagKey: jsii.String("tagKey"),
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnTagAssociation_LFTagPairProperty struct {
	// The identifier for the Data Catalog .
	//
	// By default, it is the account ID of the caller.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The key-name for the LF-tag.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

