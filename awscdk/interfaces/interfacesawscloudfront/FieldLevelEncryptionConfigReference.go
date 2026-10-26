package interfacesawscloudfront


// A reference to a FieldLevelEncryptionConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldLevelEncryptionConfigReference := &FieldLevelEncryptionConfigReference{
//   	FieldLevelEncryptionConfigArn: jsii.String("fieldLevelEncryptionConfigArn"),
//   }
//
type FieldLevelEncryptionConfigReference struct {
	// The Arn of the FieldLevelEncryptionConfig resource.
	FieldLevelEncryptionConfigArn *string `field:"required" json:"fieldLevelEncryptionConfigArn" yaml:"fieldLevelEncryptionConfigArn"`
}

