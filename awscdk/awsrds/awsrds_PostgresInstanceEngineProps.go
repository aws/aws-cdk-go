package awsrds


// Properties for PostgreSQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.postgres}.
//
// Example:
//   var vpc vpc
//
//   engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
//   	Version: rds.PostgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &DatabaseInstanceProps{
//   	Engine: Engine,
//   	Vpc: Vpc,
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("postgres"), &CredentialsBaseOptions{
//   		SecretName: jsii.String("my-cool-name"),
//   		EncryptionKey: myKey,
//   		ExcludeCharacters: jsii.String("!&*^#@()"),
//   		ReplicaRegions: []replicaRegion{
//   			&replicaRegion{
//   				Region: jsii.String("eu-west-1"),
//   			},
//   			&replicaRegion{
//   				Region: jsii.String("eu-west-2"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type PostgresInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version PostgresEngineVersion `field:"required" json:"version" yaml:"version"`
}

