package previewalexaaskmixins


// Properties for CfnSkillPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var manifest interface{}
//
//   cfnSkillMixinProps := &CfnSkillMixinProps{
//   	AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	SkillPackage: &SkillPackageProperty{
//   		Overrides: &OverridesProperty{
//   			Manifest: manifest,
//   		},
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3BucketRole: jsii.String("s3BucketRole"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	VendorId: jsii.String("vendorId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html
//
type CfnSkillMixinProps struct {
	// Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
	//
	// Only Login with Amazon clients created through the  are supported. The client ID, client secret, and refresh token are required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Configuration for the skill package that contains the components of the Alexa skill.
	//
	// Skill packages are retrieved from an Amazon S3 bucket and key and used to create and update the skill. For more information about the skill package format, see the  .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-skillpackage
	//
	SkillPackage interface{} `field:"optional" json:"skillPackage" yaml:"skillPackage"`
	// The vendor ID associated with the Amazon developer account that will host the skill.
	//
	// Details for retrieving the vendor ID are in  . The provided LWA credentials must be linked to the developer account associated with this vendor ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-vendorid
	//
	VendorId *string `field:"optional" json:"vendorId" yaml:"vendorId"`
}

