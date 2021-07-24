package service

//func GetOwnersContent(condition map[string]interface{}) []map[string]types.InfoItem {
//	owners, err := models.GetOwnersList(condition)
//	if err != nil {
//		// todo log
//		fmt.Println(err)
//		return nil
//	}
//	contents := make([]map[string]types.InfoItem, 0)
//	for _, v := range owners {
//		c := map[string]types.InfoItem{
//			"id":           {Content: template.HTML(strconv.Itoa(v.Model.ID))},
//			"owner_name":   {Content: template.HTML(v.OwnerName)},
//			"project_name": {Content: template.HTML(v.ProjectName)},
//			"path_name":    {Content: template.HTML(v.PathName)},
//			"email":        {Content: template.HTML(v.Email)},
//			"description":  {Content: template.HTML(v.Description)},
//			"create_at":    {Content: template.HTML(v.CreateAt.Format("2006-01-02 15:04:05"))},
//			"update_at":    {Content: template.HTML(v.UpdateAt.Format("2006-01-02 15:04:05"))},
//		}
//		contents = append(contents, c)
//	}
//	return contents
//}
//
//func GetOwnersTitle() []types.TheadItem {
//	return []types.TheadItem{
//		{Head: "ID", Field: "id"},
//		{Head: "owner姓名", Field: "owner_name"},
//		{Head: "大仓名称", Field: "project_name"},
//		{Head: "项目子路径", Field: "path_name"},
//		{Head: "email", Field: "email"},
//		{Head: "备注", Field: "description"},
//		{Head: "创建时间", Field: "create_at"},
//		{Head: "更新时间", Field: "update_at"},
//	}
//}
//
//func GetOwnersCount(condition map[string]interface{}) (count int) {
//	count, err := models.GetOwnersCount(condition)
//	if err != nil {
//		// todo log
//		fmt.Println(err)
//		return
//	}
//	return
//}
