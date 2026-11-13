package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   securityContextProperty := &SecurityContextProperty{
//   	AllowPrivilegeEscalation: jsii.Boolean(false),
//   	Capabilities: &EksContainerLinuxCapabilitiesProperty{
//   		Add: []*string{
//   			jsii.String("add"),
//   		},
//   		Drop: []*string{
//   			jsii.String("drop"),
//   		},
//   	},
//   	Privileged: jsii.Boolean(false),
//   	ReadOnlyRootFilesystem: jsii.Boolean(false),
//   	RunAsGroup: jsii.Number(123),
//   	RunAsNonRoot: jsii.Boolean(false),
//   	RunAsUser: jsii.Number(123),
//   	SeccompProfile: &EksSeccompProfileProperty{
//   		LocalhostProfile: jsii.String("localhostProfile"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html
//
type CfnJobDefinitionPropsMixin_SecurityContextProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-allowprivilegeescalation
	//
	AllowPrivilegeEscalation interface{} `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-capabilities
	//
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-privileged
	//
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-readonlyrootfilesystem
	//
	ReadOnlyRootFilesystem interface{} `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-runasgroup
	//
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-runasnonroot
	//
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-runasuser
	//
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-securitycontext.html#cfn-batch-jobdefinition-securitycontext-seccompprofile
	//
	SeccompProfile interface{} `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
}

