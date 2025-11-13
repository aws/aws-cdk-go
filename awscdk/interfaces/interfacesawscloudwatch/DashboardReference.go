package interfacesawscloudwatch


// A reference to a Dashboard resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardReference := &DashboardReference{
//   	DashboardName: jsii.String("dashboardName"),
//   }
//
type DashboardReference struct {
	// The DashboardName of the Dashboard resource.
	DashboardName *string `field:"required" json:"dashboardName" yaml:"dashboardName"`
}

