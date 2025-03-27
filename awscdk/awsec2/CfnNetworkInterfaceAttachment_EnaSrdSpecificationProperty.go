package awsec2


// ENA Express uses AWS Scalable Reliable Datagram (SRD) technology to increase the maximum bandwidth used per stream and minimize tail latency of network traffic between EC2 instances.
//
// With ENA Express, you can communicate between two EC2 instances in the same subnet within the same account, or in different accounts. Both sending and receiving instances must have ENA Express enabled.
//
// To improve the reliability of network packet delivery, ENA Express reorders network packets on the receiving end by default. However, some UDP-based applications are designed to handle network packets that are out of order to reduce the overhead for packet delivery at the network layer. When ENA Express is enabled, you can specify whether UDP network traffic uses it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enaSrdSpecificationProperty := &EnaSrdSpecificationProperty{
//   	EnaSrdEnabled: jsii.Boolean(false),
//   	EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   		EnaSrdUdpEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdspecification.html
//
type CfnNetworkInterfaceAttachment_EnaSrdSpecificationProperty struct {
	// Indicates whether ENA Express is enabled for the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdspecification.html#cfn-ec2-networkinterfaceattachment-enasrdspecification-enasrdenabled
	//
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// Configures ENA Express for UDP network traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdspecification.html#cfn-ec2-networkinterfaceattachment-enasrdspecification-enasrdudpspecification
	//
	EnaSrdUdpSpecification interface{} `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

