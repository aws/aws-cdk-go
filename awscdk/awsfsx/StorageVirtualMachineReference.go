package awsfsx


// A reference to a StorageVirtualMachine resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageVirtualMachineReference := &StorageVirtualMachineReference{
//   	StorageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//   }
//
type StorageVirtualMachineReference struct {
	// The StorageVirtualMachineId of the StorageVirtualMachine resource.
	StorageVirtualMachineId *string `field:"required" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
}

