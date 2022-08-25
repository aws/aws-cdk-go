package awsglue


// Specifies how job bookmark data should be encrypted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobBookmarksEncryptionProperty := &jobBookmarksEncryptionProperty{
//   	jobBookmarksEncryptionMode: jsii.String("jobBookmarksEncryptionMode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnSecurityConfiguration_JobBookmarksEncryptionProperty struct {
	// The encryption mode to use for job bookmarks data.
	JobBookmarksEncryptionMode *string `field:"optional" json:"jobBookmarksEncryptionMode" yaml:"jobBookmarksEncryptionMode"`
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

