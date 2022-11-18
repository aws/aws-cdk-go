package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &cfnAccountProps{
//   	accountName: jsii.String("accountName"),
//   	email: jsii.String("email"),
//
//   	// the properties below are optional
//   	parentIds: []*string{
//   		jsii.String("parentIds"),
//   	},
//   	roleName: jsii.String("roleName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAccountProps struct {
	// `AWS::Organizations::Account.AccountName`.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// `AWS::Organizations::Account.Email`.
	Email *string `field:"required" json:"email" yaml:"email"`
	// `AWS::Organizations::Account.ParentIds`.
	ParentIds *[]*string `field:"optional" json:"parentIds" yaml:"parentIds"`
	// `AWS::Organizations::Account.RoleName`.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// `AWS::Organizations::Account.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

