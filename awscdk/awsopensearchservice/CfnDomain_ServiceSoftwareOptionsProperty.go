package awsopensearchservice


// The current status of the service software for an Amazon OpenSearch Service domain.
//
// For more information, see [Service software updates in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceSoftwareOptionsProperty := &ServiceSoftwareOptionsProperty{
//   	AutomatedUpdateDate: jsii.String("automatedUpdateDate"),
//   	Cancellable: jsii.Boolean(false),
//   	CurrentVersion: jsii.String("currentVersion"),
//   	Description: jsii.String("description"),
//   	NewVersion: jsii.String("newVersion"),
//   	OptionalDeployment: jsii.Boolean(false),
//   	UpdateAvailable: jsii.Boolean(false),
//   	UpdateStatus: jsii.String("updateStatus"),
//   }
//
type CfnDomain_ServiceSoftwareOptionsProperty struct {
	// The timestamp, in Epoch time, until which you can manually request a service software update.
	//
	// After this date, we automatically update your service software.
	AutomatedUpdateDate *string `field:"optional" json:"automatedUpdateDate" yaml:"automatedUpdateDate"`
	// True if you're able to cancel your service software version update.
	//
	// False if you can't cancel your service software update.
	Cancellable interface{} `field:"optional" json:"cancellable" yaml:"cancellable"`
	// The current service software version present on the domain.
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// A description of the service software update status.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The new service software version, if one is available.
	NewVersion *string `field:"optional" json:"newVersion" yaml:"newVersion"`
	// True if a service software is never automatically updated.
	//
	// False if a service software is automatically updated after the automated update date.
	OptionalDeployment interface{} `field:"optional" json:"optionalDeployment" yaml:"optionalDeployment"`
	// True if you're able to update your service software version.
	//
	// False if you can't update your service software version.
	UpdateAvailable interface{} `field:"optional" json:"updateAvailable" yaml:"updateAvailable"`
	// The status of your service software update.
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
}

