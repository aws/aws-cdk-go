package awsecs


// Defines the failure threshold that the deployment circuit breaker uses to monitor a deployment.
//
// The ``type`` and ``value`` together determine the number of task failures that are tolerated before the circuit breaker triggers.
//  By default, the threshold configuration uses a ``type`` of ``BOUNDED_PERCENT`` with a ``value`` of ``50``.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   thresholdConfigurationProperty := &ThresholdConfigurationProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-thresholdconfiguration.html
//
type CfnServicePropsMixin_ThresholdConfigurationProperty struct {
	// Determines how Amazon ECS uses ``value`` to calculate the failure threshold.
	//
	// For the percentage types (``BOUNDED_PERCENT`` and ``UNBOUNDED_PERCENT``), Amazon ECS multiplies ``value`` by the latest service desired count. For ``COUNT``, Amazon ECS uses ``value`` directly as the threshold. The default is ``BOUNDED_PERCENT``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-thresholdconfiguration.html#cfn-ecs-service-thresholdconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the integer that Amazon ECS uses to calculate the failure threshold.
	//
	// When ``type`` is ``COUNT``, this value is the failure threshold itself. When ``type`` is a percentage type, Amazon ECS multiplies this value by the latest service desired count to produce the failure threshold. The default is ``50``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-thresholdconfiguration.html#cfn-ecs-service-thresholdconfiguration-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

