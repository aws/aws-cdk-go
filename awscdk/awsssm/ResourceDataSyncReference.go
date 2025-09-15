package awsssm


// A reference to a ResourceDataSync resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDataSyncReference := &ResourceDataSyncReference{
//   	SyncName: jsii.String("syncName"),
//   }
//
type ResourceDataSyncReference struct {
	// The SyncName of the ResourceDataSync resource.
	SyncName *string `field:"required" json:"syncName" yaml:"syncName"`
}

