package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a parameter group.
//
// Example:
//   parameterGroup := rds.ParameterGroup_ForInstance(this, jsii.String("InstanceParameterGroup"), &ParameterGroupProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_35(),
//   	}),
//   	Description: jsii.String("Parameter group for MySQL"),
//   	Parameters: map[string]*string{
//   		"max_connections": jsii.String("150"),
//   		"slow_query_log": jsii.String("1"),
//   	},
//   })
//
type ParameterGroupProps struct {
	// The database engine for this parameter group.
	Engine IEngine `field:"required" json:"engine" yaml:"engine"`
	// Description for this parameter group.
	// Default: a CDK generated description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of this parameter group.
	// Default: - CloudFormation-generated name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters in this parameter group.
	// Default: - None.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Default: - RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

