package awskendra


// Maps a column or attribute in the data source to an index field.
//
// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceToIndexFieldMappingProperty := &DataSourceToIndexFieldMappingProperty{
//   	DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	IndexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	DateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourcetoindexfieldmapping.html
//
type CfnDataSource_DataSourceToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	//
	// You must first create the index field using the `UpdateIndex` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourcetoindexfieldmapping.html#cfn-kendra-datasource-datasourcetoindexfieldmapping-datasourcefieldname
	//
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the data source field.
	//
	// The index field type must match the data source field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourcetoindexfieldmapping.html#cfn-kendra-datasource-datasourcetoindexfieldmapping-indexfieldname
	//
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field, you must specify the date format. If the field is not a date field, an exception is thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourcetoindexfieldmapping.html#cfn-kendra-datasource-datasourcetoindexfieldmapping-datefieldformat
	//
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

