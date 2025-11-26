package previewawsdaxmixins


// Properties for CfnSubnetGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubnetGroupMixinProps := &CfnSubnetGroupMixinProps{
//   	Description: jsii.String("description"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html
//
type CfnSubnetGroupMixinProps struct {
	// The description of the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html#cfn-dax-subnetgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html#cfn-dax-subnetgroup-subnetgroupname
	//
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// A list of VPC subnet IDs for the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html#cfn-dax-subnetgroup-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

