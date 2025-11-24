package mixinsawsdeadline


// The details of the Amazon EC2 instance market options for a service managed fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceManagedEc2InstanceMarketOptionsProperty := &ServiceManagedEc2InstanceMarketOptionsProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancemarketoptions.html
//
type CfnFleetPropsMixin_ServiceManagedEc2InstanceMarketOptionsProperty struct {
	// The Amazon EC2 instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancemarketoptions.html#cfn-deadline-fleet-servicemanagedec2instancemarketoptions-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

