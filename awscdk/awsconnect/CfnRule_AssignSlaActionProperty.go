package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assignSlaActionProperty := &AssignSlaActionProperty{
//   	CaseSlaConfiguration: &CaseSlaConfigurationProperty{
//   		Name: jsii.String("name"),
//   		TargetSlaMinutes: jsii.Number(123),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		FieldId: jsii.String("fieldId"),
//   		TargetFieldValues: []interface{}{
//   			&SlaTargetFieldValueProperty{
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	SlaAssignmentType: jsii.String("slaAssignmentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-assignslaaction.html
//
type CfnRule_AssignSlaActionProperty struct {
	// The SLA configuration for cases.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-assignslaaction.html#cfn-connect-rule-assignslaaction-caseslaconfiguration
	//
	CaseSlaConfiguration interface{} `field:"required" json:"caseSlaConfiguration" yaml:"caseSlaConfiguration"`
	// The type of SLA assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-assignslaaction.html#cfn-connect-rule-assignslaaction-slaassignmenttype
	//
	SlaAssignmentType *string `field:"required" json:"slaAssignmentType" yaml:"slaAssignmentType"`
}

