package awskendra


// Maps attributes or field names of Confluence blog to Amazon Kendra index field names.
//
// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceBlogToIndexFieldMappingProperty := &ConfluenceBlogToIndexFieldMappingProperty{
//   	DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	IndexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	DateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html
//
type CfnDataSource_ConfluenceBlogToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-datasourcefieldname
	//
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-indexfieldname
	//
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-datefieldformat
	//
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

