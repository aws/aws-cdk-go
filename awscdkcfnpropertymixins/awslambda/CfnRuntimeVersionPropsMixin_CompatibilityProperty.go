package awslambda


// Compatibility matrix for the runtime version: supported integration models, architectures, and feature policies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   compatibilityProperty := &CompatibilityProperty{
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	IntegrationModels: []*string{
//   		jsii.String("integrationModels"),
//   	},
//   	ManagedInstances: jsii.String("managedInstances"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-compatibility.html
//
type CfnRuntimeVersionPropsMixin_CompatibilityProperty struct {
	// Supported CPU architectures for the runtime version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-compatibility.html#cfn-lambda-runtimeversion-compatibility-architectures
	//
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Supported integration models for the runtime version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-compatibility.html#cfn-lambda-runtimeversion-compatibility-integrationmodels
	//
	IntegrationModels *[]*string `field:"optional" json:"integrationModels" yaml:"integrationModels"`
	// Whether the runtime version supports managed instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-compatibility.html#cfn-lambda-runtimeversion-compatibility-managedinstances
	//
	ManagedInstances *string `field:"optional" json:"managedInstances" yaml:"managedInstances"`
}

