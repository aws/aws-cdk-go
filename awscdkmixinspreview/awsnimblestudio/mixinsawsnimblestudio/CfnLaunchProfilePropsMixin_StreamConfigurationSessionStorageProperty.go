package mixinsawsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamConfigurationSessionStorageProperty := &StreamConfigurationSessionStorageProperty{
//   	Mode: []*string{
//   		jsii.String("mode"),
//   	},
//   	Root: &StreamingSessionStorageRootProperty{
//   		Linux: jsii.String("linux"),
//   		Windows: jsii.String("windows"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionstorage.html
//
type CfnLaunchProfilePropsMixin_StreamConfigurationSessionStorageProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionstorage.html#cfn-nimblestudio-launchprofile-streamconfigurationsessionstorage-mode
	//
	Mode *[]*string `field:"optional" json:"mode" yaml:"mode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionstorage.html#cfn-nimblestudio-launchprofile-streamconfigurationsessionstorage-root
	//
	Root interface{} `field:"optional" json:"root" yaml:"root"`
}

