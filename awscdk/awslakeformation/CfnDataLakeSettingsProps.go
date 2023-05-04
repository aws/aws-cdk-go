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
//   	Parameters: parameters,
//   	TrustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   }
//
type CfnDataLakeSettingsProps struct {
	// A list of AWS Lake Formation principals.
	Admins interface{} `field:"optional" json:"admins" yaml:"admins"`
	// `AWS::LakeFormation::DataLakeSettings.AllowExternalDataFiltering`.
	AllowExternalDataFiltering interface{} `field:"optional" json:"allowExternalDataFiltering" yaml:"allowExternalDataFiltering"`
	// `AWS::LakeFormation::DataLakeSettings.AuthorizedSessionTagValueList`.
	AuthorizedSessionTagValueList *[]*string `field:"optional" json:"authorizedSessionTagValueList" yaml:"authorizedSessionTagValueList"`
	// `AWS::LakeFormation::DataLakeSettings.CreateDatabaseDefaultPermissions`.
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// `AWS::LakeFormation::DataLakeSettings.CreateTableDefaultPermissions`.
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// `AWS::LakeFormation::DataLakeSettings.ExternalDataFilteringAllowList`.
	ExternalDataFilteringAllowList interface{} `field:"optional" json:"externalDataFilteringAllowList" yaml:"externalDataFilteringAllowList"`
	// `AWS::LakeFormation::DataLakeSettings.Parameters`.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of UTF-8 strings.
	//
	// A list of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs). The user ARNs can be logged in the resource owner's CloudTrail log. You may want to specify this property when you are in a high-trust boundary, such as the same team or company.
	TrustedResourceOwners *[]*string `field:"optional" json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

