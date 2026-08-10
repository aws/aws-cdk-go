package awssigner


// Properties for CfnSigningJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSigningJobMixinProps := &CfnSigningJobMixinProps{
//   	ProfileName: jsii.String("profileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingjob.html
//
type CfnSigningJobMixinProps struct {
	// The name of the signing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingjob.html#cfn-signer-signingjob-profilename
	//
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
}

