package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCatalogPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCatalogMixinProps := &CfnCatalogMixinProps{
//   	AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   	CatalogProperties: &CatalogPropertiesProperty{
//   		CustomProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   		DataLakeAccessProperties: &DataLakeAccessPropertiesProperty{
//   			AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   			CatalogType: jsii.String("catalogType"),
//   			DataLakeAccess: jsii.Boolean(false),
//   			DataTransferRole: jsii.String("dataTransferRole"),
//   			KmsKey: jsii.String("kmsKey"),
//   			ManagedWorkgroupName: jsii.String("managedWorkgroupName"),
//   			ManagedWorkgroupStatus: jsii.String("managedWorkgroupStatus"),
//   			RedshiftDatabaseName: jsii.String("redshiftDatabaseName"),
//   		},
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
//   	Description: jsii.String("description"),
//   	FederatedCatalog: &FederatedCatalogProperty{
//   		ConnectionName: jsii.String("connectionName"),
//   		Identifier: jsii.String("identifier"),
//   	},
//   	Name: jsii.String("name"),
//   	OverwriteChildResourcePermissionsWithDefault: jsii.String("overwriteChildResourcePermissionsWithDefault"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetRedshiftCatalog: &TargetRedshiftCatalogProperty{
//   		CatalogArn: jsii.String("catalogArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html
//
type CfnCatalogMixinProps struct {
	// Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-allowfulltableexternaldataaccess
	//
	AllowFullTableExternalDataAccess *string `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// A structure that specifies data lake access properties and other custom properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-catalogproperties
	//
	CatalogProperties interface{} `field:"optional" json:"catalogProperties" yaml:"catalogProperties"`
	// An array of PrincipalPermissions objects for default database permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-createdatabasedefaultpermissions
	//
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// An array of PrincipalPermissions objects for default table permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-createtabledefaultpermissions
	//
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A FederatedCatalog structure that references an entity outside the Glue Data Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-federatedcatalog
	//
	FederatedCatalog interface{} `field:"optional" json:"federatedCatalog" yaml:"federatedCatalog"`
	// The name of the catalog to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether to overwrite child resource permissions with the default permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-overwritechildresourcepermissionswithdefault
	//
	OverwriteChildResourcePermissionsWithDefault *string `field:"optional" json:"overwriteChildResourcePermissionsWithDefault" yaml:"overwriteChildResourcePermissionsWithDefault"`
	// A map of key-value pairs that define parameters and properties of the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A structure that describes a target catalog for resource linking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html#cfn-glue-catalog-targetredshiftcatalog
	//
	TargetRedshiftCatalog interface{} `field:"optional" json:"targetRedshiftCatalog" yaml:"targetRedshiftCatalog"`
}

