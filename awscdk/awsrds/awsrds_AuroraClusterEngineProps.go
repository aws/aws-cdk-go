package awsrds


// Creation properties of the plain Aurora database cluster engine.
//
// Used in {@link DatabaseClusterEngine.aurora}.
//
// Example:
//   var vpc vpc
//
//   rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &databaseClusterFromSnapshotProps{
//   	engine: rds.databaseClusterEngine.aurora(&auroraClusterEngineProps{
//   		version: rds.auroraEngineVersion_VER_1_22_2(),
//   	}),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   	},
//   	snapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
type AuroraClusterEngineProps struct {
	// The version of the Aurora cluster engine.
	Version AuroraEngineVersion `field:"required" json:"version" yaml:"version"`
}

