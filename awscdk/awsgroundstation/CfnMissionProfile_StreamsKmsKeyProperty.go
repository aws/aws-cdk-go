package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamsKmsKeyProperty := &StreamsKmsKeyProperty{
//   	KmsAliasArn: jsii.String("kmsAliasArn"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnMissionProfile_StreamsKmsKeyProperty struct {
	// `CfnMissionProfile.StreamsKmsKeyProperty.KmsAliasArn`.
	KmsAliasArn *string `field:"optional" json:"kmsAliasArn" yaml:"kmsAliasArn"`
	// `CfnMissionProfile.StreamsKmsKeyProperty.KmsKeyArn`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

