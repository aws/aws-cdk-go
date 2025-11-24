package mixinsawsemrserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   initialCapacityConfigKeyValuePairProperty := &InitialCapacityConfigKeyValuePairProperty{
//   	Key: jsii.String("key"),
//   	Value: &InitialCapacityConfigProperty{
//   		WorkerConfiguration: &WorkerConfigurationProperty{
//   			Cpu: jsii.String("cpu"),
//   			Disk: jsii.String("disk"),
//   			DiskType: jsii.String("diskType"),
//   			Memory: jsii.String("memory"),
//   		},
//   		WorkerCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.html
//
type CfnApplicationPropsMixin_InitialCapacityConfigKeyValuePairProperty struct {
	// Worker type for an analytics framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.html#cfn-emrserverless-application-initialcapacityconfigkeyvaluepair-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.html#cfn-emrserverless-application-initialcapacityconfigkeyvaluepair-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

