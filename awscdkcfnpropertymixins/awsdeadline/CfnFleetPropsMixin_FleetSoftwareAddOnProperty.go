package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fleetSoftwareAddOnProperty := &FleetSoftwareAddOnProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetsoftwareaddon.html
//
type CfnFleetPropsMixin_FleetSoftwareAddOnProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetsoftwareaddon.html#cfn-deadline-fleet-fleetsoftwareaddon-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

