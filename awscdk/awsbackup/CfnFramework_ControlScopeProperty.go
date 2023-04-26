package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

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
	// `CfnFramework.ControlScopeProperty.ComplianceResourceIds`.
	ComplianceResourceIds *[]*string `field:"optional" json:"complianceResourceIds" yaml:"complianceResourceIds"`
	// `CfnFramework.ControlScopeProperty.ComplianceResourceTypes`.
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// `CfnFramework.ControlScopeProperty.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

