package api

import (
	"gin-demo/api/types"
	"gin-demo/internal/logic"
	"gin-demo/model"
	"gin-demo/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IFormController interface {
	RestApi
}

type FormApi struct {
	Logic logic.IFormLogic
}

func NewFormApi() IFormController {
	formApi := FormApi{Logic: logic.NewFormLogic()}
	formApi.Logic.(logic.FormLogic).DB.AutoMigrate(model.Form{})
	return formApi
}

func (c FormApi) Create(ctx *gin.Context) {
	// 获取参数
	var requestForm types.FormRequest
	if err := ctx.ShouldBind(&requestForm); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}
	form, err := c.Logic.Insert(requestForm.Name)
	if err != nil {
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	response.Success(ctx, gin.H{"form": form}, "创建成功")
}

func (c FormApi) Update(ctx *gin.Context) {
	// 获取body 参数
	var requestForm types.FormRequest
	if err := ctx.ShouldBind(&requestForm); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	// 获取 path 参数
	formId, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	var updateForm *model.Form
	form, err := c.Logic.SelectById(formId)
	if err != nil {
		panic(err)
	}
	updateForm, err = c.Logic.Update(*form, requestForm.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"form": updateForm}, "更新成功")
}

func (c FormApi) Show(ctx *gin.Context) {
	// 获取 path 参数
	formId, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	form, err := c.Logic.SelectById(formId)

	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"form": form}, "")
}

func (c FormApi) Delete(ctx *gin.Context) {
	// 获取 path 参数
	formId, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	c.Logic.DeleteById(formId)

	response.Success(ctx, nil, "删除成功")
}
