package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   availabilityZoneAddressProperty := &AvailabilityZoneAddressProperty{
//   	AllocationIds: []*string{
//   		jsii.String("allocationIds"),
//   	},
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html
//
type CfnNatGatewayPropsMixin_AvailabilityZoneAddressProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html#cfn-ec2-natgateway-availabilityzoneaddress-allocationids
	//
	AllocationIds *[]*string `field:"optional" json:"allocationIds" yaml:"allocationIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html#cfn-ec2-natgateway-availabilityzoneaddress-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html#cfn-ec2-natgateway-availabilityzoneaddress-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
}

