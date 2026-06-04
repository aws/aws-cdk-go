package awsapplicationsignals


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   compositeSliConfigProperty := &CompositeSliConfigProperty{
//   	CompositeSliComponents: []interface{}{
//   		&CompositeSliComponentProperty{
//   			OperationName: jsii.String("operationName"),
//   		},
//   	},
//   	SelectionConfig: &SelectionConfigProperty{
//   		Pattern: jsii.String("pattern"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositesliconfig.html
//
type CfnServiceLevelObjectivePropsMixin_CompositeSliConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositesliconfig.html#cfn-applicationsignals-servicelevelobjective-compositesliconfig-compositeslicomponents
	//
	CompositeSliComponents interface{} `field:"optional" json:"compositeSliComponents" yaml:"compositeSliComponents"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositesliconfig.html#cfn-applicationsignals-servicelevelobjective-compositesliconfig-selectionconfig
	//
	SelectionConfig interface{} `field:"optional" json:"selectionConfig" yaml:"selectionConfig"`
}

