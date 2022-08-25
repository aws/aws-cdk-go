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
//   cfnDBParameterGroupProps := &cfnDBParameterGroupProps{
//   	description: jsii.String("description"),
//   	family: jsii.String("family"),
//
//   	// the properties below are optional
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
type CfnDBParameterGroupProps struct {
	// Provides the customer-specified description for this DB parameter group.
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
	Family *string `field:"required" json:"family" yaml:"family"`
	// An array of parameter names and values for the parameter update.
	//
	// At least one parameter name and value must be supplied. Subsequent arguments are optional.
	//
	// For more information about DB parameters and DB parameter groups for Amazon RDS DB engines, see [Working with DB Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide* .
	//
	// For more information about DB cluster and DB instance parameters and parameter groups for Amazon Aurora DB engines, see [Working with DB Parameter Groups and DB Cluster Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide* .
	//
	// > AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Tags to assign to the DB parameter group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

