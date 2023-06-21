package awsrds


// Properties for SQL Server Enterprise Edition instance engines.
//
// Used in `DatabaseInstanceEngine.sqlServerEe`.
//
// Example:
//   var vpc vpc
//
//
//   parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &ParameterGroupProps{
//   	Engine: rds.DatabaseInstanceEngine_SqlServerEe(&SqlServerEeInstanceEngineProps{
//   		Version: rds.SqlServerEngineVersion_VER_11(),
//   	}),
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
type SqlServerEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `field:"required" json:"version" yaml:"version"`
}

