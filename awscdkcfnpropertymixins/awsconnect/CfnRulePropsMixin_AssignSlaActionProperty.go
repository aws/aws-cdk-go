package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   assignSlaActionProperty := &AssignSlaActionProperty{
//   	CaseSlaConfiguration: &CaseSlaConfigurationProperty{
//   		FieldId: jsii.String("fieldId"),
//   		Name: jsii.String("name"),
//   		TargetFieldValues: []interface{}{
//   			&SlaTargetFieldValueProperty{
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		TargetSlaMinutes: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	SlaAssignmentType: jsii.String("slaAssignmentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-assignslaaction.html
//
type CfnRulePropsMixin_AssignSlaActionProperty struct {
	// The SLA configuration for cases.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-assignslaaction.html#cfn-connect-rule-assignslaaction-caseslaconfiguration
	//
	CaseSlaConfiguration interface{} `field:"optional" json:"caseSlaConfiguration" yaml:"caseSlaConfiguration"`
	// The type of SLA assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-assignslaaction.html#cfn-connect-rule-assignslaaction-slaassignmenttype
	//
	SlaAssignmentType *string `field:"optional" json:"slaAssignmentType" yaml:"slaAssignmentType"`
}

