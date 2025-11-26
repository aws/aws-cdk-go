package previewawskendramixins


// Maps attributes or field names of Confluence blog to Amazon Kendra index field names.
//
// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   confluenceBlogToIndexFieldMappingProperty := &ConfluenceBlogToIndexFieldMappingProperty{
//   	DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	DateFieldFormat: jsii.String("dateFieldFormat"),
//   	IndexFieldName: jsii.String("indexFieldName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html
//
type CfnDataSourcePropsMixin_ConfluenceBlogToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-datasourcefieldname
	//
	DataSourceFieldName *string `field:"optional" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-datefieldformat
	//
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-indexfieldname
	//
	IndexFieldName *string `field:"optional" json:"indexFieldName" yaml:"indexFieldName"`
}

