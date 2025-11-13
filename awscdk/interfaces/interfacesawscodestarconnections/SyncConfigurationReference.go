package interfacesawscodestarconnections


// A reference to a SyncConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncConfigurationReference := &SyncConfigurationReference{
//   	ResourceName: jsii.String("resourceName"),
//   	SyncType: jsii.String("syncType"),
//   }
//
type SyncConfigurationReference struct {
	// The ResourceName of the SyncConfiguration resource.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The SyncType of the SyncConfiguration resource.
	SyncType *string `field:"required" json:"syncType" yaml:"syncType"`
}

