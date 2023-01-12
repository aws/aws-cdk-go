package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionParameterProperty := &conditionParameterProperty{
//   	conditionKey: jsii.String("conditionKey"),
//   	conditionValue: jsii.String("conditionValue"),
//   }
//
type CfnBackupSelection_ConditionParameterProperty struct {
	// `CfnBackupSelection.ConditionParameterProperty.ConditionKey`.
	ConditionKey *string `field:"optional" json:"conditionKey" yaml:"conditionKey"`
	// `CfnBackupSelection.ConditionParameterProperty.ConditionValue`.
	ConditionValue *string `field:"optional" json:"conditionValue" yaml:"conditionValue"`
}

