package awsecs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-thresholdconfiguration.html#cfn-ecs-service-thresholdconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-thresholdconfiguration.html#cfn-ecs-service-thresholdconfiguration-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

