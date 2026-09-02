package interfacesawselasticache


// A reference to a ServerlessCacheSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessCacheSnapshotReference := &ServerlessCacheSnapshotReference{
//   	ServerlessCacheSnapshotArn: jsii.String("serverlessCacheSnapshotArn"),
//   }
//
type ServerlessCacheSnapshotReference struct {
	// The ARN of the ServerlessCacheSnapshot resource.
	ServerlessCacheSnapshotArn *string `field:"required" json:"serverlessCacheSnapshotArn" yaml:"serverlessCacheSnapshotArn"`
}

