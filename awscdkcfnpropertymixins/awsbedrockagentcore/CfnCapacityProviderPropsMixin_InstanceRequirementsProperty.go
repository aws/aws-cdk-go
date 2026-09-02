package awsbedrockagentcore


// Requirements for EC2 instance types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   instanceRequirementsProperty := &InstanceRequirementsProperty{
//   	AllowedInstanceTypes: []*string{
//   		jsii.String("allowedInstanceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-instancerequirements.html
//
type CfnCapacityProviderPropsMixin_InstanceRequirementsProperty struct {
	// List of allowed instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-instancerequirements.html#cfn-bedrockagentcore-capacityprovider-instancerequirements-allowedinstancetypes
	//
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
}

