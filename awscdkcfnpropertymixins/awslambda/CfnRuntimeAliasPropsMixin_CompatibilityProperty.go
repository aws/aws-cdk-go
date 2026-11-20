package awslambda


// Compatibility matrix for the runtime version the alias points to.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-compatibility.html
//
type CfnRuntimeAliasPropsMixin_CompatibilityProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-compatibility.html#cfn-lambda-runtimealias-compatibility-architectures
	//
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-compatibility.html#cfn-lambda-runtimealias-compatibility-integrationmodels
	//
	IntegrationModels *[]*string `field:"optional" json:"integrationModels" yaml:"integrationModels"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-compatibility.html#cfn-lambda-runtimealias-compatibility-managedinstances
	//
	ManagedInstances *string `field:"optional" json:"managedInstances" yaml:"managedInstances"`
}

