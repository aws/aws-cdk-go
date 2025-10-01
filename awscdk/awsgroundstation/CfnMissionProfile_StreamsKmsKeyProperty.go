package awsgroundstation


// KMS key info.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamsKmsKeyProperty := &StreamsKmsKeyProperty{
//   	KmsAliasArn: jsii.String("kmsAliasArn"),
//   	KmsAliasName: jsii.String("kmsAliasName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html
//
type CfnMissionProfile_StreamsKmsKeyProperty struct {
	// KMS Alias Arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html#cfn-groundstation-missionprofile-streamskmskey-kmsaliasarn
	//
	KmsAliasArn *string `field:"optional" json:"kmsAliasArn" yaml:"kmsAliasArn"`
	// KMS Alias Name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html#cfn-groundstation-missionprofile-streamskmskey-kmsaliasname
	//
	KmsAliasName *string `field:"optional" json:"kmsAliasName" yaml:"kmsAliasName"`
	// KMS Key Arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html#cfn-groundstation-missionprofile-streamskmskey-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

