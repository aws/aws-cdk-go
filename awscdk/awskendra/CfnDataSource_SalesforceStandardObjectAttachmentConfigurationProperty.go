package awskendra


// Provides the configuration information for processing attachments to Salesforce standard objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceStandardObjectAttachmentConfigurationProperty := &SalesforceStandardObjectAttachmentConfigurationProperty{
//   	DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	FieldMappings: []interface{}{
//   		&DataSourceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectattachmentconfiguration.html
//
type CfnDataSource_SalesforceStandardObjectAttachmentConfigurationProperty struct {
	// The name of the field used for the document title.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectattachmentconfiguration.html#cfn-kendra-datasource-salesforcestandardobjectattachmentconfiguration-documenttitlefieldname
	//
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// One or more objects that map fields in attachments to Amazon Kendra index fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectattachmentconfiguration.html#cfn-kendra-datasource-salesforcestandardobjectattachmentconfiguration-fieldmappings
	//
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

