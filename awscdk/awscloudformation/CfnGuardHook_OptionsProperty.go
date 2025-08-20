package awscloudformation


// Specifies the input parameters for a Guard Hook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &OptionsProperty{
//   	InputParams: &S3LocationProperty{
//   		Uri: jsii.String("uri"),
//
//   		// the properties below are optional
//   		VersionId: jsii.String("versionId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-options.html
//
type CfnGuardHook_OptionsProperty struct {
	// Specifies the S3 location where your input parameters are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-options.html#cfn-cloudformation-guardhook-options-inputparams
	//
	InputParams interface{} `field:"optional" json:"inputParams" yaml:"inputParams"`
}

