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
	"io"

	"cloud.google.com/go/spanner"
	pb "github.com/GoogleCloudPlatform/golang-samples/spanner/spanner_snippets/spanner/testdata/protos"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
)

// [START spanner_insert_data_with_null_proto_columns]

func insertDataWithProtoMsgAndEnumNullValues(w io.Writer, db string) error {
	ctx := context.Background()
	db = "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/singer_proto_db"
	endpoint := "staging-wrenchworks.sandbox.googleapis.com:443"
	options := []option.ClientOption{option.WithEndpoint(endpoint)}
	client, err := spanner.NewClient(ctx, db, options...)
	if err != nil {
		return err
	}
	defer client.Close()

	singer1ProtoEnum := pb.Genre_ROCK
	singer1NullProtoEnum := spanner.NullEnum{pb.Genre_ROCK, true}
	singer1NullProtoMsg := spanner.NullProto{&pb.SingerInfo{
		SingerId:    proto.Int64(1),
		BirthDate:   proto.String("January"),
		Nationality: proto.String("Country1"),
		Genre:       &singer1ProtoEnum,
	}, true}

	singer2ProtoEnum := pb.Genre_FOLK
	singer2NullProtoEnum := spanner.NullEnum{pb.Genre_FOLK, true}
	singer2NullProtoMsg := spanner.NullProto{&pb.SingerInfo{
		SingerId:    proto.Int64(2),
		BirthDate:   proto.String("February"),
		Nationality: proto.String("Country2"),
		Genre:       &singer2ProtoEnum,
	}, true}

	singer3ProtoEnum := pb.Genre_JAZZ
	singer3NullProtoEnum := spanner.NullEnum{pb.Genre_JAZZ, true}
	singer3NullProtoMsg := spanner.NullProto{&pb.SingerInfo{
		SingerId:    proto.Int64(3),
		BirthDate:   proto.String("March"),
		Nationality: proto.String("Country3"),
		Genre:       &singer3ProtoEnum,
	}, true}

	singer4ProtoEnum := pb.Genre_JAZZ
	singer4NullProtoEnum := spanner.NullEnum{pb.Genre_JAZZ, true}
	singer4NullProtoMsg := spanner.NullProto{&pb.SingerInfo{
		SingerId:    proto.Int64(4),
		BirthDate:   proto.String("April"),
		Nationality: proto.String("Country4"),
		Genre:       &singer4ProtoEnum,
	}, true}

	cols := []string{"SingerId", "FirstName", "LastName", "SingerInfo", "SingerGenre"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("Singers", cols, []interface{}{1, "Singer1", "Singer1", singer1NullProtoMsg, singer1NullProtoEnum}),
		spanner.InsertOrUpdate("Singers", cols, []interface{}{2, "Singer2", "Singer2", singer2NullProtoMsg, singer2NullProtoEnum}),
		spanner.InsertOrUpdate("Singers", cols, []interface{}{3, "Singer3", "Singer3", singer3NullProtoMsg, singer3NullProtoEnum}),
		spanner.InsertOrUpdate("Singers", cols, []interface{}{4, "Singer4", "Null Proto with value nil", spanner.NullProto{}, singer4NullProtoEnum}),
		spanner.InsertOrUpdate("Singers", cols, []interface{}{5, "Singer4", "Null Enum with value nil", singer4NullProtoMsg, spanner.NullEnum{}}),
	}
	_, err = client.Apply(ctx, m)
	return err
}

// [END spanner_insert_data_with_null_proto_columns]
