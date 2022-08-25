package awsssmincidents


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicSsmParameterProperty := &dynamicSsmParameterProperty{
//   	key: jsii.String("key"),
//   	value: &dynamicSsmParameterValueProperty{
//   		variable: jsii.String("variable"),
//   	},
//   }
//
type CfnResponsePlan_DynamicSsmParameterProperty struct {
	// `CfnResponsePlan.DynamicSsmParameterProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnResponsePlan.DynamicSsmParameterProperty.Value`.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

