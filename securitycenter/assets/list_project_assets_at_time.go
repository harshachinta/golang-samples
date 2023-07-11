// Copyright 2019 Google LLC
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

package assets

// [START securitycenter_list_assets_at_time]
import (
	"context"
	"fmt"
	"io"
	"time"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	"cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/iterator"
)

// listAllProjectAssets lists all GCP Projects in orgID at asOf time and prints
// out results to w. orgID is the numeric organization ID of interest.
func listAllProjectAssetsAtTime(w io.Writer, orgID string, asOf time.Time) error {
	// orgID := "12321311"
	// Instantiate a context and a security service client to make API calls.
	ctx := context.Background()
	client, err := securitycenter.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("securitycenter.NewClient: %w", err)
	}
	defer client.Close() // Closing the client safely cleans up background resources.

	// Convert the time to a Timestamp protobuf
	readTime, err := ptypes.TimestampProto(asOf)
	if err != nil {
		return fmt.Errorf("TimestampProto(%v): %w", asOf, err)
	}

	// You can also list assets in a project/ folder. To do so, modify the parent and
	// filter condition.
	req := &securitycenterpb.ListAssetsRequest{
		// Parent must be in one of the following formats:
		//		"organizations/{orgId}"
		//		"projects/{projectId}"
		//		"folders/{folderId}"
		Parent:   fmt.Sprintf("organizations/%s", orgID),
		Filter:   `security_center_properties.resource_type="google.cloud.resourcemanager.Project"`,
		ReadTime: readTime,
	}

	assetsFound := 0
	it := client.ListAssets(ctx, req)
	for {
		result, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("ListAssets: %w", err)
		}
		asset := result.Asset
		properties := asset.SecurityCenterProperties
		fmt.Fprintf(w, "Asset Name: %s,", asset.Name)
		fmt.Fprintf(w, "Resource Name %s,", properties.ResourceName)
		fmt.Fprintf(w, "Resource Type %s\n", properties.ResourceType)
		assetsFound++
	}
	return nil
}

// [END securitycenter_list_assets_at_time]
