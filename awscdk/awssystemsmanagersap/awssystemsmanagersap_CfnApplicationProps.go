package awssystemsmanagersap

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	applicationId: jsii.String("applicationId"),
//   	applicationType: jsii.String("applicationType"),
//
//   	// the properties below are optional
//   	credentials: []interface{}{
//   		&credentialProperty{
//   			credentialType: jsii.String("credentialType"),
//   			databaseName: jsii.String("databaseName"),
//   			secretId: jsii.String("secretId"),
//   		},
//   	},
//   	instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	sapInstanceNumber: jsii.String("sapInstanceNumber"),
//   	sid: jsii.String("sid"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// The ID of the application.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The type of the application.
	ApplicationType *string `field:"required" json:"applicationType" yaml:"applicationType"`
	// The credentials of the SAP application.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The Amazon EC2 instances on which your SAP application is running.
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// The SAP instance number of the application.
	SapInstanceNumber *string `field:"optional" json:"sapInstanceNumber" yaml:"sapInstanceNumber"`
	// The System ID of the application.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
	// The tags on the application.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

