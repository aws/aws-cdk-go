package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOrganizationalUnit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationalUnitProps := &cfnOrganizationalUnitProps{
//   	name: jsii.String("name"),
//   	parentId: jsii.String("parentId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnOrganizationalUnitProps struct {
	// `AWS::Organizations::OrganizationalUnit.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Organizations::OrganizationalUnit.ParentId`.
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// `AWS::Organizations::OrganizationalUnit.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

