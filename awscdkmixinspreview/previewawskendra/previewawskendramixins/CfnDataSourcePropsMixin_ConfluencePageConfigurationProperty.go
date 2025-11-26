package previewawskendramixins


// Configuration of the page settings for the Confluence data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   confluencePageConfigurationProperty := &ConfluencePageConfigurationProperty{
//   	PageFieldMappings: []interface{}{
//   		&ConfluencePageToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepageconfiguration.html
//
type CfnDataSourcePropsMixin_ConfluencePageConfigurationProperty struct {
	// Maps attributes or field names of Confluence pages to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `PageFieldMappings` parameter, you must specify at least one field mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepageconfiguration.html#cfn-kendra-datasource-confluencepageconfiguration-pagefieldmappings
	//
	PageFieldMappings interface{} `field:"optional" json:"pageFieldMappings" yaml:"pageFieldMappings"`
}

