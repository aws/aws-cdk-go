package previewawsdeadlinemixins


// Defines the fleet's capability name, minimum, and maximum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetAttributeCapabilityProperty := &FleetAttributeCapabilityProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetattributecapability.html
//
type CfnFleetPropsMixin_FleetAttributeCapabilityProperty struct {
	// The name of the fleet attribute capability for the worker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetattributecapability.html#cfn-deadline-fleet-fleetattributecapability-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of fleet attribute capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetattributecapability.html#cfn-deadline-fleet-fleetattributecapability-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

