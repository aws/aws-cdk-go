package awssupportauthz


// The grant definition: which actions on which resources, optionally constrained by time conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allActions interface{}
//   var allResourcesInRegion interface{}
//
//   permitProperty := &PermitProperty{
//   	Actions: &ActionSetProperty{
//   		Actions: []*string{
//   			jsii.String("actions"),
//   		},
//   		AllActions: allActions,
//   	},
//   	Resources: &ResourceSetProperty{
//   		AllResourcesInRegion: allResourcesInRegion,
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			AllowAfter: jsii.String("allowAfter"),
//   			AllowBefore: jsii.String("allowBefore"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-permit.html
//
type CfnSupportPermit_PermitProperty struct {
	// The set of actions a support permit grants.
	//
	// Exactly one of AllActions or Actions must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-permit.html#cfn-supportauthz-supportpermit-permit-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The set of resources a support permit applies to.
	//
	// Exactly one of AllResourcesInRegion or Resources must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-permit.html#cfn-supportauthz-supportpermit-permit-resources
	//
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// Optional time-bound conditions (at most two).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-permit.html#cfn-supportauthz-supportpermit-permit-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

