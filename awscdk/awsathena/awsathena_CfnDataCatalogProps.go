package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataCatalog`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataCatalogProps := &cfnDataCatalogProps{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataCatalogProps struct {
	// The name of the data catalog.
	//
	// The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of data catalog: `LAMBDA` for a federated catalog, `GLUE` for AWS Glue Catalog, or `HIVE` for an external hive metastore.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the data catalog.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the Lambda function or functions to use for the data catalog.
	//
	// The mapping used depends on the catalog type.
	//
	// - The `HIVE` data catalog type uses the following syntax. The `metadata-function` parameter is required. `The sdk-version` parameter is optional and defaults to the currently supported version.
	//
	// `metadata-function= *lambda_arn* , sdk-version= *version_number*`
	// - The `LAMBDA` data catalog type uses one of the following sets of required parameters, but not both.
	//
	// - When one Lambda function processes metadata and another Lambda function reads data, the following syntax is used. Both parameters are required.
	//
	// `metadata-function= *lambda_arn* , record-function= *lambda_arn*`
	// - A composite Lambda function that processes both metadata and data uses the following syntax.
	//
	// `function= *lambda_arn*`
	// - The `GLUE` type takes a catalog ID parameter and is required. The `*catalog_id*` is the account ID of the AWS account to which the Glue catalog belongs.
	//
	// `catalog-id= *catalog_id*`
	//
	// - The `GLUE` data catalog type also applies to the default `AwsDataCatalog` that already exists in your account, of which you can have only one and cannot modify.
	// - Queries that specify a GLUE data catalog other than the default `AwsDataCatalog` must be run on Athena engine version 2.
	// - In Regions where Athena engine version 2 is not available, creating new GLUE data catalogs results in an `INVALID_INPUT` error.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The tags (key-value pairs) to associate with this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

