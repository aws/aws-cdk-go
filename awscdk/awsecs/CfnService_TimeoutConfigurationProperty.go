package awsecs


// An object that represents the timeout configurations for Service Connect.
//
// > If `idleTimeout` is set to a time that is less than `perRequestTimeout` , the connection will close when the `idleTimeout` is reached and not the `perRequestTimeout` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutConfigurationProperty := &TimeoutConfigurationProperty{
//   	IdleTimeoutSeconds: jsii.Number(123),
//   	PerRequestTimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-timeoutconfiguration.html
//
type CfnService_TimeoutConfigurationProperty struct {
	// The amount of time in seconds a connection will stay active while idle.
	//
	// A value of `0` can be set to disable `idleTimeout` .
	//
	// The `idleTimeout` default for `HTTP` / `HTTP2` / `GRPC` is 5 minutes.
	//
	// The `idleTimeout` default for `TCP` is 1 hour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-timeoutconfiguration.html#cfn-ecs-service-timeoutconfiguration-idletimeoutseconds
	//
	IdleTimeoutSeconds *float64 `field:"optional" json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
	// The amount of time waiting for the upstream to respond with a complete response per request.
	//
	// A value of `0` can be set to disable `perRequestTimeout` . `perRequestTimeout` can only be set if Service Connect `appProtocol` isn't `TCP` . Only `idleTimeout` is allowed for `TCP` `appProtocol` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-timeoutconfiguration.html#cfn-ecs-service-timeoutconfiguration-perrequesttimeoutseconds
	//
	PerRequestTimeoutSeconds *float64 `field:"optional" json:"perRequestTimeoutSeconds" yaml:"perRequestTimeoutSeconds"`
}

