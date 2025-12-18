package previewawslambdamixins


// Specifications that define the characteristics and constraints for compute instances used by the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceRequirementsProperty := &InstanceRequirementsProperty{
//   	AllowedInstanceTypes: []*string{
//   		jsii.String("allowedInstanceTypes"),
//   	},
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	ExcludedInstanceTypes: []*string{
//   		jsii.String("excludedInstanceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-instancerequirements.html
//
type CfnCapacityProviderPropsMixin_InstanceRequirementsProperty struct {
	// A list of EC2 instance types that the capacity provider is allowed to use.
	//
	// If not specified, all compatible instance types are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-instancerequirements.html#cfn-lambda-capacityprovider-instancerequirements-allowedinstancetypes
	//
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// A list of supported CPU architectures for compute instances.
	//
	// Valid values include `x86_64` and `arm64` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-instancerequirements.html#cfn-lambda-capacityprovider-instancerequirements-architectures
	//
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// A list of EC2 instance types that the capacity provider should not use, even if they meet other requirements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-instancerequirements.html#cfn-lambda-capacityprovider-instancerequirements-excludedinstancetypes
	//
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
}

