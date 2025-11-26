package previewawskendramixins


// Provides the configuration information to connect to Google Drive as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   googleDriveConfigurationProperty := &GoogleDriveConfigurationProperty{
//   	ExcludeMimeTypes: []*string{
//   		jsii.String("excludeMimeTypes"),
//   	},
//   	ExcludeSharedDrives: []*string{
//   		jsii.String("excludeSharedDrives"),
//   	},
//   	ExcludeUserAccounts: []*string{
//   		jsii.String("excludeUserAccounts"),
//   	},
//   	ExclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	FieldMappings: []interface{}{
//   		&DataSourceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//   		},
//   	},
//   	InclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html
//
type CfnDataSourcePropsMixin_GoogleDriveConfigurationProperty struct {
	// A list of MIME types to exclude from the index. All documents matching the specified MIME type are excluded.
	//
	// For a list of MIME types, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludemimetypes
	//
	ExcludeMimeTypes *[]*string `field:"optional" json:"excludeMimeTypes" yaml:"excludeMimeTypes"`
	// A list of identifiers or shared drives to exclude from the index.
	//
	// All files and folders stored on the shared drive are excluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludeshareddrives
	//
	ExcludeSharedDrives *[]*string `field:"optional" json:"excludeSharedDrives" yaml:"excludeSharedDrives"`
	// A list of email addresses of the users.
	//
	// Documents owned by these users are excluded from the index. Documents shared with excluded users are indexed unless they are excluded in another way.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludeuseraccounts
	//
	ExcludeUserAccounts *[]*string `field:"optional" json:"excludeUserAccounts" yaml:"excludeUserAccounts"`
	// A list of regular expression patterns to exclude certain items in your Google Drive, including shared drives and users' My Drives.
	//
	// Items that match the patterns are excluded from the index. Items that don't match the patterns are included in the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-exclusionpatterns
	//
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Maps Google Drive data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Google Drive fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Google Drive data source field names must exist in your Google Drive custom metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-fieldmappings
	//
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain items in your Google Drive, including shared drives and users' My Drives.
	//
	// Items that match the patterns are included in the index. Items that don't match the patterns are excluded from the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-inclusionpatterns
	//
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// The Amazon Resource Name (ARN) of a AWS Secrets Manager secret that contains the credentials required to connect to Google Drive.
	//
	// For more information, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

