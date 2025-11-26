package previewawsdeadlinemixins


// The amounts and attributes of fleets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetCapabilitiesProperty := &FleetCapabilitiesProperty{
//   	Amounts: []interface{}{
//   		&FleetAmountCapabilityProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Attributes: []interface{}{
//   		&FleetAttributeCapabilityProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetcapabilities.html
//
type CfnFleetPropsMixin_FleetCapabilitiesProperty struct {
	// Amount capabilities of the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetcapabilities.html#cfn-deadline-fleet-fleetcapabilities-amounts
	//
	Amounts interface{} `field:"optional" json:"amounts" yaml:"amounts"`
	// Attribute capabilities of the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetcapabilities.html#cfn-deadline-fleet-fleetcapabilities-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

