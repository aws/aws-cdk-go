package awssigner


// Properties for defining a `CfnSigningJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSigningJobProps := &CfnSigningJobProps{
//   	ProfileName: jsii.String("profileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingjob.html
//
type CfnSigningJobProps struct {
	// The name of the signing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingjob.html#cfn-signer-signingjob-profilename
	//
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
}

