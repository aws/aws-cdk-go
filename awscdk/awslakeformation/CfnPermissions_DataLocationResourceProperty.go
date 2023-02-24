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
//   	S3Resource: jsii.String("s3Resource"),
//   }
//
type CfnPermissions_DataLocationResourceProperty struct {
	// The identifier for the Data Catalog .
	//
	// By default, it is the account ID of the caller.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The Amazon Resource Name (ARN) that uniquely identifies the data location resource.
	S3Resource *string `field:"optional" json:"s3Resource" yaml:"s3Resource"`
}

