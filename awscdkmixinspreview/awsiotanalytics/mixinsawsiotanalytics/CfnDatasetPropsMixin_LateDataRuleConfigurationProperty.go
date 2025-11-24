package mixinsawsiotanalytics


// The information needed to configure a delta time session window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lateDataRuleConfigurationProperty := &LateDataRuleConfigurationProperty{
//   	DeltaTimeSessionWindowConfiguration: &DeltaTimeSessionWindowConfigurationProperty{
//   		TimeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedataruleconfiguration.html
//
type CfnDatasetPropsMixin_LateDataRuleConfigurationProperty struct {
	// The information needed to configure a delta time session window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedataruleconfiguration.html#cfn-iotanalytics-dataset-latedataruleconfiguration-deltatimesessionwindowconfiguration
	//
	DeltaTimeSessionWindowConfiguration interface{} `field:"optional" json:"deltaTimeSessionWindowConfiguration" yaml:"deltaTimeSessionWindowConfiguration"`
}

