package mixinsawsec2


// Specifies an instance family to use as the baseline reference for CPU performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceProperty := &ReferenceProperty{
//   	InstanceFamily: jsii.String("instanceFamily"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-reference.html
//
type CfnLaunchTemplatePropsMixin_ReferenceProperty struct {
	// The instance family to use as a baseline reference.
	//
	// > Ensure that you specify the correct value for the instance family. The instance family is everything before the period ( `.` ) in the instance type name. For example, in the instance type `c6i.large` , the instance family is `c6i` , not `c6` . For more information, see [Amazon EC2 instance type naming conventions](https://docs.aws.amazon.com/ec2/latest/instancetypes/instance-type-names.html) in *Amazon EC2 Instance Types* .
	//
	// The following instance families are *not supported* for performance protection:
	//
	// - `c1`
	// - `g3` | `g3s`
	// - `hpc7g`
	// - `m1` | `m2`
	// - `mac1` | `mac2` | `mac2-m1ultra` | `mac2-m2` | `mac2-m2pro`
	// - `p3dn` | `p4d` | `p5`
	// - `t1`
	// - `u-12tb1` | `u-18tb1` | `u-24tb1` | `u-3tb1` | `u-6tb1` | `u-9tb1` | `u7i-12tb` | `u7in-16tb` | `u7in-24tb` | `u7in-32tb`
	//
	// If you enable performance protection by specifying a supported instance family, the returned instance types will exclude the above unsupported instance families.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-reference.html#cfn-ec2-launchtemplate-reference-instancefamily
	//
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
}

