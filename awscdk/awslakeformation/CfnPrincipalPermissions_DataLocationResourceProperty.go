package awslakeformation


// A structure for a data location object where permissions are granted or revoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLocationResourceProperty := &DataLocationResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type CfnPrincipalPermissions_DataLocationResourceProperty struct {
	// The identifier for the Data Catalog where the location is registered with AWS Lake Formation .
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The Amazon Resource Name (ARN) that uniquely identifies the data location resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

