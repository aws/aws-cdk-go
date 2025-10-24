package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnDBParameterGroupProps := &CfnDBParameterGroupProps{
//   	Description: jsii.String("description"),
//   	Family: jsii.String("family"),
//
//   	// the properties below are optional
//   	DbParameterGroupName: jsii.String("dbParameterGroupName"),
//   	Parameters: parameters,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html
//
type CfnDBParameterGroupProps struct {
	// Provides the customer-specified description for this DB parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB parameter group family name.
	//
	// A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family.
	//
	// To list all of the available parameter group families for a DB engine, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`
	//
	// For example, to list all of the available parameter group families for the MySQL DB engine, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine mysql`
	//
	// > The output contains duplicates.
	//
	// The following are the valid DB engine values:
	//
	// - `aurora-mysql`
	// - `aurora-postgresql`
	// - `db2-ae`
	// - `db2-se`
	// - `mysql`
	// - `oracle-ee`
	// - `oracle-ee-cdb`
	// - `oracle-se2`
	// - `oracle-se2-cdb`
	// - `postgres`
	// - `sqlserver-ee`
	// - `sqlserver-se`
	// - `sqlserver-ex`
	// - `sqlserver-web`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-family
	//
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the DB parameter group.
	//
	// Constraints:
	//
	// - Must be 1 to 255 letters, numbers, or hyphens.
	// - First character must be a letter
	// - Can't end with a hyphen or contain two consecutive hyphens
	//
	// If you don't specify a value for `DBParameterGroupName` property, a name is automatically created for the DB parameter group.
	//
	// > This value is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-dbparametergroupname
	//
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// A mapping of parameter names and values for the parameter update.
	//
	// You must specify at least one parameter name and value.
	//
	// For more information about parameter groups, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide* , or [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide* .
	//
	// > AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Tags to assign to the DB parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

