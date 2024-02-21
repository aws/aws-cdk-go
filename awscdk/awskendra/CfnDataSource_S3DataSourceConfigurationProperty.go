package awskendra


// Provides the configuration information to connect to an Amazon S3 bucket.
//
// > Amazon Kendra now supports an upgraded Amazon S3 connector.
// >
// > You must now use the [TemplateConfiguration](https://docs.aws.amazon.com/kendra/latest/APIReference/API_TemplateConfiguration.html) object instead of the `S3DataSourceConfiguration` object to configure your connector.
// >
// > Connectors configured using the older console and API architecture will continue to function as configured. However, you won't be able to edit or update them. If you want to edit or update your connector configuration, you must create a new connector.
// >
// > We recommended migrating your connector workflow to the upgraded version. Support for connectors configured using the older architecture is scheduled to end by June 2024.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataSourceConfigurationProperty := &S3DataSourceConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	AccessControlListConfiguration: &AccessControlListConfigurationProperty{
//   		KeyPath: jsii.String("keyPath"),
//   	},
//   	DocumentsMetadataConfiguration: &DocumentsMetadataConfigurationProperty{
//   		S3Prefix: jsii.String("s3Prefix"),
//   	},
//   	ExclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	InclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	InclusionPrefixes: []*string{
//   		jsii.String("inclusionPrefixes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html
//
type CfnDataSource_S3DataSourceConfigurationProperty struct {
	// The name of the bucket that contains the documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Provides the path to the S3 bucket that contains the user context filtering files for the data source.
	//
	// For the format of the file, see [Access control for S3 data sources](https://docs.aws.amazon.com/kendra/latest/dg/s3-acl.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-accesscontrollistconfiguration
	//
	AccessControlListConfiguration interface{} `field:"optional" json:"accessControlListConfiguration" yaml:"accessControlListConfiguration"`
	// Specifies document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes.
	//
	// Each metadata file contains metadata about a single document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-documentsmetadataconfiguration
	//
	DocumentsMetadataConfiguration interface{} `field:"optional" json:"documentsMetadataConfiguration" yaml:"documentsMetadataConfiguration"`
	// A list of glob patterns (patterns that can expand a wildcard pattern into a list of path names that match the given pattern) for certain file names and file types to exclude from your index.
	//
	// If a document matches both an inclusion and exclusion prefix or pattern, the exclusion prefix takes precendence and the document is not indexed. Examples of glob patterns include:
	//
	// - `/myapp/config/*` - All files inside config directory
	// - `/** /*.png` - All .png files in all directories
	// - `/** /*.{png,ico,md}` - All .png, .ico or .md files in all directories
	// - `/myapp/src/** /*.ts` - All .ts files inside src directory (and all its subdirectories)
	// - `** /!(*.module).ts` - All .ts files but not .module.ts
	// - **.png , *.jpg* excludes all PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** excludes all files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** excludes all internal-related files in a directory and its subdirectories.
	//
	// For more examples, see [Use of Exclude and Include Filters](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-exclusionpatterns
	//
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of glob patterns (patterns that can expand a wildcard pattern into a list of path names that match the given pattern) for certain file names and file types to include in your index.
	//
	// If a document matches both an inclusion and exclusion prefix or pattern, the exclusion prefix takes precendence and the document is not indexed. Examples of glob patterns include:
	//
	// - `/myapp/config/*` - All files inside config directory
	// - `/** /*.png` - All .png files in all directories
	// - `/** /*.{png,ico,md}` - All .png, .ico or .md files in all directories
	// - `/myapp/src/** /*.ts` - All .ts files inside src directory (and all its subdirectories)
	// - `** /!(*.module).ts` - All .ts files but not .module.ts
	// - **.png , *.jpg* includes all PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** includes all files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** includes all internal-related files in a directory and its subdirectories.
	//
	// For more examples, see [Use of Exclude and Include Filters](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-inclusionpatterns
	//
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// A list of S3 prefixes for the documents that should be included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-inclusionprefixes
	//
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

