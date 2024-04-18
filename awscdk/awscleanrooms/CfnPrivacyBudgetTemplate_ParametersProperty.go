package awscleanrooms


// Specifies the epislon and noise parameters for the privacy budget template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &ParametersProperty{
//   	Epsilon: jsii.Number(123),
//   	UsersNoisePerQuery: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html
//
type CfnPrivacyBudgetTemplate_ParametersProperty struct {
	// The epsilon value that you want to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html#cfn-cleanrooms-privacybudgettemplate-parameters-epsilon
	//
	Epsilon *float64 `field:"required" json:"epsilon" yaml:"epsilon"`
	// Noise added per query is measured in terms of the number of users whose contributions you want to obscure.
	//
	// This value governs the rate at which the privacy budget is depleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html#cfn-cleanrooms-privacybudgettemplate-parameters-usersnoiseperquery
	//
	UsersNoisePerQuery *float64 `field:"required" json:"usersNoisePerQuery" yaml:"usersNoisePerQuery"`
}

