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

package service

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/core"
)

func (s *coreService) CreateOneAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateAssociationKind{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CreateAssociationKind(params, inputData)
}

func (s *coreService) CreateManyAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateManyAssociationKind{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CreateManyAssociationKind(params, inputData)
}

func (s *coreService) SetOneAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.SetAssociationKind{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().SetAssociationKind(params, inputData)
}

func (s *coreService) SetManyAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.SetManyAssociationKind{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().SetManyAssociationKind(params, inputData)
}

func (s *coreService) UpdateAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.UpdateOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().UpdateAssociationKind(params, inputData)
}

func (s *coreService) DeleteAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().DeleteAssociationKind(params, inputData)
}

func (s *coreService) CascadeDeleteAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CascadeDeleteAssociationKind(params, inputData)
}

func (s *coreService) SearchAssociationKind(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.QueryCondition{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	result, err := s.core.AssociationOperation().SearchAssociationKind(params, inputData)
	if err != nil {
		return result, err
	}
	// translate
	for idx := range result.Info {
		s.TranslateAssociationType(params.Lang, &result.Info[idx])
	}
	return result, nil
}

func (s *coreService) CreateModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateModelAssociation{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CreateModelAssociation(params, inputData)
}

func (s *coreService) CreateMainlineModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateModelAssociation{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CreateMainlineModelAssociation(params, inputData)
}

func (s *coreService) SetModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.SetModelAssociation{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().SetModelAssociation(params, inputData)
}

func (s *coreService) UpdateModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.UpdateOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().UpdateModelAssociation(params, inputData)
}

func (s *coreService) SearchModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.QueryCondition{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().SearchModelAssociation(params, inputData)
}

func (s *coreService) DeleteModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().DeleteModelAssociation(params, inputData)
}

func (s *coreService) CascadeDeleteModelAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().DeleteModelAssociation(params, inputData)
}

func (s *coreService) CreateOneInstanceAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateOneInstanceAssociation{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CreateOneInstanceAssociation(params, inputData)
}

func (s *coreService) CreateManyInstanceAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateManyInstanceAssociation{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().CreateManyInstanceAssociation(params, inputData)
}

func (s *coreService) SearchInstanceAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.QueryCondition{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().SearchInstanceAssociation(params, inputData)
}

func (s *coreService) DeleteInstanceAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.AssociationOperation().DeleteInstanceAssociation(params, inputData)
}

func (s *coreService) DeleteInstanceAssociations(ctx core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	objID := inputData.Condition[common.BKObjIDField]
	instID := inputData.Condition[common.BKInstIDField]

	countBKObjID, err := s.core.AssociationOperation().DeleteInstanceAssociation(ctx, inputData)
	if nil != err {
		blog.Errorf("delete instance association error %v, rid: %s", err, ctx.ReqID)
		return nil, err
	}

	bkAsstObjIDCond := metadata.DeleteOption{Condition: mapstr.MapStr{common.BKAsstObjIDField: objID, common.BKAsstInstIDField: instID}}
	countBKAsstObjID, err := s.core.AssociationOperation().DeleteInstanceAssociation(ctx, bkAsstObjIDCond)
	if nil != err {
		blog.Errorf("delete instance to association error %v, rid: %s", err, ctx.ReqID)
		return nil, err
	}

	cnt := countBKObjID.Count + countBKAsstObjID.Count
	return &metadata.DeletedCount{Count: cnt}, nil
}
