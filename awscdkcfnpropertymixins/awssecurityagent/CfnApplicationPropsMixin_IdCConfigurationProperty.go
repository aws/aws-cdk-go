package awssecurityagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   idCConfigurationProperty := &IdCConfigurationProperty{
//   	IdCApplicationArn: jsii.String("idCApplicationArn"),
//   	IdCInstanceArn: jsii.String("idCInstanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-application-idcconfiguration.html
//
type CfnApplicationPropsMixin_IdCConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-application-idcconfiguration.html#cfn-securityagent-application-idcconfiguration-idcapplicationarn
	//
	IdCApplicationArn *string `field:"optional" json:"idCApplicationArn" yaml:"idCApplicationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-application-idcconfiguration.html#cfn-securityagent-application-idcconfiguration-idcinstancearn
	//
	IdCInstanceArn *string `field:"optional" json:"idCInstanceArn" yaml:"idCInstanceArn"`
}

