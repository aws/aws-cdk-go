package alexaask


// The `SkillPackage` property type contains configuration details for the skill package that contains the components of the Alexa skill.
//
// Skill packages are retrieved from an Amazon S3 bucket and key and used to create and update the skill. More details about the skill package format are located in the  .
//
// `SkillPackage` is a property of the `Alexa::ASK::Skill` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   skillPackageProperty := &SkillPackageProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//
//   	// the properties below are optional
//   	Overrides: &OverridesProperty{
//   		Manifest: manifest,
//   	},
//   	S3BucketRole: jsii.String("s3BucketRole"),
//   	S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html
//
type CfnSkill_SkillPackageProperty struct {
	// The name of the Amazon S3 bucket where the .zip file that contains the skill package is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html#cfn-ask-skill-skillpackage-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The location and name of the skill package .zip file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html#cfn-ask-skill-skillpackage-s3key
	//
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
	// Overrides to the skill package to apply when creating or updating the skill.
	//
	// Values provided here do not modify the contents of the original skill package. Currently, only overriding values inside of the skill manifest component of the package is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html#cfn-ask-skill-skillpackage-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// ARN of the IAM role that grants the Alexa service ( `alexa-appkit.amazon.com` ) permission to access the bucket and retrieve the skill package. This property is optional. If you do not provide it, the bucket must be publicly accessible or configured with a policy that allows this access. Otherwise, AWS CloudFormation cannot create the skill.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html#cfn-ask-skill-skillpackage-s3bucketrole
	//
	S3BucketRole *string `field:"optional" json:"s3BucketRole" yaml:"s3BucketRole"`
	// If you have S3 versioning enabled, the version ID of the skill package.zip file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html#cfn-ask-skill-skillpackage-s3objectversion
	//
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

