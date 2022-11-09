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
	"github.com/golang/protobuf/proto"
	"google.golang.org/api/option"
)

// [START spanner_insert_data_with_array_of_proto_columns]
func insertDataWithArrayOfProtoMsgAndEnum(w io.Writer, db string) error {
	ctx := context.Background()
	endpoint := "staging-wrenchworks.sandbox.googleapis.com:443"
	options := []option.ClientOption{option.WithEndpoint(endpoint)}
	client, err := spanner.NewClient(ctx, db, options...)
	if err != nil {
		return err
	}
	defer client.Close()

	singer1ProtoEnum := pb.Genre_ROCK
	singer1ProtoMsg := &pb.SingerInfo{
		SingerId:    proto.Int64(1),
		BirthDate:   proto.String("January"),
		Nationality: proto.String("Country1"),
		Genre:       &singer1ProtoEnum,
	}

	singer2ProtoEnum := pb.Genre_FOLK
	singer2ProtoMsg := &pb.SingerInfo{
		SingerId:    proto.Int64(2),
		BirthDate:   proto.String("February"),
		Nationality: proto.String("Country2"),
		Genre:       &singer2ProtoEnum,
	}

	singerInfoArray := []*pb.SingerInfo{singer1ProtoMsg, singer2ProtoMsg}
	singerGenreArray := []*pb.Genre{&singer1ProtoEnum, &singer2ProtoEnum}
	singerInfoNilArray := []*pb.SingerInfo(nil)
	singerGenreNilArray := []*pb.Genre(nil)
	singerInfoEmptyArray := []*pb.SingerInfo{}
	singerGenreEmptyArray := []*pb.Genre{}

	cols := []string{"SingerId", "FirstName", "LastName", "SingerInfo", "SingerGenre"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("Singers", cols, []interface{}{1, "Singer1", "Singer1", singerInfoArray, singerGenreArray}),
		spanner.InsertOrUpdate("Singers", cols, []interface{}{2, "Singer2", "Singer2", singerInfoNilArray, singerGenreNilArray}),
		spanner.InsertOrUpdate("Singers", cols, []interface{}{3, "Singer3", "Singer3", singerInfoEmptyArray, singerGenreEmptyArray}),
	}

	_, err = client.Apply(ctx, m)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Inserted array of protos data to SingerInfo and SingerGenre columns")
	return nil
}

// [END spanner_insert_data_with_array_of_proto_columns]
