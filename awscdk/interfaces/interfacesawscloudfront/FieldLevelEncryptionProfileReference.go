package interfacesawscloudfront


// A reference to a FieldLevelEncryptionProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldLevelEncryptionProfileReference := &FieldLevelEncryptionProfileReference{
//   	FieldLevelEncryptionProfileArn: jsii.String("fieldLevelEncryptionProfileArn"),
//   }
//
type FieldLevelEncryptionProfileReference struct {
	// The Arn of the FieldLevelEncryptionProfile resource.
	FieldLevelEncryptionProfileArn *string `field:"required" json:"fieldLevelEncryptionProfileArn" yaml:"fieldLevelEncryptionProfileArn"`
}

