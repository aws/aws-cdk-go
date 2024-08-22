package awssystemsmanagersap

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ApplicationType: jsii.String("applicationType"),
//
//   	// the properties below are optional
//   	Credentials: []interface{}{
//   		&CredentialProperty{
//   			CredentialType: jsii.String("credentialType"),
//   			DatabaseName: jsii.String("databaseName"),
//   			SecretId: jsii.String("secretId"),
//   		},
//   	},
//   	DatabaseArn: jsii.String("databaseArn"),
//   	Instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	SapInstanceNumber: jsii.String("sapInstanceNumber"),
//   	Sid: jsii.String("sid"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html
//
type CfnApplicationProps struct {
	// The ID of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The type of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-applicationtype
	//
	ApplicationType *string `field:"required" json:"applicationType" yaml:"applicationType"`
	// The credentials of the SAP application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The Amazon Resource Name (ARN) of the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-databasearn
	//
	DatabaseArn *string `field:"optional" json:"databaseArn" yaml:"databaseArn"`
	// The Amazon EC2 instances on which your SAP application is running.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-instances
	//
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// The SAP instance number of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-sapinstancenumber
	//
	SapInstanceNumber *string `field:"optional" json:"sapInstanceNumber" yaml:"sapInstanceNumber"`
	// The System ID of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-sid
	//
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
	// The tags on the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-systemsmanagersap-application.html#cfn-systemsmanagersap-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

