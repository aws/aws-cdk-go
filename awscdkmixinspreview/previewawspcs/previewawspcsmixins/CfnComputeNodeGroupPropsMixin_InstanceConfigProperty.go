package previewawspcsmixins


// An EC2 instance configuration AWS PCS uses to launch compute nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceConfigProperty := &InstanceConfigProperty{
//   	InstanceType: jsii.String("instanceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-instanceconfig.html
//
type CfnComputeNodeGroupPropsMixin_InstanceConfigProperty struct {
	// The EC2 instance type that AWS PCS can provision in the compute node group.
	//
	// Example: `t2.xlarge`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-instanceconfig.html#cfn-pcs-computenodegroup-instanceconfig-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
}

