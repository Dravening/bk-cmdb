package topo_server_test

import (
	"context"

	"configcenter/src/common/metadata"
	commonutil "configcenter/src/common/util"
	"configcenter/src/test"
	"configcenter/src/test/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("inst test", func() {
	var switchInstId1, switchInstId2, routerInstId1 int64

	It("create inst bk_obj_id='bk_switch'", func() {
		test.ClearDatabase()
		input := map[string]interface{}{
			"bk_asset_id":  "101",
			"bk_inst_name": "switch_1",
			"bk_sn":        "201",
		}
		rsp, err := instClient.CreateInst(context.Background(), "0", "bk_switch", header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data["bk_inst_name"].(string)).To(Equal("switch_1"))
		Expect(rsp.Data["bk_asset_id"].(string)).To(Equal("101"))
		Expect(rsp.Data["bk_obj_id"].(string)).To(Equal("bk_switch"))
		Expect(rsp.Data["bk_sn"].(string)).To(Equal("201"))
		switchInstId1, err = commonutil.GetInt64ByInterface(rsp.Data["bk_inst_id"])
		Expect(err).NotTo(HaveOccurred())
	})

	It("create inst bk_obj_id='bk_switch'", func() {
		input := map[string]interface{}{
			"bk_asset_id":  "102",
			"bk_inst_name": "switch_2",
			"bk_sn":        "202",
		}
		rsp, err := instClient.CreateInst(context.Background(), "0", "bk_switch", header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data["bk_inst_name"].(string)).To(Equal("switch_2"))
		Expect(rsp.Data["bk_asset_id"].(string)).To(Equal("102"))
		Expect(rsp.Data["bk_obj_id"].(string)).To(Equal("bk_switch"))
		Expect(rsp.Data["bk_sn"].(string)).To(Equal("202"))
		switchInstId2, err = commonutil.GetInt64ByInterface(rsp.Data["bk_inst_id"])
		Expect(err).NotTo(HaveOccurred())
	})

	It("create inst bk_obj_id='bk_router'", func() {
		input := map[string]interface{}{
			"bk_asset_id":  "101",
			"bk_inst_name": "router_1",
			"bk_sn":        "201",
		}
		rsp, err := instClient.CreateInst(context.Background(), "0", "bk_router", header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data["bk_inst_name"].(string)).To(Equal("router_1"))
		Expect(rsp.Data["bk_asset_id"].(string)).To(Equal("101"))
		Expect(rsp.Data["bk_obj_id"].(string)).To(Equal("bk_router"))
		Expect(rsp.Data["bk_sn"].(string)).To(Equal("201"))
		routerInstId1, err = commonutil.GetInt64ByInterface(rsp.Data["bk_inst_id"])
		Expect(err).NotTo(HaveOccurred())
	})

	It("create association ='bk_router_default_bk_switch'", func() {
		input := &metadata.Association{
			AsstKindID:           "default",
			AsstObjID:            "bk_switch",
			AssociationName:      "bk_router_default_bk_switch",
			AssociationAliasName: "",
			ObjectID:             "bk_router",
			Mapping:              "n:n",
		}
		rsp, err := asstClient.CreateObject(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data.ID).To(Equal(int64(6)))
	})

	It("create inst association ='bk_router1_default_bk_switch1'", func() {
		input := &metadata.CreateAssociationInstRequest{
			ObjectAsstID: "bk_router_default_bk_switch",
			InstID:       routerInstId1,
			AsstInstID:   switchInstId1,
		}
		rsp, err := asstClient.CreateInst(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data.ID).To(Equal(int64(1)))
	})

	It("create inst association ='bk_router1_default_bk_switch2'", func() {
		input := &metadata.CreateAssociationInstRequest{
			ObjectAsstID: "bk_router_default_bk_switch",
			InstID:       routerInstId1,
			AsstInstID:   switchInstId2,
		}
		rsp, err := asstClient.CreateInst(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data.ID).To(Equal(int64(2)))
	})
	//check "SearchAssociationRelatedInst" features available.
	It("search inst association related", func() {
		input := &metadata.SearchAssociationRelatedInstRequest{
			Fields: []string{
				"bk_asst_id",
				"bk_inst_id",
				"bk_obj_id",
				"bk_asst_inst_id",
				"bk_asst_obj_id",
				"bk_obj_asst_id",
			},
			Page: metadata.BasePage{
				Start: 0,
				Limit: 10,
			},
			Condition: metadata.AssociationRelatedInstRequestCond{
				ObjectID: "bk_router",
				InstID:   routerInstId1,
			},
		}
		rsp, err := asstClient.SearchAssociationRelatedInst(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(len(rsp.Data)).To(Equal(2))
		Expect(rsp.Data[0].ObjectAsstID).To(Equal("bk_router_default_bk_switch"))
	})
	//check "SearchAssociationRelatedInst" "limit-check<=500" function.
	It("search inst association related", func() {
		input := &metadata.SearchAssociationRelatedInstRequest{
			Fields: []string{
				"bk_asst_id",
				"bk_inst_id",
				"bk_obj_id",
				"bk_asst_inst_id",
				"bk_asst_obj_id",
				"bk_obj_asst_id",
			},
			Page: metadata.BasePage{
				Start: 0,
				Limit: 501,
			},
			Condition: metadata.AssociationRelatedInstRequestCond{
				ObjectID: "bk_router",
				InstID:   routerInstId1,
			},
		}
		rsp, err := asstClient.SearchAssociationRelatedInst(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(false))
	})
	//check "SearchAssociationRelatedInst" "fields can not be empty." function.
	It("search inst association related", func() {
		input := &metadata.SearchAssociationRelatedInstRequest{
			Fields: []string{},
			Page: metadata.BasePage{
				Start: 0,
				Limit: 10,
			},
			Condition: metadata.AssociationRelatedInstRequestCond{
				ObjectID: "bk_router",
				InstID:   routerInstId1,
			},
		}
		rsp, err := asstClient.SearchAssociationRelatedInst(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(false))
	})
	//check "DeleteAssociationRelatedInst" features available
	It("delete inst association related", func() {
		input := &metadata.DeleteAssociationRelatedInstRequest{}
		input.ObjectID = "bk_router"
		input.InstID = 3

		rsp, err := asstClient.DeleteAssociationRelatedInst(context.Background(), header, input)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
		Expect(rsp.Data).To(Equal("Successfully deleted 2 association record."))

		searchInput := &metadata.SearchAssociationRelatedInstRequest{
			Fields: []string{
				"bk_asst_id",
				"bk_inst_id",
				"bk_obj_id",
				"bk_asst_inst_id",
				"bk_asst_obj_id",
				"bk_obj_asst_id",
			},
			Page: metadata.BasePage{
				Start: 0,
				Limit: 10,
			},
			Condition: metadata.AssociationRelatedInstRequestCond{
				ObjectID: "bk_router",
				InstID:   routerInstId1,
			},
		}
		searchRsp, err := asstClient.SearchAssociationRelatedInst(context.Background(), header, searchInput)
		util.RegisterResponse(searchRsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(searchRsp.Result).To(Equal(true))
		Expect(len(searchRsp.Data)).To(Equal(0))
	})

})
