package interfacesawslightsail


// A reference to a DatabaseSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseSnapshotReference := &DatabaseSnapshotReference{
//   	DatabaseSnapshotArn: jsii.String("databaseSnapshotArn"),
//   	RelationalDatabaseSnapshotName: jsii.String("relationalDatabaseSnapshotName"),
//   }
//
type DatabaseSnapshotReference struct {
	// The ARN of the DatabaseSnapshot resource.
	DatabaseSnapshotArn *string `field:"required" json:"databaseSnapshotArn" yaml:"databaseSnapshotArn"`
	// The RelationalDatabaseSnapshotName of the DatabaseSnapshot resource.
	RelationalDatabaseSnapshotName *string `field:"required" json:"relationalDatabaseSnapshotName" yaml:"relationalDatabaseSnapshotName"`
}

