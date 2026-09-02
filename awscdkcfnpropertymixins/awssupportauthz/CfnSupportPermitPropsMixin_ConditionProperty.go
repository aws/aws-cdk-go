package awssupportauthz


// A time-bound condition controlling when the permit is active.
//
// Exactly one of AllowAfter or AllowBefore must be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   conditionProperty := &ConditionProperty{
//   	AllowAfter: jsii.String("allowAfter"),
//   	AllowBefore: jsii.String("allowBefore"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-condition.html
//
type CfnSupportPermitPropsMixin_ConditionProperty struct {
	// The permit is active only after this time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-condition.html#cfn-supportauthz-supportpermit-condition-allowafter
	//
	AllowAfter *string `field:"optional" json:"allowAfter" yaml:"allowAfter"`
	// The permit is active only before this time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-condition.html#cfn-supportauthz-supportpermit-condition-allowbefore
	//
	AllowBefore *string `field:"optional" json:"allowBefore" yaml:"allowBefore"`
}

