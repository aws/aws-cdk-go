package awsinterconnect


// The logical attachment point in your AWS network where the managed connection will be connected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attachPointProperty := &AttachPointProperty{
//   	Arn: jsii.String("arn"),
//   	DirectConnectGateway: jsii.String("directConnectGateway"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-attachpoint.html
//
type CfnConnection_AttachPointProperty struct {
	// The ARN of the resource to attach to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-attachpoint.html#cfn-interconnect-connection-attachpoint-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The ID of the Direct Connect Gateway to attach to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-attachpoint.html#cfn-interconnect-connection-attachpoint-directconnectgateway
	//
	DirectConnectGateway *string `field:"optional" json:"directConnectGateway" yaml:"directConnectGateway"`
}

