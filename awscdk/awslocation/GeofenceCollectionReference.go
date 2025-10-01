package awslocation


// A reference to a GeofenceCollection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geofenceCollectionReference := &GeofenceCollectionReference{
//   	CollectionName: jsii.String("collectionName"),
//   	GeofenceCollectionArn: jsii.String("geofenceCollectionArn"),
//   }
//
type GeofenceCollectionReference struct {
	// The CollectionName of the GeofenceCollection resource.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// The ARN of the GeofenceCollection resource.
	GeofenceCollectionArn *string `field:"required" json:"geofenceCollectionArn" yaml:"geofenceCollectionArn"`
}

