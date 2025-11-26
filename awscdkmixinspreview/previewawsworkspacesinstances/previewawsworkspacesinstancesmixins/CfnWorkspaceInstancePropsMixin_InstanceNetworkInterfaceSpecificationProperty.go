package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceNetworkInterfaceSpecificationProperty := &InstanceNetworkInterfaceSpecificationProperty{
//   	Description: jsii.String("description"),
//   	DeviceIndex: jsii.Number(123),
//   	Groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification.html
//
type CfnWorkspaceInstancePropsMixin_InstanceNetworkInterfaceSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification.html#cfn-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification.html#cfn-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification-deviceindex
	//
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification.html#cfn-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification-groups
	//
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification.html#cfn-workspacesinstances-workspaceinstance-instancenetworkinterfacespecification-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

