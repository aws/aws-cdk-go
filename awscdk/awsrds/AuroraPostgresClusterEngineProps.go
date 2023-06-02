package awsrds


// Creation properties of the Aurora PostgreSQL database cluster engine.
//
// Used in `DatabaseClusterEngine.auroraPostgres`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
//   		Version: rds.AuroraPostgresEngineVersion_VER_15_2(),
//   	}),
//   	Credentials: rds.Credentials_FromUsername(jsii.String("adminuser"), &CredentialsFromUsernameOptions{
//   		Password: cdk.secretValue_UnsafePlainText(jsii.String("7959866cacc02c2d243ecfe177464fe6")),
//   	}),
//   	InstanceProps: &InstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_X2G, ec2.InstanceSize_XLARGE),
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PUBLIC,
//   		},
//   		Vpc: *Vpc,
//   	},
//   	StorageType: rds.DBClusterStorageType_AURORA_IOPT1,
//   })
//
type AuroraPostgresClusterEngineProps struct {
	// The version of the Aurora PostgreSQL cluster engine.
	Version AuroraPostgresEngineVersion `field:"required" json:"version" yaml:"version"`
}

