package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataCatalog`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataCatalogProps := &CfnDataCatalogProps{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html
//
type CfnDataCatalogProps struct {
	// The name of the data catalog.
	//
	// The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of data catalog: `LAMBDA` for a federated catalog, `GLUE` for AWS Glue Catalog, or `HIVE` for an external hive metastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the data catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-description
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The tags (key-value pairs) to associate with this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

