// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   tagParam := awscdk.NewCfnParameter(this, jsii.String("TagName"))
//
//   stringEquals := awscdk.NewCfnJson(this, jsii.String("ConditionJson"), &CfnJsonProps{
//   	Value: map[string]*bool{
//   		fmt.Sprintf("aws:PrincipalTag/%v", tagParam.valueAsString): jsii.Boolean(true),
//   	},
//   })
//
//   principal := iam.NewAccountRootPrincipal().WithConditions(map[string]interface{}{
//   	"StringEquals": stringEquals,
//   })
//
//   iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
//   	AssumedBy: principal,
//   })
//
type CfnJsonProps struct {
	// The value to resolve.
	//
	// Can be any JavaScript object, including tokens and
	// references in keys or values.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

