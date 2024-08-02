package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBClusterParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnDBClusterParameterGroupProps := &CfnDBClusterParameterGroupProps{
//   	Description: jsii.String("description"),
//   	Family: jsii.String("family"),
//   	Parameters: parameters,
//
//   	// the properties below are optional
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html
//
type CfnDBClusterParameterGroupProps struct {
	// The description for the DB cluster parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB cluster parameter group family name.
	//
	// A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.
	//
	// *Aurora MySQL*
	//
	// Example: `aurora-mysql5.7` , `aurora-mysql8.0`
	//
	// *Aurora PostgreSQL*
	//
	// Example: `aurora-postgresql14`
	//
	// *RDS for MySQL*
	//
	// Example: `mysql8.0`
	//
	// *RDS for PostgreSQL*
	//
	// Example: `postgres13`
	//
	// To list all of the available parameter group families for a DB engine, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`
	//
	// For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql`
	//
	// > The output contains duplicates.
	//
	// The following are the valid DB engine values:
	//
	// - `aurora-mysql`
	// - `aurora-postgresql`
	// - `mysql`
	// - `postgres`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-family
	//
	Family *string `field:"required" json:"family" yaml:"family"`
	// Provides a list of parameters for the DB cluster parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-parameters
	//
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	// - Must not match the name of an existing DB cluster parameter group.
	//
	// > This value is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-dbclusterparametergroupname
	//
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Tags to assign to the DB cluster parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

