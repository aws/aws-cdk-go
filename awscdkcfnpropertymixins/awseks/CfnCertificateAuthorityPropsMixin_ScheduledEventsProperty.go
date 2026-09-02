package awseks


// The scheduled auto-activation events for the certificate authority, computed from its validity window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scheduledEventsProperty := &ScheduledEventsProperty{
//   	FinalAutoActivation: jsii.String("finalAutoActivation"),
//   	FirstAutoActivation: jsii.String("firstAutoActivation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-certificateauthority-scheduledevents.html
//
type CfnCertificateAuthorityPropsMixin_ScheduledEventsProperty struct {
	// The deadline by which EKS will auto-activate this certificate authority (notAfter minus 45 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-certificateauthority-scheduledevents.html#cfn-eks-certificateauthority-scheduledevents-finalautoactivation
	//
	FinalAutoActivation *string `field:"optional" json:"finalAutoActivation" yaml:"finalAutoActivation"`
	// The earliest date EKS may auto-activate this certificate authority (notAfter minus 6 months).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-certificateauthority-scheduledevents.html#cfn-eks-certificateauthority-scheduledevents-firstautoactivation
	//
	FirstAutoActivation *string `field:"optional" json:"firstAutoActivation" yaml:"firstAutoActivation"`
}

