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
//   	Tags: []cfnTag{
//   		&cfnTag{
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
	// A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a DB engine and engine version compatible with that DB parameter group family.
	//
	// > The DB parameter group family can't be changed when updating a DB parameter group.
	//
	// To list all of the available parameter group families, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"`
	//
	// The output contains duplicates.
	//
	// For more information, see `[CreateDBParameterGroup](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_CreateDBParameterGroup.html)` .
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
	// An array of parameter names and values for the parameter update.
	//
	// At least one parameter name and value must be supplied. Subsequent arguments are optional.
	//
	// RDS for Db2 requires you to bring your own Db2 license. You must enter your IBM customer ID ( `rds.ibm_customer_id` ) and site number ( `rds.ibm_site_id` ) before starting a Db2 instance.
	//
	// For more information about DB parameters and DB parameter groups for Amazon RDS DB engines, see [Working with DB Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide* .
	//
	// For more information about DB cluster and DB instance parameters and parameter groups for Amazon Aurora DB engines, see [Working with DB Parameter Groups and DB Cluster Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide* .
	//
	// > AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An optional array of key-value pairs to apply to this DB parameter group.
	//
	// > Currently, this is the only property that supports drift detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbparametergroup.html#cfn-rds-dbparametergroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

