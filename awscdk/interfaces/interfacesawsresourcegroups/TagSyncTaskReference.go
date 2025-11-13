package interfacesawsresourcegroups


// A reference to a TagSyncTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagSyncTaskReference := &TagSyncTaskReference{
//   	TaskArn: jsii.String("taskArn"),
//   }
//
type TagSyncTaskReference struct {
	// The TaskArn of the TagSyncTask resource.
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
}

