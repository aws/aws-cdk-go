package mixinsawsqbusiness


// Properties for CfnPermissionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPermissionMixinProps := &CfnPermissionMixinProps{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ApplicationId: jsii.String("applicationId"),
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionOperator: jsii.String("conditionOperator"),
//   			ConditionValues: []*string{
//   				jsii.String("conditionValues"),
//   			},
//   		},
//   	},
//   	Principal: jsii.String("principal"),
//   	StatementId: jsii.String("statementId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html
//
type CfnPermissionMixinProps struct {
	// The list of Amazon Q Business actions that the ISV is allowed to perform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// The unique identifier of the Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Provides user and group information used for filtering documents to use for generating Amazon Q Business conversation responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// A unique identifier for the policy statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-statementid
	//
	StatementId *string `field:"optional" json:"statementId" yaml:"statementId"`
}

