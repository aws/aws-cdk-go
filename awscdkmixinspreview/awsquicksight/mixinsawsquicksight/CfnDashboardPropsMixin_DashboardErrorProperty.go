package mixinsawsquicksight


// Dashboard error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashboardErrorProperty := &DashboardErrorProperty{
//   	Message: jsii.String("message"),
//   	Type: jsii.String("type"),
//   	ViolatedEntities: []interface{}{
//   		&EntityProperty{
//   			Path: jsii.String("path"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboarderror.html
//
type CfnDashboardPropsMixin_DashboardErrorProperty struct {
	// Message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboarderror.html#cfn-quicksight-dashboard-dashboarderror-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboarderror.html#cfn-quicksight-dashboard-dashboarderror-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Lists the violated entities that caused the dashboard error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboarderror.html#cfn-quicksight-dashboard-dashboarderror-violatedentities
	//
	ViolatedEntities interface{} `field:"optional" json:"violatedEntities" yaml:"violatedEntities"`
}

