package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionsProperty := &ConditionsProperty{
//   	StringEquals: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	StringLike: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	StringNotEquals: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	StringNotLike: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   }
//
type CfnBackupSelection_ConditionsProperty struct {
	// `CfnBackupSelection.ConditionsProperty.StringEquals`.
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// `CfnBackupSelection.ConditionsProperty.StringLike`.
	StringLike interface{} `field:"optional" json:"stringLike" yaml:"stringLike"`
	// `CfnBackupSelection.ConditionsProperty.StringNotEquals`.
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
	// `CfnBackupSelection.ConditionsProperty.StringNotLike`.
	StringNotLike interface{} `field:"optional" json:"stringNotLike" yaml:"stringNotLike"`
}

