// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

func ExampleAutoMLProblemTypeConfig_outputUsage() {
	var union types.AutoMLProblemTypeConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AutoMLProblemTypeConfigMemberImageClassificationJobConfig:
		_ = v.Value // Value is types.ImageClassificationJobConfig

	case *types.AutoMLProblemTypeConfigMemberTabularJobConfig:
		_ = v.Value // Value is types.TabularJobConfig

	case *types.AutoMLProblemTypeConfigMemberTextClassificationJobConfig:
		_ = v.Value // Value is types.TextClassificationJobConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TextClassificationJobConfig
var _ *types.ImageClassificationJobConfig
var _ *types.TabularJobConfig

func ExampleAutoMLProblemTypeResolvedAttributes_outputUsage() {
	var union types.AutoMLProblemTypeResolvedAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AutoMLProblemTypeResolvedAttributesMemberTabularResolvedAttributes:
		_ = v.Value // Value is types.TabularResolvedAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TabularResolvedAttributes

func ExampleTrialComponentParameterValue_outputUsage() {
	var union types.TrialComponentParameterValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TrialComponentParameterValueMemberNumberValue:
		_ = v.Value // Value is float64

	case *types.TrialComponentParameterValueMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *float64
