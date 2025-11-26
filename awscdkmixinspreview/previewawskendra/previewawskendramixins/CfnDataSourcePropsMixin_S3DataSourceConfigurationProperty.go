package previewawskendramixins


// Provides the configuration information to connect to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3DataSourceConfigurationProperty := &S3DataSourceConfigurationProperty{
//   	AccessControlListConfiguration: &AccessControlListConfigurationProperty{
//   		KeyPath: jsii.String("keyPath"),
//   	},
//   	BucketName: jsii.String("bucketName"),
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
type CfnDataSourcePropsMixin_S3DataSourceConfigurationProperty struct {
	// Provides the path to the S3 bucket that contains the user context filtering files for the data source.
	//
	// For the format of the file, see [Access control for S3 data sources](https://docs.aws.amazon.com/kendra/latest/dg/s3-acl.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-accesscontrollistconfiguration
	//
	AccessControlListConfiguration interface{} `field:"optional" json:"accessControlListConfiguration" yaml:"accessControlListConfiguration"`
	// The name of the bucket that contains the documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
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
	// - * /myapp/config/** —All files inside config directory.
	// - *** /*.png* —All .png files in all directories.
	// - *** /*.{png, ico, md}* —All .png, .ico or .md files in all directories.
	// - * /myapp/src/** /*.ts* —All .ts files inside src directory (and all its subdirectories).
	// - *** /!(*.module).ts* —All .ts files but not .module.ts
	// - **.png , *.jpg* —All PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** —All files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** —All internal-related files in a directory and its subdirectories.
	//
	// For more examples, see [Use of Exclude and Include Filters](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-exclusionpatterns
	//
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of glob patterns (patterns that can expand a wildcard pattern into a list of path names that match the given pattern) for certain file names and file types to include in your index.
	//
	// If a document matches both an inclusion and exclusion prefix or pattern, the exclusion prefix takes precendence and the document is not indexed. Examples of glob patterns include:
	//
	// - * /myapp/config/** —All files inside config directory.
	// - *** /*.png* —All .png files in all directories.
	// - *** /*.{png, ico, md}* —All .png, .ico or .md files in all directories.
	// - * /myapp/src/** /*.ts* —All .ts files inside src directory (and all its subdirectories).
	// - *** /!(*.module).ts* —All .ts files but not .module.ts
	// - **.png , *.jpg* —All PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** —All files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** —All internal-related files in a directory and its subdirectories.
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

