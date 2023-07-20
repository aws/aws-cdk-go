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
	// A friendly description for this DB cluster parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB cluster parameter group family name.
	//
	// A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
	//
	// > The DB cluster parameter group family can't be changed when updating a DB cluster parameter group.
	//
	// To list all of the available parameter group families, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"`
	//
	// The output contains duplicates.
	//
	// For more information, see `[CreateDBClusterParameterGroup](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_CreateDBClusterParameterGroup.html)` .
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
	// If you don't specify a value for `DBClusterParameterGroupName` property, a name is automatically created for the DB cluster parameter group.
	//
	// > This value is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-dbclusterparametergroupname
	//
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// An optional array of key-value pairs to apply to this DB cluster parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

