package awsemrserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initialCapacityConfigKeyValuePairProperty := &InitialCapacityConfigKeyValuePairProperty{
//   	Key: jsii.String("key"),
//   	Value: &InitialCapacityConfigProperty{
//   		WorkerConfiguration: &WorkerConfigurationProperty{
//   			Cpu: jsii.String("cpu"),
//   			Memory: jsii.String("memory"),
//
//   			// the properties below are optional
//   			Disk: jsii.String("disk"),
//   		},
//   		WorkerCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.html
//
type CfnApplication_InitialCapacityConfigKeyValuePairProperty struct {
	// Worker type for an analytics framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.html#cfn-emrserverless-application-initialcapacityconfigkeyvaluepair-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.html#cfn-emrserverless-application-initialcapacityconfigkeyvaluepair-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

