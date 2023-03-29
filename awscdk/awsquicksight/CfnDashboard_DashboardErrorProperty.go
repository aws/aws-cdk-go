package awsquicksight


// Dashboard error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDashboard_DashboardErrorProperty struct {
	// Message.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Type.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// `CfnDashboard.DashboardErrorProperty.ViolatedEntities`.
	ViolatedEntities interface{} `field:"optional" json:"violatedEntities" yaml:"violatedEntities"`
}

