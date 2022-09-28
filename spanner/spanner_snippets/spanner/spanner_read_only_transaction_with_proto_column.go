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
	pb "github.com/GoogleCloudPlatform/golang-samples/spanner/spanner_snippets/spanner/testdata"
	"google.golang.org/api/iterator"
)

// [START spanner_read_only_transaction_with_proto_column]

func readOnlyTransactionProtoMsgEnum(w io.Writer, db string) error {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	ro := client.ReadOnlyTransaction()
	defer ro.Close()
	stmt := spanner.Statement{SQL: `SELECT BookId, BookProto, Genre FROM Library`}
	iter := ro.Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		var bookID int64
		bookProto := &pb.Book{}
		var genre pb.Genre
		if err := row.Columns(&bookID, bookProto, &genre); err != nil {
			return err
		}
		fmt.Fprintf(w, "%d %d %s\n", bookID, bookProto, genre)
	}

	iter = ro.Read(ctx, "Library", spanner.AllKeys(), []string{"BookId", "BookProto", "Genre"})
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		var bookID int64
		bookProto := &pb.Book{}
		var genre pb.Genre
		if err := row.Columns(&bookID, bookProto, &genre); err != nil {
			return err
		}
		fmt.Fprintf(w, "%d %d %s\n", bookID, bookProto, genre)
	}
}

// [END spanner_read_only_transaction_with_proto_column]
