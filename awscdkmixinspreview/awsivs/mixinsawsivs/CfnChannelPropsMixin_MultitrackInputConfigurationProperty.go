package mixinsawsivs


// A complex type that specifies multitrack input configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multitrackInputConfigurationProperty := &MultitrackInputConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	MaximumResolution: jsii.String("maximumResolution"),
//   	Policy: jsii.String("policy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-channel-multitrackinputconfiguration.html
//
type CfnChannelPropsMixin_MultitrackInputConfigurationProperty struct {
	// Indicates whether multitrack input is enabled.
	//
	// Can be set to true only if channel type is STANDARD. Setting enabled to true with any other channel type will cause an exception. If true, then policy, maximumResolution, and containerFormat are required, and containerFormat must be set to FRAGMENTED_MP4. Default: false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-channel-multitrackinputconfiguration.html#cfn-ivs-channel-multitrackinputconfiguration-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Maximum resolution for multitrack input.
	//
	// Required if enabled is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-channel-multitrackinputconfiguration.html#cfn-ivs-channel-multitrackinputconfiguration-maximumresolution
	//
	MaximumResolution *string `field:"optional" json:"maximumResolution" yaml:"maximumResolution"`
	// Indicates whether multitrack input is allowed or required.
	//
	// Required if enabled is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-channel-multitrackinputconfiguration.html#cfn-ivs-channel-multitrackinputconfiguration-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

