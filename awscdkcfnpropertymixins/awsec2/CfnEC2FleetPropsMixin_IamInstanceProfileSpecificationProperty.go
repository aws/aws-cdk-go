package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   iamInstanceProfileSpecificationProperty := &IamInstanceProfileSpecificationProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-iaminstanceprofilespecification.html
//
type CfnEC2FleetPropsMixin_IamInstanceProfileSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-iaminstanceprofilespecification.html#cfn-ec2-ec2fleet-iaminstanceprofilespecification-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-iaminstanceprofilespecification.html#cfn-ec2-ec2fleet-iaminstanceprofilespecification-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

