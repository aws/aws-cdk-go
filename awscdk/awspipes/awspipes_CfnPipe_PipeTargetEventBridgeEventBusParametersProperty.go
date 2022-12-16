package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetEventBridgeEventBusParametersProperty := &pipeTargetEventBridgeEventBusParametersProperty{
//   	detailType: jsii.String("detailType"),
//   	endpointId: jsii.String("endpointId"),
//   	resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	source: jsii.String("source"),
//   	time: jsii.String("time"),
//   }
//
type CfnPipe_PipeTargetEventBridgeEventBusParametersProperty struct {
	// `CfnPipe.PipeTargetEventBridgeEventBusParametersProperty.DetailType`.
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// `CfnPipe.PipeTargetEventBridgeEventBusParametersProperty.EndpointId`.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// `CfnPipe.PipeTargetEventBridgeEventBusParametersProperty.Resources`.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// `CfnPipe.PipeTargetEventBridgeEventBusParametersProperty.Source`.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// `CfnPipe.PipeTargetEventBridgeEventBusParametersProperty.Time`.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

