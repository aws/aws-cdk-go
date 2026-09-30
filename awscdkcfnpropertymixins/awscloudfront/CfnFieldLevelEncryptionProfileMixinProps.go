package awscloudfront


// Properties for CfnFieldLevelEncryptionProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFieldLevelEncryptionProfileMixinProps := &CfnFieldLevelEncryptionProfileMixinProps{
//   	FieldLevelEncryptionProfileConfig: &FieldLevelEncryptionProfileConfigProperty{
//   		CallerReference: jsii.String("callerReference"),
//   		Comment: jsii.String("comment"),
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionprofile.html
//
type CfnFieldLevelEncryptionProfileMixinProps struct {
	// The configuration of a field-level encryption profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionprofile.html#cfn-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig
	//
	FieldLevelEncryptionProfileConfig interface{} `field:"optional" json:"fieldLevelEncryptionProfileConfig" yaml:"fieldLevelEncryptionProfileConfig"`
}

