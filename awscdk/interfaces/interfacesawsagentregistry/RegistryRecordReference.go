package interfacesawsagentregistry


// A reference to a RegistryRecord resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryRecordReference := &RegistryRecordReference{
//   	RecordArn: jsii.String("recordArn"),
//   }
//
type RegistryRecordReference struct {
	// The RecordArn of the RegistryRecord resource.
	RecordArn *string `field:"required" json:"recordArn" yaml:"recordArn"`
}

