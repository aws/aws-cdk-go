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
//   	ConnectionType: jsii.String("connectionType"),
//   	Description: jsii.String("description"),
//   	Error: jsii.String("error"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
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
	// The type of connection for a `FEDERATED` data catalog (for example, `REDSHIFT` , `MYSQL` , or `SQLSERVER` ).
	//
	// For information about individual connectors, see [Available data source connectors](https://docs.aws.amazon.com/athena/latest/ug/connectors-available.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-connectiontype
	//
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// A description of the data catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Text of the error that occurred during data catalog creation or deletion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-error
	//
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Specifies the Lambda function or functions to use for creating the data catalog.
	//
	// This is a mapping whose values depend on the catalog type.
	//
	// - For the `HIVE` data catalog type, use the following syntax. The `metadata-function` parameter is required. `The sdk-version` parameter is optional and defaults to the currently supported version.
	//
	// `metadata-function= *lambda_arn* , sdk-version= *version_number*`
	// - For the `LAMBDA` data catalog type, use one of the following sets of required parameters, but not both.
	//
	// - If you have one Lambda function that processes metadata and another for reading the actual data, use the following syntax. Both parameters are required.
	//
	// `metadata-function= *lambda_arn* , record-function= *lambda_arn*`
	// - If you have a composite Lambda function that processes both metadata and data, use the following syntax to specify your Lambda function.
	//
	// `function= *lambda_arn*`
	// - The `GLUE` type takes a catalog ID parameter and is required. The `*catalog_id*` is the account ID of the AWS account to which the AWS Glue Data Catalog belongs.
	//
	// `catalog-id= *catalog_id*`
	//
	// - The `GLUE` data catalog type also applies to the default `AwsDataCatalog` that already exists in your account, of which you can have only one and cannot modify.
	// - The `FEDERATED` data catalog type uses one of the following parameters, but not both. Use `connection-arn` for an existing AWS Glue connection. Use `connection-type` and `connection-properties` to specify the configuration setting for a new connection.
	//
	// - `connection-arn: *<glue_connection_arn_to_reuse>*`
	// - `lambda-role-arn` (optional): The execution role to use for the Lambda function. If not provided, one is created.
	// - `connection-type:MYSQL|REDSHIFT|...., connection-properties:" *<json_string>* "`
	//
	// For *`<json_string>`* , use escaped JSON text, as in the following example.
	//
	// `"{\"spill_bucket\":\"my_spill\",\"spill_prefix\":\"athena-spill\",\"host\":\"abc12345.snowflakecomputing.com\",\"port\":\"1234\",\"warehouse\":\"DEV_WH\",\"database\":\"TEST\",\"schema\":\"PUBLIC\",\"SecretArn\":\"arn:aws:secretsmanager:ap-south-1:111122223333:secret:snowflake-XHb67j\"}"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The status of the creation or deletion of the data catalog.
	//
	// - The `LAMBDA` , `GLUE` , and `HIVE` data catalog types are created synchronously. Their status is either `CREATE_COMPLETE` or `CREATE_FAILED` .
	// - The `FEDERATED` data catalog type is created asynchronously.
	//
	// Data catalog creation status:
	//
	// - `CREATE_IN_PROGRESS` : Federated data catalog creation in progress.
	// - `CREATE_COMPLETE` : Data catalog creation complete.
	// - `CREATE_FAILED` : Data catalog could not be created.
	// - `CREATE_FAILED_CLEANUP_IN_PROGRESS` : Federated data catalog creation failed and is being removed.
	// - `CREATE_FAILED_CLEANUP_COMPLETE` : Federated data catalog creation failed and was removed.
	// - `CREATE_FAILED_CLEANUP_FAILED` : Federated data catalog creation failed but could not be removed.
	//
	// Data catalog deletion status:
	//
	// - `DELETE_IN_PROGRESS` : Federated data catalog deletion in progress.
	// - `DELETE_COMPLETE` : Federated data catalog deleted.
	// - `DELETE_FAILED` : Federated data catalog could not be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags (key-value pairs) to associate with this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

