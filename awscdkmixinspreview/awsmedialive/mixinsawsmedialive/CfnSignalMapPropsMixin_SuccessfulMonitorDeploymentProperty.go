package mixinsawsmedialive


// Represents the latest successful monitor deployment of a signal map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   successfulMonitorDeploymentProperty := &SuccessfulMonitorDeploymentProperty{
//   	DetailsUri: jsii.String("detailsUri"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-successfulmonitordeployment.html
//
type CfnSignalMapPropsMixin_SuccessfulMonitorDeploymentProperty struct {
	// URI associated with a signal map's monitor deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-successfulmonitordeployment.html#cfn-medialive-signalmap-successfulmonitordeployment-detailsuri
	//
	DetailsUri *string `field:"optional" json:"detailsUri" yaml:"detailsUri"`
	// A signal map's monitor deployment status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-successfulmonitordeployment.html#cfn-medialive-signalmap-successfulmonitordeployment-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

