package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A collection of the fields in a PolicyStatement that can be used to identify a principal.
//
// This consists of the JSON used in the "Principal" field, and optionally a
// set of "Condition"s that need to be applied to the policy.
//
// Generally, a principal looks like:
//
//     { '<TYPE>': ['ID', 'ID', ...] }
//
// And this is also the type of the field `principalJson`.  However, there is a
// special type of principal that is just the string '*', which is treated
// differently by some services. To represent that principal, `principalJson`
// should contain `{ 'LiteralString': ['*'] }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//
//   principalPolicyFragment := awscdk.Aws_iam.NewPrincipalPolicyFragment(map[string][]*string{
//   	"principalJsonKey": []*string{
//   		jsii.String("principalJson"),
//   	},
//   }, map[string]interface{}{
//   	"conditionsKey": conditions,
//   })
//
type PrincipalPolicyFragment interface {
	// The conditions under which the policy is in effect.
	//
	// See [the IAM documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html).
	// conditions that need to be applied to this policy.
	Conditions() *map[string]interface{}
	// JSON of the "Principal" section in a policy statement.
	PrincipalJson() *map[string]*[]*string
}

// The jsii proxy struct for PrincipalPolicyFragment
type jsiiProxy_PrincipalPolicyFragment struct {
	_ byte // padding
}

func (j *jsiiProxy_PrincipalPolicyFragment) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalPolicyFragment) PrincipalJson() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"principalJson",
		&returns,
	)
	return returns
}


func NewPrincipalPolicyFragment(principalJson *map[string]*[]*string, conditions *map[string]interface{}) PrincipalPolicyFragment {
	_init_.Initialize()

	if err := validateNewPrincipalPolicyFragmentParameters(principalJson); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrincipalPolicyFragment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PrincipalPolicyFragment",
		[]interface{}{principalJson, conditions},
		&j,
	)

	return &j
}

func NewPrincipalPolicyFragment_Override(p PrincipalPolicyFragment, principalJson *map[string]*[]*string, conditions *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PrincipalPolicyFragment",
		[]interface{}{principalJson, conditions},
		p,
	)
}

