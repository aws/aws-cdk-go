package awsgamelift


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLocationProperty := &storageLocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnBuild_StorageLocationProperty struct {
	// `CfnBuild.StorageLocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnBuild.StorageLocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnBuild.StorageLocationProperty.RoleArn`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnBuild.StorageLocationProperty.ObjectVersion`.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

