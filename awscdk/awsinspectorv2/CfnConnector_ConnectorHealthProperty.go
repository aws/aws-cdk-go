package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorHealthProperty := &ConnectorHealthProperty{
//   	ConnectorStatus: jsii.String("connectorStatus"),
//   	LastCheckedAt: jsii.String("lastCheckedAt"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-connectorhealth.html
//
type CfnConnector_ConnectorHealthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-connectorhealth.html#cfn-inspectorv2-connector-connectorhealth-connectorstatus
	//
	ConnectorStatus *string `field:"optional" json:"connectorStatus" yaml:"connectorStatus"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-connectorhealth.html#cfn-inspectorv2-connector-connectorhealth-lastcheckedat
	//
	LastCheckedAt *string `field:"optional" json:"lastCheckedAt" yaml:"lastCheckedAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-connectorhealth.html#cfn-inspectorv2-connector-connectorhealth-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

