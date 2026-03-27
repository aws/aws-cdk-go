package awsglue


// A structure that describes a target catalog for resource linking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetRedshiftCatalogProperty := &TargetRedshiftCatalogProperty{
//   	CatalogArn: jsii.String("catalogArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-targetredshiftcatalog.html
//
type CfnCatalog_TargetRedshiftCatalogProperty struct {
	// The Amazon Resource Name (ARN) of the catalog resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-targetredshiftcatalog.html#cfn-glue-catalog-targetredshiftcatalog-catalogarn
	//
	CatalogArn *string `field:"required" json:"catalogArn" yaml:"catalogArn"`
}

