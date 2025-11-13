package interfacesawsrds


// A reference to a DBShardGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBShardGroupReference := &DBShardGroupReference{
//   	DbShardGroupIdentifier: jsii.String("dbShardGroupIdentifier"),
//   }
//
type DBShardGroupReference struct {
	// The DBShardGroupIdentifier of the DBShardGroup resource.
	DbShardGroupIdentifier *string `field:"required" json:"dbShardGroupIdentifier" yaml:"dbShardGroupIdentifier"`
}

