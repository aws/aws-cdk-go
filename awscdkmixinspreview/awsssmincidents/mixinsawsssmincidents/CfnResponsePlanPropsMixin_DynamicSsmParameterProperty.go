package mixinsawsssmincidents


// When you add a runbook to a response plan, you can specify the parameters for the runbook to use at runtime.
//
// Response plans support parameters with both static and dynamic values. For static values, you enter the value when you define the parameter in the response plan. For dynamic values, the system determines the correct parameter value by collecting information from the incident. Incident Manager supports the following dynamic parameters:
//
// *Incident ARN*
//
// When Incident Manager creates an incident, the system captures the Amazon Resource Name (ARN) of the corresponding incident record and enters it for this parameter in the runbook.
//
// > This value can only be assigned to parameters of type `String` . If assigned to a parameter of any other type, the runbook fails to run.
//
// *Involved resources*
//
// When Incident Manager creates an incident, the system captures the ARNs of the resources involved in the incident. These resource ARNs are then assigned to this parameter in the runbook.
//
// > This value can only be assigned to parameters of type `StringList` . If assigned to a parameter of any other type, the runbook fails to run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamicSsmParameterProperty := &DynamicSsmParameterProperty{
//   	Key: jsii.String("key"),
//   	Value: &DynamicSsmParameterValueProperty{
//   		Variable: jsii.String("variable"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-dynamicssmparameter.html
//
type CfnResponsePlanPropsMixin_DynamicSsmParameterProperty struct {
	// The key parameter to use when running the Systems Manager Automation runbook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-dynamicssmparameter.html#cfn-ssmincidents-responseplan-dynamicssmparameter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The dynamic parameter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-dynamicssmparameter.html#cfn-ssmincidents-responseplan-dynamicssmparameter-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

