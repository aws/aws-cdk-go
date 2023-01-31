package awslocation


// Properties for defining a `CfnGeofenceCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGeofenceCollectionProps := &cfnGeofenceCollectionProps{
//   	collectionName: jsii.String("collectionName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnGeofenceCollectionProps struct {
	// A custom name for the geofence collection.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique geofence collection name.
	// - No spaces allowed. For example, `ExampleGeofenceCollection` .
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// An optional description for the geofence collection.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key identifier for an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html) . Enter a key ID, key ARN, alias name, or alias ARN.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

