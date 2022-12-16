package awslakeformation


// A structure for a data location object where permissions are granted or revoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLocationResourceProperty := &dataLocationResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	s3Resource: jsii.String("s3Resource"),
//   }
//
type CfnPermissions_DataLocationResourceProperty struct {
	// `CfnPermissions.DataLocationResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Currently not supported by AWS CloudFormation .
	S3Resource *string `field:"optional" json:"s3Resource" yaml:"s3Resource"`
}

