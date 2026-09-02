package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reservedCapacityFallbackOptionsRequestProperty := &ReservedCapacityFallbackOptionsRequestProperty{
//   	MarketTypes: []*string{
//   		jsii.String("marketTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityfallbackoptionsrequest.html
//
type CfnEC2FleetPropsMixin_ReservedCapacityFallbackOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityfallbackoptionsrequest.html#cfn-ec2-ec2fleet-reservedcapacityfallbackoptionsrequest-markettypes
	//
	MarketTypes *[]*string `field:"optional" json:"marketTypes" yaml:"marketTypes"`
}

