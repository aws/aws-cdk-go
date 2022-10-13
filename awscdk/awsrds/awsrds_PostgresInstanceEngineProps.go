package awsrds


// Properties for PostgreSQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.postgres}.
//
// Example:
//   var vpc vpc
//
//   engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   	version: rds.postgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &databaseInstanceProps{
//   	engine: engine,
//   	vpc: vpc,
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("postgres"), &credentialsBaseOptions{
//   		secretName: jsii.String("my-cool-name"),
//   		encryptionKey: myKey,
//   		excludeCharacters: jsii.String("!&*^#@()"),
//   		replicaRegions: []replicaRegion{
//   			&replicaRegion{
//   				region: jsii.String("eu-west-1"),
//   			},
//   			&replicaRegion{
//   				region: jsii.String("eu-west-2"),
//   			},
//   		},
//   	}),
//   })
//
type PostgresInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version PostgresEngineVersion `field:"required" json:"version" yaml:"version"`
}

