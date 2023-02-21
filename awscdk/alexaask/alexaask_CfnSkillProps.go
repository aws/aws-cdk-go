package alexaask


// Properties for defining a `CfnSkill`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   cfnSkillProps := &cfnSkillProps{
//   	authenticationConfiguration: &authenticationConfigurationProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	skillPackage: &skillPackageProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//
//   		// the properties below are optional
//   		overrides: &overridesProperty{
//   			manifest: manifest,
//   		},
//   		s3BucketRole: jsii.String("s3BucketRole"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	vendorId: jsii.String("vendorId"),
//   }
//
type CfnSkillProps struct {
	// Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
	//
	// Only Login with Amazon clients created through the  are supported. The client ID, client secret, and refresh token are required.
	AuthenticationConfiguration interface{} `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Configuration for the skill package that contains the components of the Alexa skill.
	//
	// Skill packages are retrieved from an Amazon S3 bucket and key and used to create and update the skill. For more information about the skill package format, see the  .
	SkillPackage interface{} `field:"required" json:"skillPackage" yaml:"skillPackage"`
	// The vendor ID associated with the Amazon developer account that will host the skill.
	//
	// Details for retrieving the vendor ID are in  . The provided LWA credentials must be linked to the developer account associated with this vendor ID.
	VendorId *string `field:"required" json:"vendorId" yaml:"vendorId"`
}

