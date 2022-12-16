package awskendra


// Provides the configuration information to connect to OneDrive as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oneDriveConfigurationProperty := &oneDriveConfigurationProperty{
//   	oneDriveUsers: &oneDriveUsersProperty{
//   		oneDriveUserList: []*string{
//   			jsii.String("oneDriveUserList"),
//   		},
//   		oneDriveUserS3Path: &s3PathProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	secretArn: jsii.String("secretArn"),
//   	tenantDomain: jsii.String("tenantDomain"),
//
//   	// the properties below are optional
//   	disableLocalGroups: jsii.Boolean(false),
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
type CfnDataSource_OneDriveConfigurationProperty struct {
	// A list of user accounts whose documents should be indexed.
	OneDriveUsers interface{} `field:"required" json:"oneDriveUsers" yaml:"oneDriveUsers"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the user name and password to connect to OneDrive.
	//
	// The user named should be the application ID for the OneDrive application, and the password is the application key for the OneDrive application.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The Azure Active Directory domain of the organization.
	TenantDomain *string `field:"required" json:"tenantDomain" yaml:"tenantDomain"`
	// A Boolean value that specifies whether local groups are disabled ( `True` ) or enabled ( `False` ).
	DisableLocalGroups interface{} `field:"optional" json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// A list of regular expression patterns to exclude certain documents in your OneDrive.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the file name.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map OneDrive data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to OneDrive fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The OneDrive data source field names must exist in your OneDrive custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain documents in your OneDrive.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the file name.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
}

