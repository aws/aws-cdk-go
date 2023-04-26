package awsssmincidents


// The dynamic parameter value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicSsmParameterValueProperty := &DynamicSsmParameterValueProperty{
//   	Variable: jsii.String("variable"),
//   }
//
type CfnResponsePlan_DynamicSsmParameterValueProperty struct {
	// Variable dynamic parameters.
	//
	// A parameter value is determined when an incident is created.
	Variable *string `field:"optional" json:"variable" yaml:"variable"`
}

