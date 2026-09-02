package interfacesawsmgn


// A reference to a NetworkMigrationDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkMigrationDefinitionReference := &NetworkMigrationDefinitionReference{
//   	NetworkMigrationDefinitionArn: jsii.String("networkMigrationDefinitionArn"),
//   }
//
type NetworkMigrationDefinitionReference struct {
	// The Arn of the NetworkMigrationDefinition resource.
	NetworkMigrationDefinitionArn *string `field:"required" json:"networkMigrationDefinitionArn" yaml:"networkMigrationDefinitionArn"`
}

