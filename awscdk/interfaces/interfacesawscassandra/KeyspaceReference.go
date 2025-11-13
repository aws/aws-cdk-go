package interfacesawscassandra


// A reference to a Keyspace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyspaceReference := &KeyspaceReference{
//   	KeyspaceName: jsii.String("keyspaceName"),
//   }
//
type KeyspaceReference struct {
	// The KeyspaceName of the Keyspace resource.
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
}

