package previewawsgluemixins


// A structure that describes a target table for resource linking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableIdentifierProperty := &TableIdentifierProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableidentifier.html
//
type CfnTablePropsMixin_TableIdentifierProperty struct {
	// The ID of the Data Catalog in which the table resides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableidentifier.html#cfn-glue-table-tableidentifier-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database that contains the target table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableidentifier.html#cfn-glue-table-tableidentifier-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the target table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableidentifier.html#cfn-glue-table-tableidentifier-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Region of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableidentifier.html#cfn-glue-table-tableidentifier-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

