package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A framework consists of one or more controls.
//
// Each control has its own control scope. The control scope can include one or more resource types, a combination of a tag key and value, or a combination of one resource type and one resource ID. If no scope is specified, evaluations for the rule are triggered when any resource in your recording group changes in configuration.
//
// > To set a control scope that includes all of a particular resource, leave the `ControlScope` empty or do not pass it when calling `CreateFramework` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlScopeProperty := &ControlScopeProperty{
//   	ComplianceResourceIds: []*string{
//   		jsii.String("complianceResourceIds"),
//   	},
//   	ComplianceResourceTypes: []*string{
//   		jsii.String("complianceResourceTypes"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFramework_ControlScopeProperty struct {
	// The ID of the only AWS resource that you want your control scope to contain.
	ComplianceResourceIds *[]*string `field:"optional" json:"complianceResourceIds" yaml:"complianceResourceIds"`
	// Describes whether the control scope includes one or more types of resources, such as `EFS` or `RDS` .
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// The tag key-value pair applied to those AWS resources that you want to trigger an evaluation for a rule.
	//
	// A maximum of one key-value pair can be provided. The tag value is optional, but it cannot be an empty string. The structure to assign a tag is: `[{"Key":"string","Value":"string"}]` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

