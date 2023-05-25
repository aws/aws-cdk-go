package awskendra


// Configuration of attachment settings for the Confluence data source.
//
// Attachment settings are optional, if you don't specify settings attachments, Amazon Kendra won't index them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceAttachmentConfigurationProperty := &ConfluenceAttachmentConfigurationProperty{
//   	AttachmentFieldMappings: []interface{}{
//   		&ConfluenceAttachmentToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	CrawlAttachments: jsii.Boolean(false),
//   }
//
type CfnDataSource_ConfluenceAttachmentConfigurationProperty struct {
	// Maps attributes or field names of Confluence attachments to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `AttachentFieldMappings` parameter, you must specify at least one field mapping.
	AttachmentFieldMappings interface{} `field:"optional" json:"attachmentFieldMappings" yaml:"attachmentFieldMappings"`
	// `TRUE` to index attachments of pages and blogs in Confluence.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
}

