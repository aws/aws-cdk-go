package awsrds


// Properties for SQL Server Enterprise Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerEe}.
//
// Example:
//   var vpc vpc
//
//
//   parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &parameterGroupProps{
//   	engine: rds.databaseInstanceEngine.sqlServerEe(&sqlServerEeInstanceEngineProps{
//   		version: rds.sqlServerEngineVersion_VER_11(),
//   	}),
//   	parameters: map[string]*string{
//   		"locks": jsii.String("100"),
//   	},
//   })
//
//   rds.NewDatabaseInstance(this, jsii.String("Database"), &databaseInstanceProps{
//   	engine: rds.*databaseInstanceEngine_SQL_SERVER_EE(),
//   	vpc: vpc,
//   	parameterGroup: parameterGroup,
//   })
//
type SqlServerEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `field:"required" json:"version" yaml:"version"`
}

