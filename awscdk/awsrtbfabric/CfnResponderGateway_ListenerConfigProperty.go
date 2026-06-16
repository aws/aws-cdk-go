package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerConfigProperty := &ListenerConfigProperty{
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-listenerconfig.html
//
type CfnResponderGateway_ListenerConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-listenerconfig.html#cfn-rtbfabric-respondergateway-listenerconfig-protocols
	//
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
}

