package awseks


// Configuration settings for an ACK (AWS Controllers for Kubernetes) capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ackProperty := &AckProperty{
//   	DisabledServices: []*string{
//   		jsii.String("disabledServices"),
//   	},
//   	EnableCrossNamespace: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ack.html
//
type CfnCapabilityPropsMixin_AckProperty struct {
	// A list of ACK service names to disable.
	//
	// Controllers for services in this list are not installed or managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ack.html#cfn-eks-capability-ack-disabledservices
	//
	DisabledServices *[]*string `field:"optional" json:"disabledServices" yaml:"disabledServices"`
	// Whether cross-namespace references are enabled for ACK controllers.
	//
	// When not specified, the service default applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ack.html#cfn-eks-capability-ack-enablecrossnamespace
	//
	EnableCrossNamespace interface{} `field:"optional" json:"enableCrossNamespace" yaml:"enableCrossNamespace"`
}

