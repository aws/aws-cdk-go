package awskendra


// Specifies configuration information for indexing a single standard object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceStandardObjectConfigurationProperty := &SalesforceStandardObjectConfigurationProperty{
//   	DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectconfiguration.html
//
type CfnDataSource_SalesforceStandardObjectConfigurationProperty struct {
	// The name of the field in the standard object table that contains the document contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectconfiguration.html#cfn-kendra-datasource-salesforcestandardobjectconfiguration-documentdatafieldname
	//
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the standard object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectconfiguration.html#cfn-kendra-datasource-salesforcestandardobjectconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the field in the standard object table that contains the document title.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectconfiguration.html#cfn-kendra-datasource-salesforcestandardobjectconfiguration-documenttitlefieldname
	//
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps attributes or field names of the standard object to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Salesforce fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Salesforce data source field names must exist in your Salesforce custom metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectconfiguration.html#cfn-kendra-datasource-salesforcestandardobjectconfiguration-fieldmappings
	//
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

