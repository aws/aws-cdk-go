package awsrds


// Properties for MySQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.mysql}.
//
// Example:
//   var vpc vpc
//
//   role := iam.NewRole(this, jsii.String("RDSDirectoryServicesRole"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
//   	managedPolicies: []iManagedPolicy{
//   		iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
//   	},
//   })
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
//   		version: rds.mysqlEngineVersion_VER_8_0_19(),
//   	}),
//   	vpc: vpc,
//   	domain: jsii.String("d-????????"),
//   	 // The ID of the domain for the instance to join.
//   	domainRole: role,
//   })
//
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

