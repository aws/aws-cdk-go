package awsglue


// A structure that specifies data lake access properties and other custom properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogPropertiesProperty := &CatalogPropertiesProperty{
//   	CustomProperties: map[string]*string{
//   		"customPropertiesKey": jsii.String("customProperties"),
//   	},
//   	DataLakeAccessProperties: &DataLakeAccessPropertiesProperty{
//   		AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   		CatalogType: jsii.String("catalogType"),
//   		DataLakeAccess: jsii.Boolean(false),
//   		DataTransferRole: jsii.String("dataTransferRole"),
//   		KmsKey: jsii.String("kmsKey"),
//   		ManagedWorkgroupName: jsii.String("managedWorkgroupName"),
//   		ManagedWorkgroupStatus: jsii.String("managedWorkgroupStatus"),
//   		RedshiftDatabaseName: jsii.String("redshiftDatabaseName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-catalogproperties.html
//
type CfnCatalog_CatalogPropertiesProperty struct {
	// Additional key-value properties for the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-catalogproperties.html#cfn-glue-catalog-catalogproperties-customproperties
	//
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// Data lake access properties for the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-catalogproperties.html#cfn-glue-catalog-catalogproperties-datalakeaccessproperties
	//
	DataLakeAccessProperties interface{} `field:"optional" json:"dataLakeAccessProperties" yaml:"dataLakeAccessProperties"`
}

