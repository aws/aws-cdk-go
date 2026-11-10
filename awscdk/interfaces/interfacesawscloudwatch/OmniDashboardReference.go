package interfacesawscloudwatch


// A reference to a OmniDashboard resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   omniDashboardReference := &OmniDashboardReference{
//   	DashboardId: jsii.String("dashboardId"),
//   	OmniDashboardArn: jsii.String("omniDashboardArn"),
//   	SpaceId: jsii.String("spaceId"),
//   }
//
type OmniDashboardReference struct {
	// The DashboardId of the OmniDashboard resource.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
	// The ARN of the OmniDashboard resource.
	OmniDashboardArn *string `field:"required" json:"omniDashboardArn" yaml:"omniDashboardArn"`
	// The SpaceId of the OmniDashboard resource.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

