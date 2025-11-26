package previewawsssmincidentsmixins


// The key-value pair parameters to use when running the Automation runbook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ssmParameterProperty := &SsmParameterProperty{
//   	Key: jsii.String("key"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmparameter.html
//
type CfnResponsePlanPropsMixin_SsmParameterProperty struct {
	// The key parameter to use when running the Automation runbook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmparameter.html#cfn-ssmincidents-responseplan-ssmparameter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value parameter to use when running the Automation runbook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmparameter.html#cfn-ssmincidents-responseplan-ssmparameter-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

