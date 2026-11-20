package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trustedExecutionConfigurationProperty := &TrustedExecutionConfigurationProperty{
//   	AttestationPolicy: jsii.String("attestationPolicy"),
//   	IsolationMode: jsii.String("isolationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-trustedexecutionconfiguration.html
//
type CfnTaskDefinitionPropsMixin_TrustedExecutionConfigurationProperty struct {
	// A base64-encoded attestation policy that verifies the workload before the task runs.
	//
	// Only valid when IsolationMode is CONFIDENTIAL_SANDBOX.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-trustedexecutionconfiguration.html#cfn-ecs-taskdefinition-trustedexecutionconfiguration-attestationpolicy
	//
	AttestationPolicy *string `field:"optional" json:"attestationPolicy" yaml:"attestationPolicy"`
	// The isolation mode for the task.
	//
	// SHARED uses standard shared placement. ISOLATED_SANDBOX runs the task in its own isolated sandbox on a shared instance with hypervisor-level isolation. CONFIDENTIAL_SANDBOX runs the task in an isolated sandbox and supports an attestation policy to verify the workload before it runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-trustedexecutionconfiguration.html#cfn-ecs-taskdefinition-trustedexecutionconfiguration-isolationmode
	//
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
}

