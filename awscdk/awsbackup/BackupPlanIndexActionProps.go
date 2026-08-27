package awsbackup


// Properties for a BackupPlanIndexAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var indexActionResourceType IndexActionResourceType
//
//   backupPlanIndexActionProps := &BackupPlanIndexActionProps{
//   	ResourceTypes: []IndexActionResourceType{
//   		indexActionResourceType,
//   	},
//   }
//
type BackupPlanIndexActionProps struct {
	// Specifies the resource types to include in the index action.
	//
	// A backup index is only created when this is set, so at least one resource
	// type must be provided.
	ResourceTypes *[]IndexActionResourceType `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

