package processor

import (
	"cloud.google.com/go/bigquery"
	"context"
)

type Dataset interface {
	// Create creates a dataset in the BigQuery service. An error will be returned if the
	// dataset already exists. Pass in a DatasetMetadata value to configure the dataset.
	Create(ctx context.Context, md *bigquery.DatasetMetadata) (err error)
	// Delete deletes the dataset.  Delete will fail if the dataset is not empty.
	Delete(ctx context.Context) (err error)
	// DeleteWithContents deletes the dataset, as well as contained resources.
	DeleteWithContents(ctx context.Context) (err error)
	// Metadata fetches the metadata for the dataset.
	Metadata(ctx context.Context) (md *bigquery.DatasetMetadata, err error)
	// Update modifies specific Dataset metadata fields.
	// To perform a read-modify-write that protects against intervening reads,
	// set the etag argument to the DatasetMetadata.ETag field from the read.
	// Pass the empty string for etag for a "blind write" that will always succeed.
	Update(ctx context.Context, dm bigquery.DatasetMetadataToUpdate, etag string) (md *bigquery.DatasetMetadata, err error)
	// Table creates a handle to a BigQuery table in the dataset.
	// To determine if a table exists, call Table.Metadata.
	// If the table does not already exist, use Table.Create to create it.
	Table(tableID string) *bigquery.Table
	// Tables returns an iterator over the tables in the Dataset.
	Tables(ctx context.Context) *bigquery.TableIterator
	// Model creates a handle to a BigQuery model in the dataset.
	// To determine if a model exists, call Model.Metadata.
	// If the model does not already exist, you can create it via execution
	// of a CREATE MODEL query.
	Model(modelID string) *bigquery.Model
	// Models returns an iterator over the models in the Dataset.
	Models(ctx context.Context) *bigquery.ModelIterator
	// Routine creates a handle to a BigQuery routine in the dataset.
	// To determine if a routine exists, call Routine.Metadata.
	Routine(routineID string) *bigquery.Routine
	// Routines returns an iterator over the routines in the Dataset.
	Routines(ctx context.Context) *bigquery.RoutineIterator
}
