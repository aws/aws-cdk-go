package interfacesawscloudtrail


// A reference to a Dashboard resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardReference := &DashboardReference{
//   	DashboardArn: jsii.String("dashboardArn"),
//   }
//
type DashboardReference struct {
	// The DashboardArn of the Dashboard resource.
	DashboardArn *string `field:"required" json:"dashboardArn" yaml:"dashboardArn"`
}

