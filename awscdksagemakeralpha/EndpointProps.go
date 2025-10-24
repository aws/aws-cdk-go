package awscdksagemakeralpha


// Construction properties for a SageMaker Endpoint.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var endpointConfig EndpointConfig
//
//
//   endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &EndpointProps{
//   	EndpointConfig: EndpointConfig,
//   })
//   productionVariant := endpoint.FindInstanceProductionVariant(jsii.String("my-variant"))
//   productionVariant.MetricModelLatency().CreateAlarm(this, jsii.String("ModelLatencyAlarm"), &CreateAlarmOptions{
//   	Threshold: jsii.Number(100000),
//   	EvaluationPeriods: jsii.Number(3),
//   })
//
// Experimental.
type EndpointProps struct {
	// The endpoint configuration to use for this endpoint.
	// Experimental.
	EndpointConfig IEndpointConfig `field:"required" json:"endpointConfig" yaml:"endpointConfig"`
	// Name of the endpoint.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that ID for the
	// endpoint's name.
	//
	// Experimental.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
}

