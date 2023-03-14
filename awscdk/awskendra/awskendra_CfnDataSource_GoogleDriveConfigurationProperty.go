package awskendra


// Provides the configuration information to connect to Google Drive as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   googleDriveConfigurationProperty := &googleDriveConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	excludeMimeTypes: []*string{
//   		jsii.String("excludeMimeTypes"),
//   	},
//   	excludeSharedDrives: []*string{
//   		jsii.String("excludeSharedDrives"),
//   	},
//   	excludeUserAccounts: []*string{
//   		jsii.String("excludeUserAccounts"),
//   	},
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   }
//
type CfnDataSource_GoogleDriveConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of a AWS Secrets Manager secret that contains the credentials required to connect to Google Drive.
	//
	// For more information, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// A list of MIME types to exclude from the index. All documents matching the specified MIME type are excluded.
	//
	// For a list of MIME types, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	ExcludeMimeTypes *[]*string `field:"optional" json:"excludeMimeTypes" yaml:"excludeMimeTypes"`
	// A list of identifiers or shared drives to exclude from the index.
	//
	// All files and folders stored on the shared drive are excluded.
	ExcludeSharedDrives *[]*string `field:"optional" json:"excludeSharedDrives" yaml:"excludeSharedDrives"`
	// A list of email addresses of the users.
	//
	// Documents owned by these users are excluded from the index. Documents shared with excluded users are indexed unless they are excluded in another way.
	ExcludeUserAccounts *[]*string `field:"optional" json:"excludeUserAccounts" yaml:"excludeUserAccounts"`
	// A list of regular expression patterns to exclude certain items in your Google Drive, including shared drives and users' My Drives.
	//
	// Items that match the patterns are excluded from the index. Items that don't match the patterns are included in the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Maps Google Drive data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Google Drive fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Google Drive data source field names must exist in your Google Drive custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain items in your Google Drive, including shared drives and users' My Drives.
	//
	// Items that match the patterns are included in the index. Items that don't match the patterns are excluded from the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
}

