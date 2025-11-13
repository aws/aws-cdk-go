package interfacesawsiotsitewise


// A reference to a Dashboard resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardReference := &DashboardReference{
//   	DashboardArn: jsii.String("dashboardArn"),
//   	DashboardId: jsii.String("dashboardId"),
//   }
//
type DashboardReference struct {
	// The ARN of the Dashboard resource.
	DashboardArn *string `field:"required" json:"dashboardArn" yaml:"dashboardArn"`
	// The DashboardId of the Dashboard resource.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
}

