package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a parameter group.
//
// Example:
//   var vpc Vpc
//
//
//   parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &ParameterGroupProps{
//   	Engine: rds.DatabaseInstanceEngine_SqlServerEe(&SqlServerEeInstanceEngineProps{
//   		Version: rds.SqlServerEngineVersion_VER_11(),
//   	}),
//   	Name: jsii.String("my-parameter-group"),
//   	Parameters: map[string]*string{
//   		"locks": jsii.String("100"),
//   	},
//   })
//
//   rds.NewDatabaseInstance(this, jsii.String("Database"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_SQL_SERVER_EE(),
//   	Vpc: Vpc,
//   	ParameterGroup: ParameterGroup,
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

