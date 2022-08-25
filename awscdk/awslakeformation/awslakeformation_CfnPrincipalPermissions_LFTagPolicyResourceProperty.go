package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagPolicyResourceProperty := &lFTagPolicyResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	expression: []interface{}{
//   		&lFTagProperty{
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	resourceType: jsii.String("resourceType"),
//   }
//
type CfnPrincipalPermissions_LFTagPolicyResourceProperty struct {
	// `CfnPrincipalPermissions.LFTagPolicyResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.LFTagPolicyResourceProperty.Expression`.
	Expression interface{} `field:"required" json:"expression" yaml:"expression"`
	// `CfnPrincipalPermissions.LFTagPolicyResourceProperty.ResourceType`.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

