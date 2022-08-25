package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLocationResourceProperty := &dataLocationResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnPrincipalPermissions_DataLocationResourceProperty struct {
	// `CfnPrincipalPermissions.DataLocationResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.DataLocationResourceProperty.ResourceArn`.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

