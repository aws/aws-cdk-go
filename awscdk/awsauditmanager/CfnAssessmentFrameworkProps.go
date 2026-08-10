package awsauditmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAssessmentFramework`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssessmentFrameworkProps := &CfnAssessmentFrameworkProps{
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
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ComplianceType: jsii.String("complianceType"),
//   	Description: jsii.String("description"),
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
type CfnAssessmentFrameworkProps struct {
	// The control sets that are associated with the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-controlsets
	//
	ControlSets interface{} `field:"required" json:"controlSets" yaml:"controlSets"`
	// The name of the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The compliance type that the framework supports, such as CIS or HIPAA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-compliancetype
	//
	ComplianceType *string `field:"optional" json:"complianceType" yaml:"complianceType"`
	// The description of the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessmentframework.html#cfn-auditmanager-assessmentframework-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

