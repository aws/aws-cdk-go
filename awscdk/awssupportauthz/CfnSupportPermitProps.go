package awssupportauthz

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSupportPermit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allActions interface{}
//   var allResourcesInRegion interface{}
//
//   cfnSupportPermitProps := &CfnSupportPermitProps{
//   	Name: jsii.String("name"),
//   	Permit: &PermitProperty{
//   		Actions: &ActionSetProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			AllActions: allActions,
//   		},
//   		Resources: &ResourceSetProperty{
//   			AllResourcesInRegion: allResourcesInRegion,
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				AllowAfter: jsii.String("allowAfter"),
//   				AllowBefore: jsii.String("allowBefore"),
//   			},
//   		},
//   	},
//   	SigningKeyInfo: &SigningKeyInfoProperty{
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	SupportCaseDisplayId: jsii.String("supportCaseDisplayId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html
//
type CfnSupportPermitProps struct {
	// The name of the support permit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html#cfn-supportauthz-supportpermit-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The grant definition: which actions on which resources, optionally constrained by time conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html#cfn-supportauthz-supportpermit-permit
	//
	Permit interface{} `field:"required" json:"permit" yaml:"permit"`
	// The signing key used by the permit.
	//
	// Exactly one key type must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html#cfn-supportauthz-supportpermit-signingkeyinfo
	//
	SigningKeyInfo interface{} `field:"required" json:"signingKeyInfo" yaml:"signingKeyInfo"`
	// An optional description of the support permit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html#cfn-supportauthz-supportpermit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The support case display identifier associated with the permit.
	//
	// When provided, the permit is linked to the specified AWS Support case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html#cfn-supportauthz-supportpermit-supportcasedisplayid
	//
	SupportCaseDisplayId *string `field:"optional" json:"supportCaseDisplayId" yaml:"supportCaseDisplayId"`
	// A list of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html#cfn-supportauthz-supportpermit-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

