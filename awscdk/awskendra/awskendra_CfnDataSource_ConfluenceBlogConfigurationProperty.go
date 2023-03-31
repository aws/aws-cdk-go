package awskendra


// Configuration of blog settings for the Confluence data source.
//
// Blogs are always indexed unless filtered from the index by the `ExclusionPatterns` or `InclusionPatterns` fields in the `ConfluenceConfiguration` object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceBlogConfigurationProperty := &confluenceBlogConfigurationProperty{
//   	blogFieldMappings: []interface{}{
//   		&confluenceBlogToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceBlogConfigurationProperty struct {
	// Maps attributes or field names of Confluence blogs to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `BlogFieldMappings` parameter, you must specify at least one field mapping.
	BlogFieldMappings interface{} `field:"optional" json:"blogFieldMappings" yaml:"blogFieldMappings"`
}

