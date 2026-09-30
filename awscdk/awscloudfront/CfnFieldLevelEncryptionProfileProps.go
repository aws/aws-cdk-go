package awscloudfront


// Properties for defining a `CfnFieldLevelEncryptionProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFieldLevelEncryptionProfileProps := &CfnFieldLevelEncryptionProfileProps{
//   	FieldLevelEncryptionProfileConfig: &FieldLevelEncryptionProfileConfigProperty{
//   		CallerReference: jsii.String("callerReference"),
//   		EncryptionEntities: []interface{}{
//   			&EncryptionEntityProperty{
//   				FieldPatterns: []*string{
//   					jsii.String("fieldPatterns"),
//   				},
//   				ProviderId: jsii.String("providerId"),
//   				PublicKeyId: jsii.String("publicKeyId"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Comment: jsii.String("comment"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionprofile.html
//
type CfnFieldLevelEncryptionProfileProps struct {
	// The configuration of a field-level encryption profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionprofile.html#cfn-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig
	//
	FieldLevelEncryptionProfileConfig interface{} `field:"required" json:"fieldLevelEncryptionProfileConfig" yaml:"fieldLevelEncryptionProfileConfig"`
}

