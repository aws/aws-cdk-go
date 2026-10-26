package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTenant`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTenantProps := &CfnTenantProps{
//   	TenantName: jsii.String("tenantName"),
//
//   	// the properties below are optional
//   	ReputationEntity: &ReputationEntityProperty{
//   		Policy: &ReputationPolicyProperty{
//   			ReputationEntityPolicy: jsii.String("reputationEntityPolicy"),
//   		},
//   	},
//   	ResourceAssociations: []interface{}{
//   		&ResourceAssociationProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	SuppressionAttributes: &TenantSuppressionAttributesProperty{
//   		SuppressedReasons: []*string{
//   			jsii.String("suppressedReasons"),
//   		},
//   		SuppressionScope: jsii.String("suppressionScope"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-tenant.html
//
type CfnTenantProps struct {
	// The name of a tenant.
	//
	// The name can contain up to 64 alphanumeric characters, including letters, numbers, hyphens (-) and underscores (_) only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-tenant.html#cfn-ses-tenant-tenantname
	//
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
	// The reputation configuration for the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-tenant.html#cfn-ses-tenant-reputationentity
	//
	ReputationEntity interface{} `field:"optional" json:"reputationEntity" yaml:"reputationEntity"`
	// The list of resources to associate with the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-tenant.html#cfn-ses-tenant-resourceassociations
	//
	ResourceAssociations interface{} `field:"optional" json:"resourceAssociations" yaml:"resourceAssociations"`
	// The suppression attributes for the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-tenant.html#cfn-ses-tenant-suppressionattributes
	//
	SuppressionAttributes interface{} `field:"optional" json:"suppressionAttributes" yaml:"suppressionAttributes"`
	// An array of objects that define the tags (keys and values) associated with the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-tenant.html#cfn-ses-tenant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

