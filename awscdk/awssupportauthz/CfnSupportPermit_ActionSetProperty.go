package awssupportauthz


// The set of actions a support permit grants.
//
// Exactly one of AllActions or Actions must be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allActions interface{}
//
//   actionSetProperty := &ActionSetProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	AllActions: allActions,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-actionset.html
//
type CfnSupportPermit_ActionSetProperty struct {
	// An explicit list of actions to grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-actionset.html#cfn-supportauthz-supportpermit-actionset-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Grants all actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-actionset.html#cfn-supportauthz-supportpermit-actionset-allactions
	//
	AllActions interface{} `field:"optional" json:"allActions" yaml:"allActions"`
}

