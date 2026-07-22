package awsconnect


// The SLA configuration for cases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caseSlaConfigurationProperty := &CaseSlaConfigurationProperty{
//   	Name: jsii.String("name"),
//   	TargetSlaMinutes: jsii.Number(123),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	FieldId: jsii.String("fieldId"),
//   	TargetFieldValues: []interface{}{
//   		&SlaTargetFieldValueProperty{
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html
//
type CfnRule_CaseSlaConfigurationProperty struct {
	// The name of the SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The target SLA time in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-targetslaminutes
	//
	TargetSlaMinutes *float64 `field:"required" json:"targetSlaMinutes" yaml:"targetSlaMinutes"`
	// The type of SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The field Id for the SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The target field values for the SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-targetfieldvalues
	//
	TargetFieldValues interface{} `field:"optional" json:"targetFieldValues" yaml:"targetFieldValues"`
}

