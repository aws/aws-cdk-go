package previewawskendramixins


// Provides information about how Amazon Kendra should use the columns of a database in an index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnConfigurationProperty := &ColumnConfigurationProperty{
//   	ChangeDetectingColumns: []*string{
//   		jsii.String("changeDetectingColumns"),
//   	},
//   	DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   	DocumentIdColumnName: jsii.String("documentIdColumnName"),
//   	DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   	FieldMappings: []interface{}{
//   		&DataSourceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html
//
type CfnDataSourcePropsMixin_ColumnConfigurationProperty struct {
	// One to five columns that indicate when a document in the database has changed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-changedetectingcolumns
	//
	ChangeDetectingColumns *[]*string `field:"optional" json:"changeDetectingColumns" yaml:"changeDetectingColumns"`
	// The column that contains the contents of the document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documentdatacolumnname
	//
	DocumentDataColumnName *string `field:"optional" json:"documentDataColumnName" yaml:"documentDataColumnName"`
	// The column that provides the document's identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documentidcolumnname
	//
	DocumentIdColumnName *string `field:"optional" json:"documentIdColumnName" yaml:"documentIdColumnName"`
	// The column that contains the title of the document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documenttitlecolumnname
	//
	DocumentTitleColumnName *string `field:"optional" json:"documentTitleColumnName" yaml:"documentTitleColumnName"`
	// An array of objects that map database column names to the corresponding fields in an index.
	//
	// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-fieldmappings
	//
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

