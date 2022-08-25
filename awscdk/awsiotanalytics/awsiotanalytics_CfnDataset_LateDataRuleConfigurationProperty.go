package awsiotanalytics


// The information needed to configure a delta time session window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lateDataRuleConfigurationProperty := &lateDataRuleConfigurationProperty{
//   	deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   		timeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnDataset_LateDataRuleConfigurationProperty struct {
	// The information needed to configure a delta time session window.
	DeltaTimeSessionWindowConfiguration interface{} `field:"optional" json:"deltaTimeSessionWindowConfiguration" yaml:"deltaTimeSessionWindowConfiguration"`
}

