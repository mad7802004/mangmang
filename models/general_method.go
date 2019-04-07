package models

// 批量新建
func Create(data ...interface{}) bool {
	tx := Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false
	}
	for _, model := range data {

		if err := tx.Create(model).Error; err != nil {
			tx.Rollback()
			return false
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return false
	}
	return true

}
