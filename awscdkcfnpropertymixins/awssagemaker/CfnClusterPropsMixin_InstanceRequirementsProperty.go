package awssagemaker


// The instance requirements for the instance group.
//
// Specifies a list of instance types that can be used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   instanceRequirementsProperty := &InstanceRequirementsProperty{
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-instancerequirements.html
//
type CfnClusterPropsMixin_InstanceRequirementsProperty struct {
	// A list of instance types that can be used for this instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-instancerequirements.html#cfn-sagemaker-cluster-instancerequirements-instancetypes
	//
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
}

