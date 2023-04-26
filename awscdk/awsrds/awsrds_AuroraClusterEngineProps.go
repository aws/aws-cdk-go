package awsrds


// Creation properties of the plain Aurora database cluster engine.
//
// Used in {@link DatabaseClusterEngine.aurora}.
//
// Example:
//   var vpc vpc
//
//   rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &DatabaseClusterFromSnapshotProps{
//   	Engine: rds.DatabaseClusterEngine_Aurora(&AuroraClusterEngineProps{
//   		Version: rds.AuroraEngineVersion_VER_1_22_2(),
//   	}),
//   	InstanceProps: &InstanceProps{
//   		Vpc: *Vpc,
//   	},
//   	SnapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
// Experimental.
type AuroraClusterEngineProps struct {
	// The version of the Aurora cluster engine.
	// Experimental.
	Version AuroraEngineVersion `field:"required" json:"version" yaml:"version"`
}

