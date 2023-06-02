package awsrds


// Creation properties of the plain Aurora database cluster engine.
//
// Used in `DatabaseClusterEngine.aurora`.
//
// Example:
//   var vpc vpc
//
//   rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &DatabaseClusterFromSnapshotProps{
//   	Engine: rds.DatabaseClusterEngine_Aurora(&AuroraClusterEngineProps{
//   		Version: rds.AuroraEngineVersion_VER_1_22_2(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
//   	Vpc: Vpc,
//   	SnapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
type AuroraClusterEngineProps struct {
	// The version of the Aurora cluster engine.
	Version AuroraEngineVersion `field:"required" json:"version" yaml:"version"`
}

