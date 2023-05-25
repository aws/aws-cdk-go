package awslakeformation


// A structure containing an LF-tag key and values for a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagKeyResourceProperty := &LFTagKeyResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	TagKey: jsii.String("tagKey"),
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnPrincipalPermissions_LFTagKeyResourceProperty struct {
	// The identifier for the Data Catalog where the location is registered with Data Catalog .
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The key-name for the LF-tag.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// A list of possible values for the corresponding `TagKey` of an LF-tag key-value pair.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

