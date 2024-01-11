package awscodetest


// Properties for defining a `CfnSeries`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var runDefinition interface{}
//
//   cfnSeriesProps := &CfnSeriesProps{
//   	PersistentConfigurationId: jsii.String("persistentConfigurationId"),
//   	RunDefinition: runDefinition,
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-series.html
//
type CfnSeriesProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-series.html#cfn-codetest-series-persistentconfigurationid
	//
	PersistentConfigurationId *string `field:"required" json:"persistentConfigurationId" yaml:"persistentConfigurationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-series.html#cfn-codetest-series-rundefinition
	//
	RunDefinition interface{} `field:"required" json:"runDefinition" yaml:"runDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-series.html#cfn-codetest-series-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-series.html#cfn-codetest-series-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

