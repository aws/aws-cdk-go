package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetRedshiftDataParametersProperty := &pipeTargetRedshiftDataParametersProperty{
//   	database: jsii.String("database"),
//   	sqls: []*string{
//   		jsii.String("sqls"),
//   	},
//
//   	// the properties below are optional
//   	dbUser: jsii.String("dbUser"),
//   	secretManagerArn: jsii.String("secretManagerArn"),
//   	statementName: jsii.String("statementName"),
//   	withEvent: jsii.Boolean(false),
//   }
//
type CfnPipe_PipeTargetRedshiftDataParametersProperty struct {
	// `CfnPipe.PipeTargetRedshiftDataParametersProperty.Database`.
	Database *string `field:"required" json:"database" yaml:"database"`
	// `CfnPipe.PipeTargetRedshiftDataParametersProperty.Sqls`.
	Sqls *[]*string `field:"required" json:"sqls" yaml:"sqls"`
	// `CfnPipe.PipeTargetRedshiftDataParametersProperty.DbUser`.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// `CfnPipe.PipeTargetRedshiftDataParametersProperty.SecretManagerArn`.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// `CfnPipe.PipeTargetRedshiftDataParametersProperty.StatementName`.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// `CfnPipe.PipeTargetRedshiftDataParametersProperty.WithEvent`.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

