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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html
//
type CfnDomain_ServiceSoftwareOptionsProperty struct {
	// The timestamp, in Epoch time, until which you can manually request a service software update.
	//
	// After this date, we automatically update your service software.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-automatedupdatedate
	//
	AutomatedUpdateDate *string `field:"optional" json:"automatedUpdateDate" yaml:"automatedUpdateDate"`
	// True if you're able to cancel your service software version update.
	//
	// False if you can't cancel your service software update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-cancellable
	//
	Cancellable interface{} `field:"optional" json:"cancellable" yaml:"cancellable"`
	// The current service software version present on the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-currentversion
	//
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// A description of the service software update status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The new service software version, if one is available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-newversion
	//
	NewVersion *string `field:"optional" json:"newVersion" yaml:"newVersion"`
	// True if a service software is never automatically updated.
	//
	// False if a service software is automatically updated after the automated update date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-optionaldeployment
	//
	OptionalDeployment interface{} `field:"optional" json:"optionalDeployment" yaml:"optionalDeployment"`
	// True if you're able to update your service software version.
	//
	// False if you can't update your service software version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-updateavailable
	//
	UpdateAvailable interface{} `field:"optional" json:"updateAvailable" yaml:"updateAvailable"`
	// The status of your service software update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-updatestatus
	//
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
}

