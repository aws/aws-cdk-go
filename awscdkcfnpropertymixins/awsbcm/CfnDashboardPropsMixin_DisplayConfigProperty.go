package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var table interface{}
//
//   displayConfigProperty := &DisplayConfigProperty{
//   	Graph: map[string]interface{}{
//   		"graphKey": &GraphDisplayConfigProperty{
//   			"visualType": jsii.String("visualType"),
//   		},
//   	},
//   	Table: table,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-displayconfig.html
//
type CfnDashboardPropsMixin_DisplayConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-displayconfig.html#cfn-bcm-dashboard-displayconfig-graph
	//
	Graph interface{} `field:"optional" json:"graph" yaml:"graph"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-displayconfig.html#cfn-bcm-dashboard-displayconfig-table
	//
	Table interface{} `field:"optional" json:"table" yaml:"table"`
}

