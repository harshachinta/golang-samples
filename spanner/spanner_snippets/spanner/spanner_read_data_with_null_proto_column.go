// Copyright 2022 Google LLC
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

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
	pb "github.com/GoogleCloudPlatform/golang-samples/spanner/spanner_snippets/spanner/testdata/protos"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// [START spanner_read_data_with_null_proto_columns]

func readDataWithProtoMsgAndEnumNullValues(w io.Writer, db string) error {
	ctx := context.Background()
	db = "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/singer_proto_db"
	endpoint := "staging-wrenchworks.sandbox.googleapis.com:443"
	client, err := spanner.NewClient(ctx, db, option.WithEndpoint(endpoint))
	if err != nil {
		return err
	}
	defer client.Close()

	iter := client.Single().Read(ctx, "Singers", spanner.AllKeys(),
		[]string{"SingerId", "FirstName", "LastName", "SingerInfo", "SingerGenre"})
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		var singerId int64
		var firstName string
		var lastName string
		singerInfo := spanner.NullProtoMessage{ProtoMessageVal: &pb.SingerInfo{}}
		var genreVal pb.Genre
		singerGenre := spanner.NullProtoEnum{ProtoEnumVal: &genreVal}

		if err := row.Columns(&singerId, &firstName, &lastName, &singerInfo, &singerGenre); err != nil {
			return err
		}
		fmt.Fprintf(w, "%d %s %s %v %s\n", singerId, firstName, lastName, singerInfo, singerGenre)
	}
}

// [END spanner_read_data_with_null_proto_columns]
