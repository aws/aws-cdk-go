package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eventSourceMappingProperty := &EventSourceMappingProperty{
//   	Arn: jsii.String("arn"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eventsourcemapping.html
//
type CfnPlanPropsMixin_EventSourceMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eventsourcemapping.html#cfn-arcregionswitch-plan-eventsourcemapping-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eventsourcemapping.html#cfn-arcregionswitch-plan-eventsourcemapping-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eventsourcemapping.html#cfn-arcregionswitch-plan-eventsourcemapping-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

