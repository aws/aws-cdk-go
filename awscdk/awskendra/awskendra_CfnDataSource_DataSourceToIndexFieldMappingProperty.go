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
//   dataSourceToIndexFieldMappingProperty := &dataSourceToIndexFieldMappingProperty{
//   	dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	indexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	dateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
type CfnDataSource_DataSourceToIndexFieldMappingProperty struct {
	// The name of the column or attribute in the data source.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the field in the index.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The type of data stored in the column or attribute.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

