package mixinsawsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionProperty := &ConditionProperty{
//   	ConditionKey: jsii.String("conditionKey"),
//   	ConditionOperator: jsii.String("conditionOperator"),
//   	ConditionValues: []*string{
//   		jsii.String("conditionValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-permission-condition.html
//
type CfnPermissionPropsMixin_ConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-permission-condition.html#cfn-qbusiness-permission-condition-conditionkey
	//
	ConditionKey *string `field:"optional" json:"conditionKey" yaml:"conditionKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-permission-condition.html#cfn-qbusiness-permission-condition-conditionoperator
	//
	ConditionOperator *string `field:"optional" json:"conditionOperator" yaml:"conditionOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-permission-condition.html#cfn-qbusiness-permission-condition-conditionvalues
	//
	ConditionValues *[]*string `field:"optional" json:"conditionValues" yaml:"conditionValues"`
}

