//go:build go1.20
// +build go1.20

// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

// [START spanner_opentelemetry_gfe_metric]
// Ensure that your Go runtime version is supported by the OpenTelemetry-Go compatibility policy before enabling OpenTelemetry instrumentation.
// Refer to compatibility here https://github.com/googleapis/google-cloud-go/blob/main/debug.md#opentelemetry

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func queryWithGFELatencyMetric(w io.Writer, db string) error {
	// db = `projects/<project>/instances/<instance-id>/database/<database-id>`
	ctx := context.Background()

	// Create a new resource to uniquely identify the application
	res, err := newResource()
	if err != nil {
		log.Fatal(err)
	}

	// Enable OpenTelemetry metrics before injecting meter provider.
	spanner.EnableOpenTelemetryMetrics()

	// Create a new meter provider
	meterProvider := getOtlpMeterProvider(ctx, res)
	defer meterProvider.ForceFlush(ctx)

	// Inject meter provider locally via ClientConfig when creating a spanner client.
	client, err := spanner.NewClientWithConfig(ctx, db, spanner.ClientConfig{OpenTelemetryMeterProvider: meterProvider})
	if err != nil {
		return err
	}
	defer client.Close()

	stmt := spanner.Statement{SQL: `SELECT SingerId, AlbumId, AlbumTitle FROM Albums`}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		var singerID, albumID int64
		var albumTitle string
		if err := row.Columns(&singerID, &albumID, &albumTitle); err != nil {
			return err
		}
		fmt.Fprintf(w, "%d %d %s\n", singerID, albumID, albumTitle)
	}
}

// [END spanner_opentelemetry_gfe_metric]
