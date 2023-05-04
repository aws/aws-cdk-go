package awsquicksight


// Dashboard source entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardSourceEntityProperty := &DashboardSourceEntityProperty{
//   	SourceTemplate: &DashboardSourceTemplateProperty{
//   		Arn: jsii.String("arn"),
//   		DataSetReferences: []interface{}{
//   			&DataSetReferenceProperty{
//   				DataSetArn: jsii.String("dataSetArn"),
//   				DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_DashboardSourceEntityProperty struct {
	// Source template.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

