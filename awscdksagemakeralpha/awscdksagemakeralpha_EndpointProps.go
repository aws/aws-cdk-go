// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha


// Construction properties for a SageMaker Endpoint.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var endpointConfig endpointConfig
//
//
//   endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &endpointProps{
//   	endpointConfig: endpointConfig,
//   })
//   productionVariant := endpoint.findInstanceProductionVariant(jsii.String("my-variant"))
//   productionVariant.metricModelLatency().createAlarm(this, jsii.String("ModelLatencyAlarm"), &createAlarmOptions{
//   	threshold: jsii.Number(100000),
//   	evaluationPeriods: jsii.Number(3),
//   })
//
// Experimental.
type EndpointProps struct {
	// The endpoint configuration to use for this endpoint.
	// Experimental.
	EndpointConfig IEndpointConfig `field:"required" json:"endpointConfig" yaml:"endpointConfig"`
	// Name of the endpoint.
	// Experimental.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
}

