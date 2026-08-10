package awsauditmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAssessmentFrameworkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAssessmentFrameworkMixinProps := &CfnAssessmentFrameworkMixinProps{
//   	ComplianceType: jsii.String("complianceType"),
//   	ControlSets: []interface{}{
//   		&ControlSetProperty{
//   			Controls: []interface{}{
//   				&ControlSetControlProperty{
//   					Id: jsii.String("id"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html
//
type CfnAssessmentFrameworkMixinProps struct {
	// The compliance type that the framework supports, such as CIS or HIPAA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-compliancetype
	//
	ComplianceType *string `field:"optional" json:"complianceType" yaml:"complianceType"`
	// The control sets that are associated with the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-controlsets
	//
	ControlSets interface{} `field:"optional" json:"controlSets" yaml:"controlSets"`
	// The description of the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags associated with the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

