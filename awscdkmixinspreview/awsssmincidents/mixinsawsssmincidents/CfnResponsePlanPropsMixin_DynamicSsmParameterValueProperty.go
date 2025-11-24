package mixinsawsssmincidents


// The dynamic parameter value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamicSsmParameterValueProperty := &DynamicSsmParameterValueProperty{
//   	Variable: jsii.String("variable"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-dynamicssmparametervalue.html
//
type CfnResponsePlanPropsMixin_DynamicSsmParameterValueProperty struct {
	// Variable dynamic parameters.
	//
	// A parameter value is determined when an incident is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-dynamicssmparametervalue.html#cfn-ssmincidents-responseplan-dynamicssmparametervalue-variable
	//
	Variable *string `field:"optional" json:"variable" yaml:"variable"`
}

