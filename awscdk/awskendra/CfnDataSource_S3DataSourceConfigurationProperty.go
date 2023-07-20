package awskendra


// Provides the configuration information to connect to an Amazon S3 bucket.
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
	// A list of glob patterns for documents that should not be indexed.
	//
	// If a document that matches an inclusion prefix or inclusion pattern also matches an exclusion pattern, the document is not indexed.
	//
	// Some [examples](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) are:
	//
	// - **.png , *.jpg* will exclude all PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** will exclude all files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** will exclude all internal-related files in a directory and its subdirectories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-exclusionpatterns
	//
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of glob patterns for documents that should be indexed.
	//
	// If a document that matches an inclusion pattern also matches an exclusion pattern, the document is not indexed.
	//
	// Some [examples](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) are:
	//
	// - **.txt* will include all text files in a directory (files with the extension .txt).
	// - *** /*.txt* will include all text files in a directory and its subdirectories.
	// - **tax** will include all files in a directory that contain 'tax' in the file name, such as 'tax', 'taxes', 'income_tax'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-inclusionpatterns
	//
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// A list of S3 prefixes for the documents that should be included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-inclusionprefixes
	//
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

