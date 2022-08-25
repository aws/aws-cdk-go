package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagKeyResourceProperty := &lFTagKeyResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnPrincipalPermissions_LFTagKeyResourceProperty struct {
	// `CfnPrincipalPermissions.LFTagKeyResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.LFTagKeyResourceProperty.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `CfnPrincipalPermissions.LFTagKeyResourceProperty.TagValues`.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

