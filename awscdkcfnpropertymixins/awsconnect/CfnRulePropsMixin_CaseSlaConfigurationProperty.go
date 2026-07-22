package awsconnect


// The SLA configuration for cases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   caseSlaConfigurationProperty := &CaseSlaConfigurationProperty{
//   	FieldId: jsii.String("fieldId"),
//   	Name: jsii.String("name"),
//   	TargetFieldValues: []interface{}{
//   		&SlaTargetFieldValueProperty{
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	TargetSlaMinutes: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html
//
type CfnRulePropsMixin_CaseSlaConfigurationProperty struct {
	// The field Id for the SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The name of the SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target field values for the SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-targetfieldvalues
	//
	TargetFieldValues interface{} `field:"optional" json:"targetFieldValues" yaml:"targetFieldValues"`
	// The target SLA time in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-targetslaminutes
	//
	TargetSlaMinutes *float64 `field:"optional" json:"targetSlaMinutes" yaml:"targetSlaMinutes"`
	// The type of SLA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-caseslaconfiguration.html#cfn-connect-rule-caseslaconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

