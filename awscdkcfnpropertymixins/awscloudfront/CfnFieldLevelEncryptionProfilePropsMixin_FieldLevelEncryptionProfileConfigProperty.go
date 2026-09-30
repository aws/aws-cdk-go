package awscloudfront


// The configuration of a field-level encryption profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fieldLevelEncryptionProfileConfigProperty := &FieldLevelEncryptionProfileConfigProperty{
//   	CallerReference: jsii.String("callerReference"),
//   	Comment: jsii.String("comment"),
//   	EncryptionEntities: []interface{}{
//   		&EncryptionEntityProperty{
//   			FieldPatterns: []*string{
//   				jsii.String("fieldPatterns"),
//   			},
//   			ProviderId: jsii.String("providerId"),
//   			PublicKeyId: jsii.String("publicKeyId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig.html
//
type CfnFieldLevelEncryptionProfilePropsMixin_FieldLevelEncryptionProfileConfigProperty struct {
	// A unique value that identifies the creation request.
	//
	// Caller references are unique within an AWS account and cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig.html#cfn-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig-callerreference
	//
	CallerReference *string `field:"optional" json:"callerReference" yaml:"callerReference"`
	// An optional comment describing the field-level encryption profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig.html#cfn-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The encryption entities of the field-level encryption profile.
	//
	// At least one entity is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig.html#cfn-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig-encryptionentities
	//
	EncryptionEntities interface{} `field:"optional" json:"encryptionEntities" yaml:"encryptionEntities"`
	// The name of the field-level encryption profile.
	//
	// Names are unique within an AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig.html#cfn-cloudfront-fieldlevelencryptionprofile-fieldlevelencryptionprofileconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

