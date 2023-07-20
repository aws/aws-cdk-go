package awsconfig


// The static value of the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticValueProperty := &StaticValueProperty{
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-staticvalue.html
//
type CfnRemediationConfiguration_StaticValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-staticvalue.html#cfn-config-remediationconfiguration-staticvalue-value
	//
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
	// A list of values.
	//
	// For example, the ARN of the assumed role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-staticvalue.html#cfn-config-remediationconfiguration-staticvalue-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

