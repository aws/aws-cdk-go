package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   switchoverDetailProperty := &SwitchoverDetailProperty{
//   	SourceMember: jsii.String("sourceMember"),
//   	SwitchoverDetailStatus: jsii.String("switchoverDetailStatus"),
//   	TargetMember: jsii.String("targetMember"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-switchoverdetail.html
//
type CfnDeploymentPropsMixin_SwitchoverDetailProperty struct {
	// The source database member in the Blue/Green Deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-switchoverdetail.html#cfn-rds-deployment-switchoverdetail-sourcemember
	//
	SourceMember *string `field:"optional" json:"sourceMember" yaml:"sourceMember"`
	// The status of this source and target member pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-switchoverdetail.html#cfn-rds-deployment-switchoverdetail-switchoverdetailstatus
	//
	SwitchoverDetailStatus *string `field:"optional" json:"switchoverDetailStatus" yaml:"switchoverDetailStatus"`
	// The target database member in the Blue/Green Deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-switchoverdetail.html#cfn-rds-deployment-switchoverdetail-targetmember
	//
	TargetMember *string `field:"optional" json:"targetMember" yaml:"targetMember"`
}

