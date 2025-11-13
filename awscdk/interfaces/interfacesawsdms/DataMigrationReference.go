package interfacesawsdms


// A reference to a DataMigration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataMigrationReference := &DataMigrationReference{
//   	DataMigrationArn: jsii.String("dataMigrationArn"),
//   }
//
type DataMigrationReference struct {
	// The DataMigrationArn of the DataMigration resource.
	DataMigrationArn *string `field:"required" json:"dataMigrationArn" yaml:"dataMigrationArn"`
}

