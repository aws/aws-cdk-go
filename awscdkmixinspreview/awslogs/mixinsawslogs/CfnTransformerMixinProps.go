package mixinsawslogs


// Properties for CfnTransformerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerMixinProps := &CfnTransformerMixinProps{
//   	LogGroupIdentifier: jsii.String("logGroupIdentifier"),
//   	TransformerConfig: []interface{}{
//   		&ProcessorProperty{
//   			AddKeys: &AddKeysProperty{
//   				Entries: []interface{}{
//   					&AddKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						OverwriteIfExists: jsii.Boolean(false),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			CopyValue: &CopyValueProperty{
//   				Entries: []interface{}{
//   					&CopyValueEntryProperty{
//   						OverwriteIfExists: jsii.Boolean(false),
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//   					},
//   				},
//   			},
//   			Csv: &CsvProperty{
//   				Columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				Delimiter: jsii.String("delimiter"),
//   				QuoteCharacter: jsii.String("quoteCharacter"),
//   				Source: jsii.String("source"),
//   			},
//   			DateTimeConverter: &DateTimeConverterProperty{
//   				Locale: jsii.String("locale"),
//   				MatchPatterns: []*string{
//   					jsii.String("matchPatterns"),
//   				},
//   				Source: jsii.String("source"),
//   				SourceTimezone: jsii.String("sourceTimezone"),
//   				Target: jsii.String("target"),
//   				TargetFormat: jsii.String("targetFormat"),
//   				TargetTimezone: jsii.String("targetTimezone"),
//   			},
//   			DeleteKeys: &DeleteKeysProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			Grok: &GrokProperty{
//   				Match: jsii.String("match"),
//   				Source: jsii.String("source"),
//   			},
//   			ListToMap: &ListToMapProperty{
//   				Flatten: jsii.Boolean(false),
//   				FlattenedElement: jsii.String("flattenedElement"),
//   				Key: jsii.String("key"),
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//   				ValueKey: jsii.String("valueKey"),
//   			},
//   			LowerCaseString: &LowerCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			MoveKeys: &MoveKeysProperty{
//   				Entries: []interface{}{
//   					&MoveKeyEntryProperty{
//   						OverwriteIfExists: jsii.Boolean(false),
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//   					},
//   				},
//   			},
//   			ParseCloudfront: &ParseCloudfrontProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseJson: &ParseJSONProperty{
//   				Destination: jsii.String("destination"),
//   				Source: jsii.String("source"),
//   			},
//   			ParseKeyValue: &ParseKeyValueProperty{
//   				Destination: jsii.String("destination"),
//   				FieldDelimiter: jsii.String("fieldDelimiter"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   				KeyValueDelimiter: jsii.String("keyValueDelimiter"),
//   				NonMatchValue: jsii.String("nonMatchValue"),
//   				OverwriteIfExists: jsii.Boolean(false),
//   				Source: jsii.String("source"),
//   			},
//   			ParsePostgres: &ParsePostgresProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseRoute53: &ParseRoute53Property{
//   				Source: jsii.String("source"),
//   			},
//   			ParseToOcsf: &ParseToOCSFProperty{
//   				EventSource: jsii.String("eventSource"),
//   				OcsfVersion: jsii.String("ocsfVersion"),
//   				Source: jsii.String("source"),
//   			},
//   			ParseVpc: &ParseVPCProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseWaf: &ParseWAFProperty{
//   				Source: jsii.String("source"),
//   			},
//   			RenameKeys: &RenameKeysProperty{
//   				Entries: []interface{}{
//   					&RenameKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						OverwriteIfExists: jsii.Boolean(false),
//   						RenameTo: jsii.String("renameTo"),
//   					},
//   				},
//   			},
//   			SplitString: &SplitStringProperty{
//   				Entries: []interface{}{
//   					&SplitStringEntryProperty{
//   						Delimiter: jsii.String("delimiter"),
//   						Source: jsii.String("source"),
//   					},
//   				},
//   			},
//   			SubstituteString: &SubstituteStringProperty{
//   				Entries: []interface{}{
//   					&SubstituteStringEntryProperty{
//   						From: jsii.String("from"),
//   						Source: jsii.String("source"),
//   						To: jsii.String("to"),
//   					},
//   				},
//   			},
//   			TrimString: &TrimStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			TypeConverter: &TypeConverterProperty{
//   				Entries: []interface{}{
//   					&TypeConverterEntryProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			UpperCaseString: &UpperCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html
//
type CfnTransformerMixinProps struct {
	// Specify either the name or ARN of the log group to create the transformer for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html#cfn-logs-transformer-loggroupidentifier
	//
	LogGroupIdentifier *string `field:"optional" json:"logGroupIdentifier" yaml:"logGroupIdentifier"`
	// This structure is an array that contains the configuration of this log transformer.
	//
	// A log transformer is an array of processors, where each processor applies one type of transformation to the log events that are ingested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html#cfn-logs-transformer-transformerconfig
	//
	TransformerConfig interface{} `field:"optional" json:"transformerConfig" yaml:"transformerConfig"`
}

