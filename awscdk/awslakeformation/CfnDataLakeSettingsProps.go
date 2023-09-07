package awslakeformation


// Properties for defining a `CfnDataLakeSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnDataLakeSettingsProps := &CfnDataLakeSettingsProps{
//   	Admins: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	AllowExternalDataFiltering: jsii.Boolean(false),
//   	AllowFullTableExternalDataAccess: jsii.Boolean(false),
//   	AuthorizedSessionTagValueList: []*string{
//   		jsii.String("authorizedSessionTagValueList"),
//   	},
//   	CreateDatabaseDefaultPermissions: []interface{}{
//   		&PrincipalPermissionsProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	CreateTableDefaultPermissions: []interface{}{
//   		&PrincipalPermissionsProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	ExternalDataFilteringAllowList: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	MutationType: jsii.String("mutationType"),
//   	Parameters: parameters,
//   	TrustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html
//
type CfnDataLakeSettingsProps struct {
	// A list of AWS Lake Formation principals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-admins
	//
	Admins interface{} `field:"optional" json:"admins" yaml:"admins"`
	// Whether to allow Amazon EMR clusters or other third-party query engines to access data managed by Lake Formation .
	//
	// If set to true, you allow Amazon EMR clusters or other third-party engines to access data in Amazon S3 locations that are registered with Lake Formation .
	//
	// If false or null, no third-party query engines will be able to access data in Amazon S3 locations that are registered with Lake Formation.
	//
	// For more information, see [External data filtering setting](https://docs.aws.amazon.com/lake-formation/latest/dg/initial-LF-setup.html#external-data-filter) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-allowexternaldatafiltering
	//
	AllowExternalDataFiltering interface{} `field:"optional" json:"allowExternalDataFiltering" yaml:"allowExternalDataFiltering"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-allowfulltableexternaldataaccess
	//
	AllowFullTableExternalDataAccess interface{} `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	//
	// Lake Formation will publish the acceptable key-value pair, for example key = "LakeFormationTrustedCaller" and value = "TRUE" and the third party integrator must properly tag the temporary security credentials that will be used to call Lake Formation 's administrative API operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-authorizedsessiontagvaluelist
	//
	AuthorizedSessionTagValueList *[]*string `field:"optional" json:"authorizedSessionTagValueList" yaml:"authorizedSessionTagValueList"`
	// Specifies whether access control on a newly created database is managed by Lake Formation permissions or exclusively by IAM permissions.
	//
	// A null value indicates that the access is controlled by Lake Formation permissions. `ALL` permissions assigned to `IAM_ALLOWED_PRINCIPALS` group indicates that the user's IAM permissions determine the access to the database. This is referred to as the setting "Use only IAM access control," and is to support backward compatibility with the AWS Glue permission model implemented by IAM permissions.
	//
	// The only permitted values are an empty array or an array that contains a single JSON object that grants `ALL` to `IAM_ALLOWED_PRINCIPALS` .
	//
	// For more information, see [Changing the default security settings for your data lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-createdatabasedefaultpermissions
	//
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// Specifies whether access control on a newly created table is managed by Lake Formation permissions or exclusively by IAM permissions.
	//
	// A null value indicates that the access is controlled by Lake Formation permissions. `ALL` permissions assigned to `IAM_ALLOWED_PRINCIPALS` group indicate that the user's IAM permissions determine the access to the table. This is referred to as the setting "Use only IAM access control," and is to support the backward compatibility with the AWS Glue permission model implemented by IAM permissions.
	//
	// The only permitted values are an empty array or an array that contains a single JSON object that grants `ALL` permissions to `IAM_ALLOWED_PRINCIPALS` .
	//
	// For more information, see [Changing the default security settings for your data lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-createtabledefaultpermissions
	//
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A list of the account IDs of AWS accounts with Amazon EMR clusters or third-party engines that are allwed to perform data filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-externaldatafilteringallowlist
	//
	ExternalDataFilteringAllowList interface{} `field:"optional" json:"externalDataFilteringAllowList" yaml:"externalDataFilteringAllowList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-mutationtype
	//
	MutationType *string `field:"optional" json:"mutationType" yaml:"mutationType"`
	// A key-value map that provides an additional configuration on your data lake.
	//
	// `CrossAccountVersion` is the key you can configure in the `Parameters` field. Accepted values for the `CrossAccountVersion` key are 1, 2, and 3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of UTF-8 strings.
	//
	// A list of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs). The user ARNs can be logged in the resource owner's CloudTrail log. You may want to specify this property when you are in a high-trust boundary, such as the same team or company.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-trustedresourceowners
	//
	TrustedResourceOwners *[]*string `field:"optional" json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

