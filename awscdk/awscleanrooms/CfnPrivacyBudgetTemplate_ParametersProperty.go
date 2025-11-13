package awscleanrooms


// Specifies the epsilon and noise parameters for the privacy budget template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &ParametersProperty{
//   	BudgetParameters: []interface{}{
//   		&BudgetParameterProperty{
//   			Budget: jsii.Number(123),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			AutoRefresh: jsii.String("autoRefresh"),
//   		},
//   	},
//   	Epsilon: jsii.Number(123),
//   	ResourceArn: jsii.String("resourceArn"),
//   	UsersNoisePerQuery: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html
//
type CfnPrivacyBudgetTemplate_ParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html#cfn-cleanrooms-privacybudgettemplate-parameters-budgetparameters
	//
	BudgetParameters interface{} `field:"optional" json:"budgetParameters" yaml:"budgetParameters"`
	// The epsilon value that you want to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html#cfn-cleanrooms-privacybudgettemplate-parameters-epsilon
	//
	Epsilon *float64 `field:"optional" json:"epsilon" yaml:"epsilon"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html#cfn-cleanrooms-privacybudgettemplate-parameters-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// Noise added per query is measured in terms of the number of users whose contributions you want to obscure.
	//
	// This value governs the rate at which the privacy budget is depleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-parameters.html#cfn-cleanrooms-privacybudgettemplate-parameters-usersnoiseperquery
	//
	UsersNoisePerQuery *float64 `field:"optional" json:"usersNoisePerQuery" yaml:"usersNoisePerQuery"`
}

