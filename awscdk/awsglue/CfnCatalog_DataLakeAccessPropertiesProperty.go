package awsglue


// Data lake access properties for the catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakeAccessPropertiesProperty := &DataLakeAccessPropertiesProperty{
//   	AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   	CatalogType: jsii.String("catalogType"),
//   	DataLakeAccess: jsii.Boolean(false),
//   	DataTransferRole: jsii.String("dataTransferRole"),
//   	KmsKey: jsii.String("kmsKey"),
//   	ManagedWorkgroupName: jsii.String("managedWorkgroupName"),
//   	ManagedWorkgroupStatus: jsii.String("managedWorkgroupStatus"),
//   	RedshiftDatabaseName: jsii.String("redshiftDatabaseName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html
//
type CfnCatalog_DataLakeAccessPropertiesProperty struct {
	// Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-allowfulltableexternaldataaccess
	//
	AllowFullTableExternalDataAccess *string `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// Specifies a federated catalog type for the native catalog resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-catalogtype
	//
	CatalogType *string `field:"optional" json:"catalogType" yaml:"catalogType"`
	// Turns on or off data lake access for Apache Spark applications that access Amazon Redshift databases in the Data Catalog from any non-Redshift engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-datalakeaccess
	//
	DataLakeAccess interface{} `field:"optional" json:"dataLakeAccess" yaml:"dataLakeAccess"`
	// A role that will be assumed by Glue for transferring data into/out of the staging bucket during a query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-datatransferrole
	//
	DataTransferRole *string `field:"optional" json:"dataTransferRole" yaml:"dataTransferRole"`
	// An encryption key that will be used for the staging bucket that will be created along with the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The name of the managed workgroup associated with the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-managedworkgroupname
	//
	ManagedWorkgroupName *string `field:"optional" json:"managedWorkgroupName" yaml:"managedWorkgroupName"`
	// The status of the managed workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-managedworkgroupstatus
	//
	ManagedWorkgroupStatus *string `field:"optional" json:"managedWorkgroupStatus" yaml:"managedWorkgroupStatus"`
	// The name of the Redshift database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-datalakeaccessproperties.html#cfn-glue-catalog-datalakeaccessproperties-redshiftdatabasename
	//
	RedshiftDatabaseName *string `field:"optional" json:"redshiftDatabaseName" yaml:"redshiftDatabaseName"`
}

