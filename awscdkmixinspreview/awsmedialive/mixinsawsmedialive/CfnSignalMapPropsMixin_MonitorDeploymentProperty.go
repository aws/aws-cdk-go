package mixinsawsmedialive


// Represents the latest monitor deployment of a signal map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitorDeploymentProperty := &MonitorDeploymentProperty{
//   	DetailsUri: jsii.String("detailsUri"),
//   	ErrorMessage: jsii.String("errorMessage"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-monitordeployment.html
//
type CfnSignalMapPropsMixin_MonitorDeploymentProperty struct {
	// URI associated with a signal map's monitor deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-monitordeployment.html#cfn-medialive-signalmap-monitordeployment-detailsuri
	//
	DetailsUri *string `field:"optional" json:"detailsUri" yaml:"detailsUri"`
	// Error message associated with a failed monitor deployment of a signal map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-monitordeployment.html#cfn-medialive-signalmap-monitordeployment-errormessage
	//
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// The signal map monitor deployment status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-monitordeployment.html#cfn-medialive-signalmap-monitordeployment-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

