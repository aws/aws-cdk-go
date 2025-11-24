package mixinsawspinpointemail


// Enable or disable collection of reputation metrics for emails that you send using this configuration set in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   reputationOptionsProperty := &ReputationOptionsProperty{
//   	ReputationMetricsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-reputationoptions.html
//
type CfnConfigurationSetPropsMixin_ReputationOptionsProperty struct {
	// If `true` , tracking of reputation metrics is enabled for the configuration set.
	//
	// If `false` , tracking of reputation metrics is disabled for the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-reputationoptions.html#cfn-pinpointemail-configurationset-reputationoptions-reputationmetricsenabled
	//
	ReputationMetricsEnabled interface{} `field:"optional" json:"reputationMetricsEnabled" yaml:"reputationMetricsEnabled"`
}

