package awsgameliftstreams


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcTransitConfigurationProperty := &VpcTransitConfigurationProperty{
//   	Ipv4CidrBlocks: []*string{
//   		jsii.String("ipv4CidrBlocks"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-vpctransitconfiguration.html
//
type CfnStreamGroup_VpcTransitConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-vpctransitconfiguration.html#cfn-gameliftstreams-streamgroup-vpctransitconfiguration-ipv4cidrblocks
	//
	Ipv4CidrBlocks *[]*string `field:"required" json:"ipv4CidrBlocks" yaml:"ipv4CidrBlocks"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-vpctransitconfiguration.html#cfn-gameliftstreams-streamgroup-vpctransitconfiguration-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

