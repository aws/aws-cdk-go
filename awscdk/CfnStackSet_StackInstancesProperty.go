package awscdk


// Stack instances in some specific accounts and Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
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
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//
//   	// the properties below are optional
//   	ParameterOverrides: []interface{}{
//   		&ParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html
//
type CfnStackSet_StackInstancesProperty struct {
	// The AWS `OrganizationalUnitIds` or `Accounts` for which to create stack instances in the specified Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-deploymenttargets
	//
	DeploymentTargets interface{} `field:"required" json:"deploymentTargets" yaml:"deploymentTargets"`
	// The names of one or more Regions where you want to create stack instances using the specified AWS accounts .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-regions
	//
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// A list of stack set parameters whose values you want to override in the selected stack instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-parameteroverrides
	//
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
}

