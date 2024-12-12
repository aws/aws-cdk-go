package awsmedialive


// Settings that apply only if the input is a push type of input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDestinationRequestProperty := &InputDestinationRequestProperty{
//   	Network: jsii.String("network"),
//   	NetworkRoutes: []interface{}{
//   		&InputRequestDestinationRouteProperty{
//   			Cidr: jsii.String("cidr"),
//   			Gateway: jsii.String("gateway"),
//   		},
//   	},
//   	StaticIpAddress: jsii.String("staticIpAddress"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html
//
type CfnInput_InputDestinationRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html#cfn-medialive-input-inputdestinationrequest-network
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html#cfn-medialive-input-inputdestinationrequest-networkroutes
	//
	NetworkRoutes interface{} `field:"optional" json:"networkRoutes" yaml:"networkRoutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html#cfn-medialive-input-inputdestinationrequest-staticipaddress
	//
	StaticIpAddress *string `field:"optional" json:"staticIpAddress" yaml:"staticIpAddress"`
	// The stream name (application name/application instance) for the location the RTMP source content will be pushed to in MediaLive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html#cfn-medialive-input-inputdestinationrequest-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

