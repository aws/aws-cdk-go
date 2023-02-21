package awsquicksight


// Dashboard source entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardSourceEntityProperty := &dashboardSourceEntityProperty{
//   	sourceTemplate: &dashboardSourceTemplateProperty{
//   		arn: jsii.String("arn"),
//   		dataSetReferences: []interface{}{
//   			&dataSetReferenceProperty{
//   				dataSetArn: jsii.String("dataSetArn"),
//   				dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_DashboardSourceEntityProperty struct {
	// Source template.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

