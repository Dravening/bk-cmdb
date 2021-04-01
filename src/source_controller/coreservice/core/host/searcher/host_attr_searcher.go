/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package searcher

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/common/util"
	"configcenter/src/storage/driver/mongodb"
	"context"
	"fmt"
)

// list host attribute distinct values for IAM.
func (s *Searcher) ListHostAttributeDistinctValues(ctx context.Context, filter mapstr.MapStr, attribute string) (*metadata.ListHostAttrValueResult, error) {
	rid := util.ExtractRequestIDFromContext(ctx)
	// todo: filter is not supported for now
	filter = mapstr.MapStr{}

	attrsList := metadata.ListHostAttrValueResult{
		Count: 0,
		Info:  []string{},
	}

	attrs, err := mongodb.Client().Table(common.BKTableNameBaseHost).Distinct(ctx, attribute, filter)
	if err != nil {
		blog.Errorf("list host attribute values failed, db distinct failed, filter: %+v, err: %+v, rid: %s", filter, err, rid)
		return nil, err
	}

	for _, attr := range attrs {
		attrStr, ok := attr.(string)
		if !ok {
			return nil, fmt.Errorf("unsupported attr type: %v", attr)
		}
		if attrStr != "" {
			attrsList.Info = append(attrsList.Info, attrStr)
		}
	}
	attrsList.Count = len(attrsList.Info)

	return &attrsList, err
}
