package api

type GetOwnersPanelReq struct {
	OwnerName   string `form:"owner_name" json:"owner_name"`
	PathName    string `form:"path_name" json:"path_name"`
	ProjectName string `form:"project_name" json:"project_name"`
}

func getOwnersPaneReq(req GetOwnersPanelReq) (condition map[string]interface{}) {
	condition = map[string]interface{}{}
	if req.PathName != "" {
		condition["path_name"] = req.PathName
	}
	if req.ProjectName != "" {
		condition["project_name"] = req.ProjectName
	}
	if req.OwnerName != "" {
		condition["owner_name"] = req.OwnerName
	}
	return
}

//func GetOwnersPanel(ctx *gin.Context) (types.Panel, error) {
//	var req GetOwnersPanelReq
//	ctx.BindQuery(&req)
//	reqMap := getOwnersPaneReq(req)
//	comp := template2.Get(config.GetTheme())
//	// get owners title
//	ownerTitle := service.GetOwnersTitle()
//	// get owners content
//	ownerContent := service.GetOwnersContent(reqMap)
//	ownerCount := service.GetOwnersCount(reqMap)
//	//opTable := table2.NewDefaultTable()
//	//info := opTable.GetInfo().AddXssJsFilter().
//	//	HideFilterArea().HideDetailButton().HideEditButton().HideNewButton()
//	//info.AddField("标志", "project_name", db.Varchar).FieldFilterable()
//	//fmt.Println(info)
//	//info.AddField("角色", "owner_name", db.Varchar).FieldJoin(types.Join{
//	//	Table:     config.GetAuthUserTable(),
//	//	JoinField: "id",
//	//	Field:     "user_id",
//	//}).FieldDisplay(func(value types.FieldModel) interface{} {
//	//	return template2.Default().
//	//		Link().
//	//		SetURL(config.Url("/info/manager/detail?__goadmin_detail_pk=") + strconv.Itoa(int(value.Row["user_id"].(int64)))).
//	//		SetContent(template.HTML(value.Value)).
//	//		OpenInNewTab().
//	//		SetTabTitle("Manager Detail").
//	//		GetContent()
//	//}).FieldFilterable()
//	//info.SetTable("goadmin_operation_log").
//	//	SetTitle("333").
//	//	SetDescription("666")
//	//opTable := table2.NewDefaultTable()
//	//formList := opTable.GetForm().AddXssJsFilter()
//
//	table := comp.DataTable().
//		SetHasFilter(true).
//		//SetHideFilterArea(true).
//		// export
//		SetExportUrl("http://baidu.com").
//		SetInfoList(ownerContent).
//		SetPrimaryKey("id").
//		SetThead(ownerTitle)
//
//	allBtns := make(types.Buttons, 0)
//
//	// Add a ajax button action
//	//allBtns = append(allBtns, types.GetDefaultButton("Click me", icon.ArrowLeft, action.Ajax("ajax_id",
//	//	func(ctx *context.Context) (success bool, msg string, data interface{}) {
//	//		fmt.Println("ctx request", ctx.FormValue("id"))
//	//		return true, "ok", nil
//	//	})))
//	//
//	allBtns = append(allBtns, types.GetColumnButton("f1", icon.User, action.Ajax("ajax_id",
//		func(ctx *context.Context) (success bool, msg string, data interface{}) {
//			fmt.Println("ctx request", ctx.FormValue("id"))
//			return true, "ok", nil
//		})))
//	//allBtns = append(allBtns, types.GetActionButton("f2", icon.User, action.Ajax("ajax_id",
//	//	func(ctx *context.Context) (success bool, msg string, data interface{}) {
//	//		fmt.Println("ctx request", ctx.FormValue("id"))
//	//		return true, "ok", nil
//	//	})))
//	btns, btnsJs := allBtns.Content()
//	table = table.SetButtons(btns).SetActionJs(btnsJs)
//	cbs := make(types.Callbacks, 0)
//	for _, btn := range allBtns {
//		cbs = append(cbs, btn.GetAction().GetCallbacks())
//	}
//	body := table.GetContent()
//	//info := table.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
//	//table.AddField("Name", "name", db.Varchar)
//	header := table.GetDataTableHeader()
//	content := comp.Box().
//		SetBody(body).
//		SetNoPadding().
//		SetHeader(header).
//		WithHeadBorder().
//		SetFooter(paginator.Get(paginator.Config{
//			Size:         ownerCount,
//			PageSizeList: pageSizeList,
//			Param:        parameter.GetParam(ctx.Request.URL, defaultSize),
//		}).GetContent()).
//		GetContent()
//	title := "Owner列表"
//	description := "Owner列表"
//	appG := util.Gin{C: ctx}
//	return appG.ResponsePanel(string(content), title, description)
//}
//
//type GetOwnerInfoPanelReq struct {
//	Id string `form:"id" json:"id"`
//}
