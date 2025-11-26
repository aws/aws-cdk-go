package previewawsdeadlinemixins


// The configuration options for a service managed fleet's VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConfigurationProperty := &VpcConfigurationProperty{
//   	ResourceConfigurationArns: []*string{
//   		jsii.String("resourceConfigurationArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-vpcconfiguration.html
//
type CfnFleetPropsMixin_VpcConfigurationProperty struct {
	// The ARNs of the VPC Lattice resource configurations attached to the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-vpcconfiguration.html#cfn-deadline-fleet-vpcconfiguration-resourceconfigurationarns
	//
	ResourceConfigurationArns *[]*string `field:"optional" json:"resourceConfigurationArns" yaml:"resourceConfigurationArns"`
}

