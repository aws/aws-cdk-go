package awsopensearchservice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceSoftwareOptionsProperty := &serviceSoftwareOptionsProperty{
//   	automatedUpdateDate: jsii.String("automatedUpdateDate"),
//   	cancellable: jsii.Boolean(false),
//   	currentVersion: jsii.String("currentVersion"),
//   	description: jsii.String("description"),
//   	newVersion: jsii.String("newVersion"),
//   	optionalDeployment: jsii.Boolean(false),
//   	updateAvailable: jsii.Boolean(false),
//   	updateStatus: jsii.String("updateStatus"),
//   }
//
type CfnDomain_ServiceSoftwareOptionsProperty struct {
	// `CfnDomain.ServiceSoftwareOptionsProperty.AutomatedUpdateDate`.
	AutomatedUpdateDate *string `field:"optional" json:"automatedUpdateDate" yaml:"automatedUpdateDate"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.Cancellable`.
	Cancellable interface{} `field:"optional" json:"cancellable" yaml:"cancellable"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.CurrentVersion`.
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.NewVersion`.
	NewVersion *string `field:"optional" json:"newVersion" yaml:"newVersion"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.OptionalDeployment`.
	OptionalDeployment interface{} `field:"optional" json:"optionalDeployment" yaml:"optionalDeployment"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.UpdateAvailable`.
	UpdateAvailable interface{} `field:"optional" json:"updateAvailable" yaml:"updateAvailable"`
	// `CfnDomain.ServiceSoftwareOptionsProperty.UpdateStatus`.
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
}

