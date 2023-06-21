package awsrds


// The storage type to be associated with the DB cluster.
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
type DBClusterStorageType string

const (
	// Storage type for Aurora DB standard clusters.
	DBClusterStorageType_AURORA DBClusterStorageType = "AURORA"
	// Storage type for Aurora DB I/O-Optimized clusters.
	DBClusterStorageType_AURORA_IOPT1 DBClusterStorageType = "AURORA_IOPT1"
)

