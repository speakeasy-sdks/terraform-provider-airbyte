// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceE2eTestCloudMockCatalogMultiSchemaType string

const (
	SourceE2eTestCloudMockCatalogMultiSchemaTypeMultiStream SourceE2eTestCloudMockCatalogMultiSchemaType = "MULTI_STREAM"
)

func (e SourceE2eTestCloudMockCatalogMultiSchemaType) ToPointer() *SourceE2eTestCloudMockCatalogMultiSchemaType {
	return &e
}

func (e *SourceE2eTestCloudMockCatalogMultiSchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MULTI_STREAM":
		*e = SourceE2eTestCloudMockCatalogMultiSchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceE2eTestCloudMockCatalogMultiSchemaType: %v", v)
	}
}

// SourceE2eTestCloudMockCatalogMultiSchema - A catalog with multiple data streams, each with a different schema.
type SourceE2eTestCloudMockCatalogMultiSchema struct {
	// A Json object specifying multiple data streams and their schemas. Each key in this object is one stream name. Each value is the schema for that stream. The schema should be compatible with <a href="https://json-schema.org/draft-07/json-schema-release-notes.html">draft-07</a>. See <a href="https://cswr.github.io/JsonSchema/spec/introduction/">this doc</a> for examples.
	StreamSchemas string                                       `json:"stream_schemas"`
	Type          SourceE2eTestCloudMockCatalogMultiSchemaType `json:"type"`
}

type SourceE2eTestCloudMockCatalogSingleSchemaType string

const (
	SourceE2eTestCloudMockCatalogSingleSchemaTypeSingleStream SourceE2eTestCloudMockCatalogSingleSchemaType = "SINGLE_STREAM"
)

func (e SourceE2eTestCloudMockCatalogSingleSchemaType) ToPointer() *SourceE2eTestCloudMockCatalogSingleSchemaType {
	return &e
}

func (e *SourceE2eTestCloudMockCatalogSingleSchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SINGLE_STREAM":
		*e = SourceE2eTestCloudMockCatalogSingleSchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceE2eTestCloudMockCatalogSingleSchemaType: %v", v)
	}
}

// SourceE2eTestCloudMockCatalogSingleSchema - A catalog with one or multiple streams that share the same schema.
type SourceE2eTestCloudMockCatalogSingleSchema struct {
	// Duplicate the stream for easy load testing. Each stream name will have a number suffix. For example, if the stream name is "ds", the duplicated streams will be "ds_0", "ds_1", etc.
	StreamDuplication *int64 `json:"stream_duplication,omitempty"`
	// Name of the data stream.
	StreamName string `json:"stream_name"`
	// A Json schema for the stream. The schema should be compatible with <a href="https://json-schema.org/draft-07/json-schema-release-notes.html">draft-07</a>. See <a href="https://cswr.github.io/JsonSchema/spec/introduction/">this doc</a> for examples.
	StreamSchema string                                        `json:"stream_schema"`
	Type         SourceE2eTestCloudMockCatalogSingleSchemaType `json:"type"`
}

type SourceE2eTestCloudMockCatalogType string

const (
	SourceE2eTestCloudMockCatalogTypeSourceE2eTestCloudMockCatalogSingleSchema SourceE2eTestCloudMockCatalogType = "source-e2e-test-cloud_Mock Catalog_Single Schema"
	SourceE2eTestCloudMockCatalogTypeSourceE2eTestCloudMockCatalogMultiSchema  SourceE2eTestCloudMockCatalogType = "source-e2e-test-cloud_Mock Catalog_Multi Schema"
)

type SourceE2eTestCloudMockCatalog struct {
	SourceE2eTestCloudMockCatalogSingleSchema *SourceE2eTestCloudMockCatalogSingleSchema
	SourceE2eTestCloudMockCatalogMultiSchema  *SourceE2eTestCloudMockCatalogMultiSchema

	Type SourceE2eTestCloudMockCatalogType
}

func CreateSourceE2eTestCloudMockCatalogSourceE2eTestCloudMockCatalogSingleSchema(sourceE2eTestCloudMockCatalogSingleSchema SourceE2eTestCloudMockCatalogSingleSchema) SourceE2eTestCloudMockCatalog {
	typ := SourceE2eTestCloudMockCatalogTypeSourceE2eTestCloudMockCatalogSingleSchema

	return SourceE2eTestCloudMockCatalog{
		SourceE2eTestCloudMockCatalogSingleSchema: &sourceE2eTestCloudMockCatalogSingleSchema,
		Type: typ,
	}
}

func CreateSourceE2eTestCloudMockCatalogSourceE2eTestCloudMockCatalogMultiSchema(sourceE2eTestCloudMockCatalogMultiSchema SourceE2eTestCloudMockCatalogMultiSchema) SourceE2eTestCloudMockCatalog {
	typ := SourceE2eTestCloudMockCatalogTypeSourceE2eTestCloudMockCatalogMultiSchema

	return SourceE2eTestCloudMockCatalog{
		SourceE2eTestCloudMockCatalogMultiSchema: &sourceE2eTestCloudMockCatalogMultiSchema,
		Type:                                     typ,
	}
}

func (u *SourceE2eTestCloudMockCatalog) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceE2eTestCloudMockCatalogSingleSchema := new(SourceE2eTestCloudMockCatalogSingleSchema)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceE2eTestCloudMockCatalogSingleSchema); err == nil {
		u.SourceE2eTestCloudMockCatalogSingleSchema = sourceE2eTestCloudMockCatalogSingleSchema
		u.Type = SourceE2eTestCloudMockCatalogTypeSourceE2eTestCloudMockCatalogSingleSchema
		return nil
	}

	sourceE2eTestCloudMockCatalogMultiSchema := new(SourceE2eTestCloudMockCatalogMultiSchema)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceE2eTestCloudMockCatalogMultiSchema); err == nil {
		u.SourceE2eTestCloudMockCatalogMultiSchema = sourceE2eTestCloudMockCatalogMultiSchema
		u.Type = SourceE2eTestCloudMockCatalogTypeSourceE2eTestCloudMockCatalogMultiSchema
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceE2eTestCloudMockCatalog) MarshalJSON() ([]byte, error) {
	if u.SourceE2eTestCloudMockCatalogSingleSchema != nil {
		return json.Marshal(u.SourceE2eTestCloudMockCatalogSingleSchema)
	}

	if u.SourceE2eTestCloudMockCatalogMultiSchema != nil {
		return json.Marshal(u.SourceE2eTestCloudMockCatalogMultiSchema)
	}

	return nil, nil
}

type SourceE2eTestCloudE2eTestCloud string

const (
	SourceE2eTestCloudE2eTestCloudE2eTestCloud SourceE2eTestCloudE2eTestCloud = "e2e-test-cloud"
)

func (e SourceE2eTestCloudE2eTestCloud) ToPointer() *SourceE2eTestCloudE2eTestCloud {
	return &e
}

func (e *SourceE2eTestCloudE2eTestCloud) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "e2e-test-cloud":
		*e = SourceE2eTestCloudE2eTestCloud(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceE2eTestCloudE2eTestCloud: %v", v)
	}
}

type SourceE2eTestCloudType string

const (
	SourceE2eTestCloudTypeContinuousFeed SourceE2eTestCloudType = "CONTINUOUS_FEED"
)

func (e SourceE2eTestCloudType) ToPointer() *SourceE2eTestCloudType {
	return &e
}

func (e *SourceE2eTestCloudType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONTINUOUS_FEED":
		*e = SourceE2eTestCloudType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceE2eTestCloudType: %v", v)
	}
}

type SourceE2eTestCloud struct {
	// Number of records to emit per stream. Min 1. Max 100 billion.
	MaxMessages int64 `json:"max_messages"`
	// Interval between messages in ms. Min 0 ms. Max 60000 ms (1 minute).
	MessageIntervalMs *int64                        `json:"message_interval_ms,omitempty"`
	MockCatalog       SourceE2eTestCloudMockCatalog `json:"mock_catalog"`
	// When the seed is unspecified, the current time millis will be used as the seed. Range: [0, 1000000].
	Seed       *int64                         `json:"seed,omitempty"`
	SourceType SourceE2eTestCloudE2eTestCloud `json:"sourceType"`
	Type       *SourceE2eTestCloudType        `json:"type,omitempty"`
}
