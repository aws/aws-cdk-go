package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   graphDisplayConfigProperty := &GraphDisplayConfigProperty{
//   	VisualType: jsii.String("visualType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-graphdisplayconfig.html
//
type CfnDashboardPropsMixin_GraphDisplayConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-graphdisplayconfig.html#cfn-bcm-dashboard-graphdisplayconfig-visualtype
	//
	VisualType *string `field:"optional" json:"visualType" yaml:"visualType"`
}

