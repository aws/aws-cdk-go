package previewawsroute53mixins


// Represents the features configuration for a hosted zone, including the status of various features and any associated failure reasons.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hostedZoneFeaturesProperty := &HostedZoneFeaturesProperty{
//   	EnableAcceleratedRecovery: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonefeatures.html
//
type CfnHostedZonePropsMixin_HostedZoneFeaturesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonefeatures.html#cfn-route53-hostedzone-hostedzonefeatures-enableacceleratedrecovery
	//
	EnableAcceleratedRecovery interface{} `field:"optional" json:"enableAcceleratedRecovery" yaml:"enableAcceleratedRecovery"`
}

