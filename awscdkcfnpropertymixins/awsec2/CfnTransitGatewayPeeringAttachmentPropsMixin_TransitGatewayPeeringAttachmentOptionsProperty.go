package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transitGatewayPeeringAttachmentOptionsProperty := &TransitGatewayPeeringAttachmentOptionsProperty{
//   	AllowMultipathResilience: jsii.Boolean(false),
//   	MultipathResilienceMode: jsii.String("multipathResilienceMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypeeringattachment-transitgatewaypeeringattachmentoptions.html
//
type CfnTransitGatewayPeeringAttachmentPropsMixin_TransitGatewayPeeringAttachmentOptionsProperty struct {
	// Whether to enable MultipathResilience on this side of the peering attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypeeringattachment-transitgatewaypeeringattachmentoptions.html#cfn-ec2-transitgatewaypeeringattachment-transitgatewaypeeringattachmentoptions-allowmultipathresilience
	//
	AllowMultipathResilience interface{} `field:"optional" json:"allowMultipathResilience" yaml:"allowMultipathResilience"`
	// The MultipathResilience mode.
	//
	// Only valid when AllowMultipathResilience is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypeeringattachment-transitgatewaypeeringattachmentoptions.html#cfn-ec2-transitgatewaypeeringattachment-transitgatewaypeeringattachmentoptions-multipathresiliencemode
	//
	MultipathResilienceMode *string `field:"optional" json:"multipathResilienceMode" yaml:"multipathResilienceMode"`
}

