package awsrds


// Properties for MySQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.mysql}.
//
// Example:
//   var vpc vpc
//
//   role := iam.NewRole(this, jsii.String("RDSDirectoryServicesRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
//   	ManagedPolicies: []iManagedPolicy{
//   		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
//   	},
//   })
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_19(),
//   	}),
//   	Vpc: Vpc,
//   	Domain: jsii.String("d-????????"),
//   	 // The ID of the domain for the instance to join.
//   	DomainRole: role,
//   })
//
// Experimental.
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version MysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

