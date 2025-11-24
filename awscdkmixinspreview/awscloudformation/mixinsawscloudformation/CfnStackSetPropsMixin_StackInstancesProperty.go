package mixinsawscloudformation


// Stack instances in some specific accounts and Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stackInstancesProperty := &StackInstancesProperty{
//   	DeploymentTargets: &DeploymentTargetsProperty{
//   		AccountFilterType: jsii.String("accountFilterType"),
//   		Accounts: []*string{
//   			jsii.String("accounts"),
//   		},
//   		AccountsUrl: jsii.String("accountsUrl"),
//   		OrganizationalUnitIds: []*string{
//   			jsii.String("organizationalUnitIds"),
//   		},
//   	},
//   	ParameterOverrides: []interface{}{
//   		&ParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html
//
type CfnStackSetPropsMixin_StackInstancesProperty struct {
	// The AWS Organizations accounts or AWS accounts to deploy stacks to in the specified Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-deploymenttargets
	//
	DeploymentTargets interface{} `field:"optional" json:"deploymentTargets" yaml:"deploymentTargets"`
	// A list of StackSet parameters whose values you want to override in the selected stack instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-parameteroverrides
	//
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The names of one or more Regions where you want to create stack instances using the specified AWS accounts .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

